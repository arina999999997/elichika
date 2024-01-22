package client

import (
	"elichika/enum"

	"fmt"
)

type UserAccessoryLevelUpItem struct {
	AccessoryLevelUpItemMasterId int32 `json:"accessory_level_up_item_master_id"`
	Amount                       int32 `json:"amount"`
}

func (ualui *UserAccessoryLevelUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeAccessoryLevelUp { // 24
		panic(fmt.Sprintln("Wrong content for AccessoryLevelUpItem: ", content))
	}
	ualui.AccessoryLevelUpItemMasterId = content.ContentId
	ualui.Amount = content.ContentAmount
}
func (ualui *UserAccessoryLevelUpItem) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeAccessoryLevelUp,
		ContentId:     contentId,
		ContentAmount: ualui.Amount,
	}
}
