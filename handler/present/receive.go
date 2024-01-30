package present

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func receive(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ReceivePresentRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.ReceivePresentResponse{
		UserModelDiff: &session.UserModel,
	}

	for _, id := range req.Ids.Slice {
		user_present.Receive(session, id, &resp)
	}

	session.Finalize()
	resp.PresentItems = user_present.FetchPresentItems(session)
	resp.PresentHistoryItems = user_present.FetchPresentHistoryItems(session)
	resp.PresentCount = user_present.FetchPresentCount(session)
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/present/receive", receive)
}
