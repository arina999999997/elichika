package gacha

import (
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
		_, exists := groups[card.GroupMasterID]
		if !exists {
			totalWeight += gamedata.GachaGroup[card.GroupMasterID].GroupWeight
		}
		groups[card.GroupMasterID] = append(groups[card.GroupMasterID], card.CardMasterID)
	}
	groupRand := rand.Int63n(totalWeight)
	for groupID, cardIDs := range groups {
		if gamedata.GachaGroup[groupID].GroupWeight > groupRand { // this group
			return cardIDs[rand.Intn(len(cardIDs))]
		} else {
			groupRand -= gamedata.GachaGroup[groupID].GroupWeight
		}
	}
	panic("this shouldn't happen")
}

func MakeResultCard(session *userdata.Session, cardMasterID int, isGuaranteed bool) model.ResultCard {
	card := session.GetUserCard(cardMasterID)
	cardRarity := session.Gamedata.Card[cardMasterID].CardRarityType
	member := session.GetMember(session.Gamedata.Card[cardMasterID].Member.ID)
	resultCard := model.ResultCard{
		GachaLotType:         1,
		CardMasterID:         cardMasterID,
		Level:                1,
		BeforeGrade:          card.Grade,
		AfterGrade:           card.Grade + 1,
		Content:              nil,
		LimitExceeded:        false,
		BeforeLoveLevelLimit: session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit),
		AfterLoveLevelLimit:  0,
	}
	if isGuaranteed {
		// if more than 1 card have this then the the client might refuse to show the result.
		// it's not doing anything visible, so might as well not set it
		// resultCard.GachaLotType = 2
	}
	if resultCard.AfterGrade == 6 { // maxed out card
		resultCard.AfterGrade = 5
		resultCard.Content = &model.Content{
			ContentType:   13,
			ContentID:     1800,
			ContentAmount: 1,
		}
		// 30 20 10 for UR, SR, R
		for i := cardRarity; i > 10; i -= 10 {
			resultCard.Content.ContentAmount *= 5
		}
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
			session.AddTriggerCardGradeUp(model.TriggerCardGradeUp{
				CardMasterID:         card.CardMasterID,
				BeforeLoveLevelLimit: resultCard.AfterLoveLevelLimit, // this is correct
				AfterLoveLevelLimit:  resultCard.AfterLoveLevelLimit,
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
	draw := *gamedata.GachaDraw[req.GachaDrawMasterID]
	gacha := *gamedata.Gacha[draw.GachaMasterID]
	cardPool := []model.GachaCard{}
	for _, group := range gacha.DbGachaGroups {
		cardPool = append(cardPool, gamedata.GachaGroup[group].Cards...)
		// allow 1 card to be in multiple group
	}
	ctx.Set("gacha_card_pool", cardPool)
	// TODO: gacha recovery and economy
	// for now just get this to work
	resultCards := []model.ResultCard{}
	for _, guaranteeID := range draw.Guarantees {
		gachaGuarantee := gamedata.GachaGuarantee[guaranteeID]
		cardMasterID := GuaranteeHandlers[gachaGuarantee.GuaranteeHandler](ctx, gachaGuarantee)
		if cardMasterID == 0 {
			continue
		}
		resultCards = append(resultCards, MakeResultCard(session, cardMasterID, true))
	}
	for i := len(resultCards); i < draw.DrawCount; i++ {
		resultCards = append(resultCards, MakeResultCard(session, ChooseRandomCard(gamedata, cardPool), false))
	}
	return gacha, resultCards
}
