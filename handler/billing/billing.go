package billing

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

// TODO(billing history): always return empty for now
// technically we can track usage but let's save that for later
func BillingHistory(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.BillingHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// userId := int32(ctx.GetInt("user_id"))
	// session := userdata.GetSession(ctx, userId)
	// defer session.Close()

	// session.Finalize()

	common.JsonResponse(ctx, &response.BillingHistoryResponse{})
}

func UpdateSubscription(ctx *gin.Context) {
	// there's no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus(13001)

	subscriptionStatus.ExpireDate = 1<<31 - 1 // preserve the subscription for now
	subscriptionStatus.PlatformExpireDate = subscriptionStatus.ExpireDate
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)
	
	session.Finalize()
	common.JsonResponse(ctx, &response.UpdateSubscriptionResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/billing/fetchBillingHistory", BillingHistory)
	router.AddHandler("/billing/updateSubscription", UpdateSubscription)

}
