package user_account_deletion

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_social"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func checkUserAccountDeleted(ctx *gin.Context) {
	req := request.UserAccountDeletionRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	// response with an empty response, or more precisely UserAccountDeletionRecoverableExceptionResponse if the user exist
	session := ctx.MustGet("session").(*userdata.Session)
	if !user_social.UserExist(session, req.UserId) {
		common.AlternativeJsonResponse(ctx, response.UserAccountDeletionRecoverableExceptionResponse{})
	} else {
		common.JsonResponse(ctx, response.EmptyResponse{})
	}
}

func init() {
	router.AddHandler("/", "POST", "/userAccountDeletion/checkUserAccountDeleted", checkUserAccountDeleted)
}
