package shop

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchShopSubscription(ctx *gin.Context) {
	// there's no request body
	userId := int32(ctx.GetInt("user_id"))
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

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/shop/fetchShopSubscription", fetchShopSubscription)
}
