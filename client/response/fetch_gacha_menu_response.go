package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchGachaMenuResponse struct {
	GachaList        generic.List[client.Gacha]                `json:"gacha_list"`
	GachaUnconfirmed generic.Nullable[client.GachaUnconfirmed] `json:"gacha_unconfirmed"` // this is not a nullable in client, it's a pointer, but it can be null
	UserModelDiff    *client.UserModel                         `json:"user_model_diff"`
}
