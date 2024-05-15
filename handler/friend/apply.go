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

// request: ApplyFriendRequest
// success response: FriendListResponse
// error response: FriendRecoverableExceptionResponse
func apply(ctx *gin.Context) {
	req := request.ApplyFriendRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	successResponse, failureResponse := user_social.ApplyFriendRequest(session, req.UserId)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	router.AddHandler("/", "POST", "/friend/apply", apply)
}
