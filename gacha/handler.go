package gacha

import (
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"github.com/gin-gonic/gin"
)

var (
	// handlers take the context
	// the context have the session and the list of available cards in the gacha builtin to it
	// handlers should choose a random card from the card pool if it match, or return 0 if there is no card that fit
	GuaranteeHandlers map[string]func(*gin.Context, *model.GachaGuarantee) int
)

// take no params, return a new card
func GuaranteedNewCard(ctx *gin.Context, gachaGuarantee *model.GachaGuarantee) int {
	cardPool := ctx.MustGet("gacha_card_pool").([]model.GachaCard)
	session := ctx.MustGet("session").(*userdata.Session)
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	newCards := []model.GachaCard{}
	{
		userCards := []int{}
		err := session.Db.Table("u_card").Where("user_id = ?", ctx.GetInt("user_id")).Cols("card_master_id").
			Find(&userCards)
		utils.CheckErr(err)
		cardSet := map[int]bool{}
		for _, id := range userCards {
			cardSet[id] = true
		}
		for _, card := range cardPool {
			_, have := cardSet[card.CardMasterID]
			if !have {
				newCards = append(newCards, card)
			}
		}
	}
	if len(newCards) == 0 { // if empty then choose from the cards that have grade < 5
		userCards := []int{}
		err := session.Db.Table("u_card").Where("user_id = ? AND grade = 5", ctx.GetInt("user_id")).Cols("card_master_id").
			Find(&userCards)
		utils.CheckErr(err)
		cardSet := map[int]bool{}
		for _, id := range userCards {
			cardSet[id] = true
		}
		for _, card := range cardPool {
			_, have := cardSet[card.CardMasterID]
			if !have {
				newCards = append(newCards, card)
			}
		}
	}
	return ChooseRandomCard(gamedata, newCards)
}

func GuaranteedCardSet(ctx *gin.Context, gachaGuarantee *model.GachaGuarantee) int {
	cardPool := ctx.MustGet("gacha_card_pool").([]model.GachaCard)
	// use SQL to get card from the set
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	cardSet := gachaGuarantee.GuaranteedCardSet
	availableCards := []model.GachaCard{}
	for _, card := range cardPool {
		if cardSet[card.CardMasterID] {
			availableCards = append(availableCards, card)
		}
	}
	return ChooseRandomCard(gamedata, availableCards)
}

func init() {
	GuaranteeHandlers = make(map[string]func(*gin.Context, *model.GachaGuarantee) int)
	GuaranteeHandlers["guaranteed_new_card"] = GuaranteedNewCard
	GuaranteeHandlers["guaranteed_card_set"] = GuaranteedCardSet
}
