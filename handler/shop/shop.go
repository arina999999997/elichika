package shop

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// Not implementing this system seriously

func FetchShopTop(ctx *gin.Context) {
	// There is no request body
	resp := response.FetchShopTopResponse{}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: true,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeBillingPack, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: true,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeBillingNormal, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: false,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeEventExchange, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: false,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeItemExchange, arr)
	}
	{
		arr := generic.Array[client.ShopTopIsOpen]{}
		arr.Append(client.ShopTopIsOpen{
			IsOpen: false,
		})
		resp.IsOpenByShopType.Set(enum.ShopTypeBillingSubscription, arr)
	}
	common.JsonResponse(ctx, &resp)
}

func FetchShopPack(ctx *gin.Context) {
	common.JsonResponse(ctx, &response.FetchShopPackResponse{})
}

func FetchShopSnsCoin(ctx *gin.Context) {
	// there is no request body
	// special behaviour to add 10000 gems if someone try to buy gem
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UserModel.UserStatus.FreeSnsCoin += 10000
	
	session.Finalize()
	common.JsonResponse(ctx, &response.FetchShopSnsCoinResponse{})
}

func FetchShopSubscription(ctx *gin.Context) {
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
	router.AddHandler("/shop/fetchShopPack", FetchShopPack)
	router.AddHandler("/shop/fetchShopSnsCoin", FetchShopSnsCoin)
	router.AddHandler("/shop/fetchShopSubscription", FetchShopSubscription)
	router.AddHandler("/shop/fetchShopTop", FetchShopTop)
}
