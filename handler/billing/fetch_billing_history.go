package billing

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// TODO(billing history): always return empty for now
// technically we can track usage but let's save that for later
func fetchBillingHistory(ctx *gin.Context) {
	req := request.BillingHistoryRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	// userId := int32(ctx.GetInt("user_id"))
	// session := userdata.GetSession(ctx, userId)
	// defer session.Close()

	// session.Finalize()

	common.JsonResponse(ctx, &response.BillingHistoryResponse{})
}

func init() {
	router.AddHandler("/billing/fetchBillingHistory", fetchBillingHistory)
}
