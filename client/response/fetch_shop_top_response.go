package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchShopTopResponse struct {
	ShopEventBadge   generic.Array[client.ShopEventBadge]                           `json:"shop_event_badge"`
	Banners          generic.Array[client.Banner1]                                  `json:"banners"`
	UserInfoTrigger  client.UserInfoTrigger                                         `json:"user_info_trigger"`
	IsOpenByShopType generic.Dictionary[int32, generic.Array[client.ShopTopIsOpen]] `json:"is_open_by_shop_type" enum:"ShopType"`
	ProductList      generic.Array[client.ShopBillingProduct]                       `json:"product_list"`
	IsSnsCoinSale    bool                                                           `json:"is_sns_coin_sale"`
}
