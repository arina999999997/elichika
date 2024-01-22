package reference_book

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

func SaveReferenceBook(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveReferenceBookRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.InsertReferenceBook(req.ReferenceBookId)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/referenceBook/saveReferenceBook", SaveReferenceBook)
}
