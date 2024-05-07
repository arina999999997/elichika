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
	session := ctx.MustGet("session").(*userdata.Session)
	if !session.UserExist(req.UserId) {
		common.JsonResponseWithResponseType(ctx, response.UserAccountDeletionRecoverableExceptionResponse{}, 1)
	} else {
		common.JsonResponseWithResponseType(ctx, response.UserAccountDeletionRecoverableExceptionResponse{}, 1)
		// TODO(multiplayer):
		// TODO(friend):
		// This prevent users from viewing other users profile, because doing so currently lead to a frozen game
		// common.JsonResponse(ctx, response.EmptyResponse{})
	}
}

func init() {
	router.AddHandler("/", "POST", "/userAccountDeletion/checkUserAccountDeleted", checkUserAccountDeleted)
}
