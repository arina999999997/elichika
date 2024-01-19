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

func SaveReferenceBook(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveReferenceBookRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.InsertReferenceBook(req.ReferenceBookId)

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}
