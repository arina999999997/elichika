package shop

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

// Not implementing this system seriously

func fetchShopTop(ctx *gin.Context) {
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

func init() {
	router.AddHandler("/", "POST", "/shop/fetchShopTop", fetchShopTop)
}
