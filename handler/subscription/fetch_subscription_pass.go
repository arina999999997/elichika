package subscription

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"
	"elichika/subsystem/user_subscription_status"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func fetchSubscriptionPass(ctx *gin.Context) {
	// TODO(subscription): everytime someone click on this, give them 1 month of reward because why not
	// reward is cyclic, after the last month it come back to normal
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchSubscriptionPassRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	subscriptionStatus := user_subscription_status.GetUserSubsriptionStatus(session, req.SubscriptionMasterId)

	// subscriptionStatus.RenewalCount++
	// subscriptionStatus.ContinueCount++
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)
	session.Finalize()

	common.JsonResponse(ctx, response.FetchSubscriptionPassResponse{
		BeforeContinueCount: generic.NewNullable(subscriptionStatus.RenewalCount),
	})
}

func init() {
	router.AddHandler("/subscription/fetchSubscriptionPass", fetchSubscriptionPass)
}
