package client

import (
	"elichika/enum"

	"fmt"
)

type UserRecoveryTowerCardUsedCountItem struct {
	RecoveryTowerCardUsedCountItemMasterId int32 `json:"recovery_tower_card_used_count_item_master_id"`
	Amount                                 int32 `json:"amount"`
}

func (urtcuci *UserRecoveryTowerCardUsedCountItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeRecoveryTowerCardUsedCount { // 31
		panic(fmt.Sprintln("Wrong content for RecoveryTowerCardUsedCountItem: ", content))
	}
	urtcuci.RecoveryTowerCardUsedCountItemMasterId = content.ContentId
	urtcuci.Amount = content.ContentAmount
}
func (urtcuci *UserRecoveryTowerCardUsedCountItem) ToContent(contentId int32) Content {
	return Content{
		ContentType:   enum.ContentTypeRecoveryTowerCardUsedCount,
		ContentId:     contentId,
		ContentAmount: urtcuci.Amount,
	}
}
