package client

import (
	"elichika/enum"

	"fmt"
)

type UserAccessoryRarityUpItem struct {
	AccessoryRarityUpItemMasterId int32 `json:"accessory_rarity_up_item_master_id"`
	Amount                        int32 `json:"amount"`
}

func (uarui *UserAccessoryRarityUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeAccessoryRarityUp { // 25
		panic(fmt.Sprintln("Wrong content for AccessoryRarityUpItem: ", content))
	}
	uarui.AccessoryRarityUpItemMasterId = content.ContentId
	uarui.Amount = content.ContentAmount
}
func (uarui *UserAccessoryRarityUpItem) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeAccessoryRarityUp,
		ContentId:     contentId,
		ContentAmount: uarui.Amount,
	}
}
