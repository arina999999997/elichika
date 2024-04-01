package live_mv

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live_mv"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func saveDeck(ctx *gin.Context) {
	req := request.SaveLiveMvDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_live_mv.SetLiveMvDeck(session, req.LiveMasterId, req.LiveMvDeckType, req.MemberMasterIdByPos, req.SuitMasterIdByPos, req.ViewStatusByPos)

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/liveMv/saveDeck", saveDeck)
}
