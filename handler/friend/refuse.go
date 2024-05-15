package friend

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_social"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// request: RefuseFriendApprovalRequest
// success response: FriendListResponse
// error response: FriendRecoverableExceptionResponse
func refuse(ctx *gin.Context) {
	req := request.RefuseFriendApprovalRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	successResponse, failureResponse := user_social.RefuseFriendRequest(session, req.UserIds.Slice, req.IsMass)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	router.AddHandler("/", "POST", "/friend/refuse", refuse)
}
