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
	"github.com/tidwall/gjson"
)

func checkUserAccountDeleted(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UserAccountDeletionRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// response with an empty response, or more precisely UserAccountDeletionRecoverableExceptionResponse if the user exist
	// do not response otherwise
	if !userdata.UserExist(req.UserId) {
		return
	}

	common.JsonResponse(ctx, response.EmptyResponse{})
}

func init() {
	router.AddHandler("/userAccountDeletion/checkUserAccountDeleted", checkUserAccountDeleted)
}
