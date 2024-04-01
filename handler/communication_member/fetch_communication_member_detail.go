package communication_member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/time"
	"elichika/subsystem/user_member"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchCommunicationMemberDetail(ctx *gin.Context) {
	req := request.FetchCommunicationMemberDetailRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchCommunicationMemberDetailResponse{}
	resp.MemberLovePanels.Append(user_member.GetMemberLovePanel(session, req.MemberId))

	resp.WeekdayState = time.GetWeekdayState(session)
	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/communicationMember/fetchCommunicationMemberDetail", fetchCommunicationMemberDetail)
}
