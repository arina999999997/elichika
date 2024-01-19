package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchSubscriptionPass(ctx *gin.Context) {
	// TODO(subscription): everytime someone click on this, give them 1 month of reward because why not
	// reward is cyclic, after the last month it come back to normal
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchSubscriptionPassRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus(req.SubscriptionMasterId)

	// subscriptionStatus.RenewalCount++
	// subscriptionStatus.ContinueCount++
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)
	session.Finalize("{}", "dummy")

	JsonResponse(ctx, response.FetchSubscriptionPassResponse{
		BeforeContinueCount: generic.NewNullable(subscriptionStatus.RenewalCount),
	})
}

func FetchShopSubscription(ctx *gin.Context) {
	// there's no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchShopSubscriptionResponse{}
	// product := response.ShopBillingProduct{
	// 	ShopProductMasterId: 13001,
	// 	BillingProductType: 3,
	// 	Price: 1,
	// 	ShopBillingProductContent: []response.ShopBillingProductContent{},
	// 	ShopBillingPlatformProduct: &response.ShopBillingPlatformProduct{
	// 		PlatformProductId: "", // this should be the google product id stuff, let's not touch this until we somehow decouple the product system from the client
	// 	},
	// 	Subscription: &response.Subscription {
	// 		SubscriptionMasterId: 13001,
	// 		IsTrial: false,
	// 	},
	// }
	// respObj.ProductList = append(respObj.ProductList, product)
	// it would be fancier if people can click on the purcharse button and it them a subscription
	//  but let's settle on giving everyone who click on this button a subscription anyway

	session.UserModel.UserSubscriptionStatusById.Set(13001, session.GetSubsriptionStatus(13001))

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, &resp)
}

func UpdateSubscription(ctx *gin.Context) {
	// there's no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus(13001)

	subscriptionStatus.ExpireDate = 1<<31 - 1 // preserve the subscription for now
	subscriptionStatus.PlatformExpireDate = subscriptionStatus.ExpireDate
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)
	session.Finalize("{}", "dummy")

	JsonResponse(ctx, &response.UpdateSubscriptionResponse{
		UserModel: &session.UserModel,
	})
}
