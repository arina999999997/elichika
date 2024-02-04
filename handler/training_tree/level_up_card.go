package training_tree

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func levelUpCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LevelUpCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingLevelUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingActivateCell
	}

	cardLevel := session.Gamedata.CardLevel[session.Gamedata.Card[req.CardMasterId].CardRarityType]
	card := user_card.GetUserCard(session, req.CardMasterId)
	user_content.RemoveContent(session, item.Gold.Amount(int32(
		cardLevel.GameMoneyPrefixSum[card.Level+req.AdditionalLevel]-cardLevel.GameMoneyPrefixSum[card.Level])))
	user_content.RemoveContent(session, item.EXP.Amount(int32(
		cardLevel.ExpPrefixSum[card.Level+req.AdditionalLevel]-cardLevel.ExpPrefixSum[card.Level])))
	card.Level += req.AdditionalLevel
	user_card.UpdateUserCard(session, card)

	session.Finalize()
	common.JsonResponse(ctx, response.LevelUpCardResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/trainingTree/levelUpCard", levelUpCard)
}
