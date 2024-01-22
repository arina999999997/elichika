package request

import (
	"elichika/client"
	"elichika/generic"
)

type AccessoryPowerUpRequest struct {
	UserAccessoryId       int64                                      `json:"user_accessory_id"`
	PowerUpAccessoryIds   generic.Array[int64]                       `json:"power_up_user_accessory_ids"`
	AccessoryLevelUpItems generic.Array[client.AccessoryLevelUpItem] `json:"accessory_level_up_items"`
}
