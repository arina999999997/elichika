package user_gacha

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/serverdata"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/userdata"

	"math/rand"

	"github.com/gin-gonic/gin"
)

// it's not too bad to call this function multiple time, but maybe it's better to have a function that return multiple
func ChooseRandomCard(gamedata *gamedata.Gamedata, cards []serverdata.GachaCard) int32 {
	if len(cards) == 0 { // no card
		return 0
	}
	groups := map[int32]([]int32){}
	totalWeight := int64(0)
	for _, card := range cards {
		_, exist := groups[card.GroupMasterId]
		if !exist {
			totalWeight += gamedata.GachaGroup[card.GroupMasterId].GroupWeight
		}
		groups[card.GroupMasterId] = append(groups[card.GroupMasterId], card.CardMasterId)
	}
	groupRand := rand.Int63n(totalWeight)
	for groupId, cardIds := range groups {
		if gamedata.GachaGroup[groupId].GroupWeight > groupRand { // this group
			return cardIds[rand.Intn(len(cardIds))]
		} else {
			groupRand -= gamedata.GachaGroup[groupId].GroupWeight
		}
	}
	panic("this shouldn't happen")
}

func MakeResultCard(session *userdata.Session, cardMasterId int32, isGuaranteed bool) client.AddedGachaCardResult {
	addedCardResult := user_card.AddUserCardByCardMasterId(session, cardMasterId)
	addedGachaCardResult := client.AddedGachaCardResult{
		GachaLotType:         enum.GachaLotTypeNormal,
		CardMasterId:         addedCardResult.CardMasterId,
		Level:                addedCardResult.Level,
		BeforeGrade:          addedCardResult.BeforeGrade,
		AfterGrade:           addedCardResult.AfterGrade,
		Content:              addedCardResult.Content,
		LimitExceeded:        addedCardResult.LimitExceeded,
		BeforeLoveLevelLimit: addedCardResult.BeforeLoveLevelLimit,
		AfterLoveLevelLimit:  addedCardResult.AfterLoveLevelLimit,
	}

	// if isGuaranteed {
	// if all 10 cards has this it can crash so let's not assign it
	// addedGachaCardResult.GachaLotType = GachaLotTypeAssurance
	// }
	return addedGachaCardResult
}

func HandleGacha(ctx *gin.Context, req request.DrawGachaRequest) (client.Gacha, generic.List[client.AddedGachaCardResult]) {

	session := ctx.MustGet("session").(*userdata.Session)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	draw := *gamedata.GachaDraw[req.GachaDrawMasterId]
	// payment
	if config.Conf.ResourceConfig().ConsumeGachaCurrency {
		contentType := enum.ContentTypeSnsCoin
		switch draw.GachaPaymentType {
		case enum.GachaPaymentTypeGachaTicket:
		case enum.GachaPaymentTypePremiumGachaTicket:
			contentType = enum.ContentTypeGachaTicket
		}
		user_content.RemoveContent(session, client.Content{
			ContentType:   contentType,
			ContentId:     draw.GachaPaymentMasterId,
			ContentAmount: draw.GachaPaymentAmount,
		})
	}

	gacha := *gamedata.Gacha[req.GachaDrawMasterId/10]
	cardPool := []serverdata.GachaCard{}
	for _, group := range gacha.GachaGroups {
		cardPool = append(cardPool, gamedata.GachaGroup[group].Cards...)
		// allow 1 card to be in multiple group
	}
	ctx.Set("gacha_card_pool", cardPool)
	// TODO: gacha recovery and economy
	// for now just get this to work
	resultCards := generic.List[client.AddedGachaCardResult]{}
	for _, guaranteeId := range gamedata.GachaDrawGuarantee[req.GachaDrawMasterId].GuaranteeIds {
		gachaGuarantee := gamedata.GachaGuarantee[guaranteeId]
		cardMasterId := GuaranteeHandlers[gachaGuarantee.GuaranteeHandler](ctx, gachaGuarantee)
		if cardMasterId == 0 {
			continue
		}
		resultCards.Append(MakeResultCard(session, cardMasterId, true))
	}
	for i := int32(resultCards.Size()); i < draw.DrawCount; i++ {
		resultCards.Append(MakeResultCard(session, int32(ChooseRandomCard(gamedata, cardPool)), false))
	}
	return gacha.ClientGacha, resultCards
}
