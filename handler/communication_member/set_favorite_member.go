package communication_member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func setFavoriteMember(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SetFavoriteMemberRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.FavoriteMemberId = int32(req.MemberMasterId)
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseFavoriateMember {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseLovePointUp
		// award the initial SR
		// TODO(magic_id)
		intiialSr := int32(100002001 + req.MemberMasterId*10000)
		user_card.AddUserCardByCardMasterId(session, intiialSr)
		session.UserStatus.RecommendCardMasterId = intiialSr // this is for the pop up
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/communicationMember/setFavoriteMember", setFavoriteMember)
}
