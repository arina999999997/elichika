package gacha

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"

	"math/rand"

	"github.com/gin-gonic/gin"
)

// it's not too bad to call this function multiple time, but maybe it's better to have a function that return multiple
func ChooseRandomCard(gamedata *gamedata.Gamedata, cards []model.GachaCard) int {
	if len(cards) == 0 { // no card
		return 0
	}
	groups := map[int]([]int){}
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

func MakeResultCard(session *userdata.Session, cardMasterId int32, isGuaranteed bool) model.ResultCard {
	card := session.GetUserCard(cardMasterId)
	cardRarity := session.Gamedata.Card[cardMasterId].CardRarityType
	member := session.GetMember(session.Gamedata.Card[cardMasterId].Member.Id)
	resultCard := model.ResultCard{
		GachaLotType:         1,
		CardMasterId:         int(cardMasterId),
		Level:                1,
		BeforeGrade:          int(card.Grade),
		AfterGrade:           int(card.Grade) + 1,
		Content:              nil,
		LimitExceeded:        false,
		BeforeLoveLevelLimit: int(session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit)),
		AfterLoveLevelLimit:  0,
	}
	// if isGuaranteed {
	// 	// if more than 1 card have this then the the client might refuse to show the result.
	// 	// it's not doing anything visible, so might as well not set it
	// 	// resultCard.GachaLotType = 2
	// }
	if resultCard.AfterGrade == 6 { // maxed out card
		resultCard.AfterGrade = 5
		resultCard.Content = &client.Content{
			ContentType:   13,
			ContentId:     1800,
			ContentAmount: 1,
		}
		// 30 20 10 for UR, SR, R
		for i := cardRarity; i > 10; i -= 10 {
			resultCard.Content.ContentAmount *= 5
		}
		session.AddResource(*resultCard.Content)
	} else {
		resultCard.AfterLoveLevelLimit = resultCard.BeforeLoveLevelLimit + int(cardRarity/10)
		if resultCard.AfterLoveLevelLimit > int(session.Gamedata.MemberLoveLevelCount) {
			resultCard.AfterLoveLevelLimit = int(session.Gamedata.MemberLoveLevelCount)
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

func HandleGacha(ctx *gin.Context, req model.GachaDrawReq) (model.Gacha, []model.ResultCard) {
	session := ctx.MustGet("session").(*userdata.Session)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	draw := *gamedata.GachaDraw[req.GachaDrawMasterId]
	gacha := *gamedata.Gacha[draw.GachaMasterId]
	cardPool := []model.GachaCard{}
	for _, group := range gacha.DbGachaGroups {
		cardPool = append(cardPool, gamedata.GachaGroup[group].Cards...)
		// allow 1 card to be in multiple group
	}
	ctx.Set("gacha_card_pool", cardPool)
	// TODO: gacha recovery and economy
	// for now just get this to work
	resultCards := []model.ResultCard{}
	for _, guaranteeId := range draw.Guarantees {
		gachaGuarantee := gamedata.GachaGuarantee[guaranteeId]
		cardMasterId := GuaranteeHandlers[gachaGuarantee.GuaranteeHandler](ctx, gachaGuarantee)
		if cardMasterId == 0 {
			continue
		}
		resultCards = append(resultCards, MakeResultCard(session, int32(cardMasterId), true))
	}
	for i := len(resultCards); i < draw.DrawCount; i++ {
		resultCards = append(resultCards, MakeResultCard(session, int32(ChooseRandomCard(gamedata, cardPool)), false))
	}
	return gacha, resultCards
}
