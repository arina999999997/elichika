package user_account_deletion

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
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
	// do not response otherwise
	// TODO(protocol): This one can have alternative response, check it (some other might have too)
	if !userdata.UserExist(req.UserId) {
		return
	}

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/", "POST", "/userAccountDeletion/checkUserAccountDeleted", checkUserAccountDeleted)
}
