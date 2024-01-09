package client

import (
	"elichika/enum"

	"fmt"
)

type UserRecoveryTowerCardUsedCountItem struct {
	RecoveryTowerCardUsedCountItemMasterId int32 `json:"recovery_tower_card_used_count_item_master_id"`
	Amount                                 int32 `json:"amount"`
}

func (urtcuci *UserRecoveryTowerCardUsedCountItem) Id() int64 {
	return int64(urtcuci.RecoveryTowerCardUsedCountItemMasterId)
}
func (urtcuci *UserRecoveryTowerCardUsedCountItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeRecoveryTowerCardUsedCount { // 31
		panic(fmt.Sprintln("Wrong content for RecoveryTowerCardUsedCountItem: ", content))
	}
	urtcuci.RecoveryTowerCardUsedCountItemMasterId = content.ContentId
	urtcuci.Amount = content.ContentAmount
}
func (urtcuci *UserRecoveryTowerCardUsedCountItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeRecoveryTowerCardUsedCount,
		ContentId:     urtcuci.RecoveryTowerCardUsedCountItemMasterId,
		ContentAmount: urtcuci.Amount,
	}
}
