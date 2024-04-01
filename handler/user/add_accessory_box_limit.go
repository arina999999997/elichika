package user

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_status"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func addAccessoryBoxLimit(ctx *gin.Context) {
	req := request.AddAccessoryBoxLimitRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	req.Count *= 10 // count is the amount of time it is performed, not the amount of slot / gem used

	user_status.AddUserAccessoryLimit(session, req.Count)
	user_content.RemoveContent(session, item.StarGem.Amount(req.Count))

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/user/addAccessoryBoxLimit", addAccessoryBoxLimit)
}
