package handler

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(billing history): always return empty for now
// technically we can track usage but let's save that for later
func BillingHistory(ctx *gin.Context) {

	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.BillingHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// userId := ctx.GetInt("user_id")
	// session := userdata.GetSession(ctx, userId)
	// defer session.Close()

	resp := response.BillingHistoryResponse{
		BillingBalanceHistoryList: []client.BillingBalanceHistory{},
		BillingDepositHistoryList: []client.BillingDepositHistory{},
		BillingConsumeHistoryList: []client.BillingConsumeHistory{},
	}
	// session.Finalize()

	JsonResponse(ctx, resp)
}
