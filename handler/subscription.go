package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/protocol/response"
	"elichika/userdata"

	"encoding/json"
	"net/http"
	"time"
	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/tidwall/gjson"
)

func FetchSubscriptionPass(ctx *gin.Context) {
	// TODO: everytime someone click on this, give them 1 month of reward because why not
	// reward is cyclic, after the last month it come back to normal

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus()

	respObj := response.FetchSubscriptionPassResponse{
		BeforeContinueCount: subscriptionStatus.RenewalCount,
	}
	// subscriptionStatus.RenewalCount++
	// subscriptionStatus.ContinueCount++
	session.UserModel.UserSubscriptionStatusByID.PushBack(subscriptionStatus)

	session.Finalize("{}", "dummy")

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchShopSubscription(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
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
	// 	ShopProductMasterID: 13001,
	// 	BillingProductType: 3,
	// 	Price: 1,
	// 	ShopBillingProductContent: []response.ShopBillingProductContent{},
	// 	ShopBillingPlatformProduct: &response.ShopBillingPlatformProduct{
	// 		PlatformProductID: "", // this should be the google product id stuff, let's not touch this until we somehow decouple the product system from the client
	// 	},
	// 	Subscription: &response.Subscription {
	// 		SubscriptionMasterID: 13001,
	// 		IsTrial: false,
	// 	},
	// }
	// respObj.ProductList = append(respObj.ProductList, product)
	// it would be fancier if people can click on the purcharse button and it them a subscription
	//  but let's settle on giving everyone who click on this button a subscription anyway

	session.UserModel.UserSubscriptionStatusByID.PushBack(
		model.UserSubscriptionStatus{
			UserID:               userID,
			SubscriptionMasterID: 13001,
			StartDate:            int(time.Now().Unix()),
			ExpireDate:           1<<31 - 1,
			PlatformExpireDate:   1<<31 - 1,
			RenewalCount:         0,
			ContinueCount:        0,
			SubscriptionPassID:   time.Now().UnixNano(),
			AttachID:             "miraizura",
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
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	subscriptionStatus := session.GetSubsriptionStatus()

	subscriptionStatus.ExpireDate = 1<<31 - 1 // preserve the subscription for now
	subscriptionStatus.PlatformExpireDate = subscriptionStatus.ExpireDate
	session.UserModel.UserSubscriptionStatusByID.PushBack(subscriptionStatus)
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
