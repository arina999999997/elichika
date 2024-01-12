package request

import (
	"elichika/client"
)

type AccessoryPowerUpRequest struct {
	UserAccessoryId       int64                         `json:"user_accessory_id"`
	PowerUpAccessoryIds   []int64                       `json:"power_up_user_accessory_ids"`
	AccessoryLevelUpItems []client.AccessoryLevelUpItem `json:"accessory_level_up_items"`
}
