package handler

import (
	"elichika/enum"

	"elichika/client"
	"elichika/client/response"
	"elichika/generic"
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
	JsonResponse(ctx, &resp)
}

func FetchShopPack(ctx *gin.Context) {
	JsonResponse(ctx, &response.FetchShopPackResponse{})
}

func FetchShopSnsCoin(ctx *gin.Context) {
	// there is no request body
	// special behaviour to add 10000 gems if someone try to buy gem
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UserModel.UserStatus.FreeSnsCoin += 10000
	session.Finalize("{}", "dummy")
	JsonResponse(ctx, &response.FetchShopSnsCoinResponse{})
}
