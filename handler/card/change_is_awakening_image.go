package card

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func changeIsAwakeningImage(ctx *gin.Context) {
	req := request.ChangeIsAwakeningImageRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	userCard := user_card.GetUserCard(session, req.CardMasterId)
	userCard.IsAwakeningImage = req.IsAwakeningImage
	user_card.UpdateUserCard(session, userCard)

	common.JsonResponse(ctx, response.ChangeIsAwakeningImageResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/card/changeIsAwakeningImage", changeIsAwakeningImage)
}
