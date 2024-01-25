package user

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_status"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func recoverLp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoverLPRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// TODO(hardcode): technically this is defined in m_recovery_lp
	// and SnsCoinForLpRecover in m_constant
	switch req.ContentId {
	case item.ShowCandy50.ContentId:
		user_status.AddUserLp(session, req.Count.Value*50)
		session.RemoveContent(item.ShowCandy50.Amount(req.Count.Value))
	case item.ShowCandy100.ContentId:
		user_status.AddUserLp(session, req.Count.Value*100)
		session.RemoveContent(item.ShowCandy100.Amount(req.Count.Value))
	case item.StarGem.ContentId:
		user_status.AddUserLp(session, req.Count.Value*100)
		session.RemoveContent(item.StarGem.Amount(req.Count.Value * 10))
	}

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/user/recoverLp", recoverLp)
}
