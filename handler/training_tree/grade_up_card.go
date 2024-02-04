package training_tree

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func gradeUpCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.GradeUpCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	masterCard := session.Gamedata.Card[req.CardMasterId]
	card := user_card.GetUserCard(session, req.CardMasterId)
	member := user_member.GetMember(session, *masterCard.MemberMasterId)

	card.Grade++
	currentLoveLevel := session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit)
	currentLoveLevel += masterCard.CardRarityType / 10 // TODO: Do not hard code this

	if currentLoveLevel > session.Gamedata.MemberLoveLevelCount {
		currentLoveLevel = session.Gamedata.MemberLoveLevelCount
	}
	member.LovePointLimit = session.Gamedata.MemberLoveLevelLovePoint[currentLoveLevel]
	user_card.UpdateUserCard(session, card)
	member.IsNew = true
	user_member.UpdateMember(session, member)
	user_content.RemoveContent(session, masterCard.CardGradeUpItem[card.Grade][req.ContentId])
	// we need to set user_info_trigger_card_grade_up_by_trigger_id
	// for the pop up after limit breaking
	// this trigger show the pop up after limit break
	user_info_trigger.AddTriggerCardGradeUp(session, client.UserInfoTriggerCardGradeUp{
		CardMasterId:         req.CardMasterId,
		BeforeLoveLevelLimit: int32(currentLoveLevel - masterCard.CardRarityType/10),
		AfterLoveLevelLimit:  int32(currentLoveLevel)})

	session.Finalize()
	common.JsonResponse(ctx, response.GradeUpCardResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/trainingTree/gradeUpCard", gradeUpCard)
}
