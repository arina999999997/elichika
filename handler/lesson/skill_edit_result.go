package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func skillEditResult(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SkillEditResultRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for cardMasterId, selectedSkills := range req.SelectedSkillIds.Map {
		card := user_card.GetUserCard(session, cardMasterId)
		for i, skillId := range selectedSkills.Slice {
			switch i {
			case 0:
				card.AdditionalPassiveSkill1Id = skillId
			case 1:
				card.AdditionalPassiveSkill2Id = skillId
			case 2:
				card.AdditionalPassiveSkill3Id = skillId
			case 3:
				card.AdditionalPassiveSkill4Id = skillId
			}
		}
		user_card.UpdateUserCard(session, card) // this is always updated even if no skill change happen
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/lesson/skillEditResult", skillEditResult)
}
