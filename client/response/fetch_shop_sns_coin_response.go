package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchShopSnsCoinResponse struct {
	ProductList         generic.Array[client.ShopBillingProduct] `json:"product_list"`
	BillingStateInfo    client.BillingStateInfo                  `json:"billing_state_info"`
	BootstrapPickupInfo client.BootstrapPickupInfo               `json:"bootstrap_pickup_info"`
}
