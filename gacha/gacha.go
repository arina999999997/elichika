package gacha

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/item"
	"elichika/serverdata"
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
	card := session.GetUserCard(cardMasterId)
	cardRarity := session.Gamedata.Card[cardMasterId].CardRarityType
	member := session.GetMember(session.Gamedata.Card[cardMasterId].Member.Id)
	resultCard := client.AddedGachaCardResult{
		GachaLotType:         enum.GachaLotTypeNormal,
		CardMasterId:         cardMasterId,
		Level:                1,
		BeforeGrade:          card.Grade,
		AfterGrade:           card.Grade + 1,
		BeforeLoveLevelLimit: session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit),
	}
	if isGuaranteed {
		resultCard.GachaLotType = enum.GachaLotTypeAssurance
	}
	if resultCard.AfterGrade == 6 { // maxed out card
		resultCard.AfterGrade = 5
		content := item.SchoolIdolRadiance
		// 30 20 10 for UR, SR, R
		for i := cardRarity; i > 10; i -= 10 {
			content.ContentAmount *= 5
		}
		resultCard.Content = generic.NewNullable(content)
		session.AddResource(content)
	} else {
		resultCard.AfterLoveLevelLimit = resultCard.BeforeLoveLevelLimit + cardRarity/10
		if resultCard.AfterLoveLevelLimit > session.Gamedata.MemberLoveLevelCount {
			resultCard.AfterLoveLevelLimit = session.Gamedata.MemberLoveLevelCount
		}
		member.LovePointLimit = session.Gamedata.MemberLoveLevelLovePoint[resultCard.AfterLoveLevelLimit]
		card.Grade++ // new grade,
		if card.Grade == 0 {
			// entirely new card
			member.OwnedCardCount++
			resultCard.BeforeGrade = 0
		} else {
			// add trigger card grade up so animation play when opening the card
			session.AddTriggerCardGradeUp(client.UserInfoTriggerCardGradeUp{
				CardMasterId:         card.CardMasterId,
				BeforeLoveLevelLimit: int32(resultCard.AfterLoveLevelLimit), // this is correct
				AfterLoveLevelLimit:  int32(resultCard.AfterLoveLevelLimit),
			})
		}
		// update the card and member
		session.UpdateUserCard(card)
		session.UpdateMember(member)
	}
	return resultCard
}

func HandleGacha(ctx *gin.Context, req request.DrawGachaRequest) (client.Gacha, generic.List[client.AddedGachaCardResult]) {
	session := ctx.MustGet("session").(*userdata.Session)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	draw := *gamedata.GachaDraw[req.GachaDrawMasterId]

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
