package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func CheckUserAccountDeleted(ctx *gin.Context) {
	// request is UserAccountDeletionRequest
	// response with an empty response, or more precisely UserAccountDeletionRecoverableExceptionResponse if the user exist
	// do not response otherwise
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UserAccountDeletionRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	if !userdata.UserExist(req.UserId) {
		return
	}

	JsonResponse(ctx, response.EmptyResponse{})
}
