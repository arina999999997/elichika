package billing

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_subscription_status"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func updateSubscription(ctx *gin.Context) {
	// there's no request body

	session := ctx.MustGet("session").(*userdata.Session)

	// TODO(subscription): Implement subscription logic better
	subscriptionStatus := user_subscription_status.GetUserSubsriptionStatus(session, 13001)

	subscriptionStatus.ExpireDate = 1<<31 - 1 // preserve the subscription for now
	subscriptionStatus.PlatformExpireDate = subscriptionStatus.ExpireDate
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)

	common.JsonResponse(ctx, &response.UpdateSubscriptionResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/billing/updateSubscription", updateSubscription)
}
