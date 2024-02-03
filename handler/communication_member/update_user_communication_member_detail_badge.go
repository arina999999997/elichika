package communication_member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func updateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateMemberDetailBadgeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	detailBadge := user_member.GetUserCommunicationMemberDetailBadge(session, req.MemberMasterId)
	switch req.CommunicationMemberDetailBadgeType {
	case enum.CommunicationMemberDetailBadgeTypeStoryMember:
		detailBadge.IsStoryMemberBadge = false
	case enum.CommunicationMemberDetailBadgeTypeStorySide:
		detailBadge.IsStorySideBadge = false
	case enum.CommunicationMemberDetailBadgeTypeVoice:
		detailBadge.IsVoiceBadge = false
	case enum.CommunicationMemberDetailBadgeTypeTheme:
		detailBadge.IsThemeBadge = false
	case enum.CommunicationMemberDetailBadgeTypeCard:
		detailBadge.IsCardBadge = false
	case enum.CommunicationMemberDetailBadgeTypeMusic:
		detailBadge.IsMusicBadge = false
	default:
		panic("unknown type")
	}
	user_member.UpdateUserCommunicationMemberDetailBadge(session, detailBadge)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/communicationMember/updateUserCommunicationMemberDetailBadge", updateUserCommunicationMemberDetailBadge)
}
