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
)

func setFavoriteMember(ctx *gin.Context) {
	req := request.SetFavoriteMemberRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	session.UserStatus.FavoriteMemberId = int32(req.MemberMasterId)
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseFavoriateMember {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseLovePointUp
		// award the initial SR
		// TODO(magic_id)
		intiialSr := int32(100002001 + req.MemberMasterId*10000)
		user_card.AddUserCardByCardMasterId(session, intiialSr)
		session.UserStatus.RecommendCardMasterId = intiialSr // this is for the pop up
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/communicationMember/setFavoriteMember", setFavoriteMember)
}
