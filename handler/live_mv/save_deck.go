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

	session := ctx.MustGet("session").(*userdata.Session)

	user_live_mv.SetLiveMvDeck(session, req.LiveMasterId, req.LiveMvDeckType, req.MemberMasterIdByPos, req.SuitMasterIdByPos, req.ViewStatusByPos)

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/liveMv/saveDeck", saveDeck)
}
