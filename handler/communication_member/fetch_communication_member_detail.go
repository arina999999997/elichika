package communication_member

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/time"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func fetchCommunicationMemberDetail(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchCommunicationMemberDetailRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchCommunicationMemberDetailResponse{}
	resp.MemberLovePanels.Append(session.GetMemberLovePanel(req.MemberId))

	resp.WeekdayState = time.GetWeekdayState(session)
	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/communicationMember/fetchCommunicationMemberDetail", fetchCommunicationMemberDetail)
}
