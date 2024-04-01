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

func recoverLp(ctx *gin.Context) {
	req := request.RecoverLPRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// TODO(hardcode): technically this is defined in m_recovery_lp
	// and SnsCoinForLpRecover in m_constant
	switch req.ContentId {
	case item.ShowCandy50.ContentId:
		user_status.AddUserLp(session, req.Count.Value*50)
		user_content.RemoveContent(session, item.ShowCandy50.Amount(req.Count.Value))
	case item.ShowCandy100.ContentId:
		user_status.AddUserLp(session, req.Count.Value*100)
		user_content.RemoveContent(session, item.ShowCandy100.Amount(req.Count.Value))
	case item.StarGem.ContentId:
		user_status.AddUserLp(session, req.Count.Value*100)
		user_content.RemoveContent(session, item.StarGem.Amount(req.Count.Value*10))
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/user/recoverLp", recoverLp)
}
