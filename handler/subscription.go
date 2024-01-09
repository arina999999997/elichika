package handler

import (
	"elichika/config"
	"elichika/client"
	"elichika/protocol/response"
	"elichika/userdata"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
)

func FetchSubscriptionPass(ctx *gin.Context) {
	// TODO: everytime someone click on this, give them 1 month of reward because why not
	// reward is cyclic, after the last month it come back to normal

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus()

	respObj := response.FetchSubscriptionPassResponse{
		BeforeContinueCount: subscriptionStatus.RenewalCount,
	}
	// subscriptionStatus.RenewalCount++
	// subscriptionStatus.ContinueCount++
	session.UserModel.UserSubscriptionStatusById.PushBack(subscriptionStatus)

	session.Finalize("{}", "dummy")

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchShopSubscription(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	// there's no request body
	respObj := response.ShopSubscriptionResponse{
		BillingStateInfo: response.BillingStateInfo{
			Age:                        42,
			CurrentMonthPurcharsePrice: 1337,
		},
		ProductList: []response.ShopBillingProduct{},
	}
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

	session.UserModel.UserSubscriptionStatusById.PushBack(
		client.UserSubscriptionStatus{
			SubscriptionMasterId: 13001,
			StartDate:            session.Time.Unix(),
			ExpireDate:           1<<31 - 1,
			PlatformExpireDate:   1<<31 - 1,
			SubscriptionPassId:   session.Time.UnixNano(),
			AttachId:             "miraizura",
			IsAutoRenew:          true,
			IsDoneTrial:          true,
		})
	session.Finalize("{}", "dummy")

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func UpdateSubscription(ctx *gin.Context) {
	// there's no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus()

	subscriptionStatus.ExpireDate = 1<<31 - 1 // preserve the subscription for now
	subscriptionStatus.PlatformExpireDate = subscriptionStatus.ExpireDate
	session.UserModel.UserSubscriptionStatusById.PushBack(subscriptionStatus)
	session.Finalize("{}", "dummy")
	respObj := response.UpdateSubscriptionResponse{
		UserModel: &session.UserModel,
		BillingStateInfo: response.BillingStateInfo{
			Age:                        42,
			CurrentMonthPurcharsePrice: 1337,
		},
	}

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
