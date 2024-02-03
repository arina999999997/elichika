package emblem

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

func activateEmblem(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ActivateEmblemRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.EmblemId = req.EmblemMasterId

	session.Finalize()
	common.JsonResponse(ctx, response.ActivateEmblemResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/emblem/activateEmblem", activateEmblem)
}
