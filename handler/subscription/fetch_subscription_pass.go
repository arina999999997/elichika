package subscription

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_subscription_status"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchSubscriptionPass(ctx *gin.Context) {
	// TODO(subscription): everytime someone click on this, give them 1 month of reward because why not
	// just make reward cyclic, after the last month it come back to normal
	req := request.FetchSubscriptionPassRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	subscriptionStatus := user_subscription_status.GetUserSubsriptionStatus(session, req.SubscriptionMasterId)

	// subscriptionStatus.RenewalCount++
	// subscriptionStatus.ContinueCount++
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)

	common.JsonResponse(ctx, response.FetchSubscriptionPassResponse{
		BeforeContinueCount: generic.NewNullable(subscriptionStatus.RenewalCount),
	})
}

func init() {
	router.AddHandler("/", "POST", "/subscription/fetchSubscriptionPass", fetchSubscriptionPass)
}
