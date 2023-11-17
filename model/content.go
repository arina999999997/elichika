// content that use can have multiple of
// so materials, gacha points, tickets, ...
package model

import (
	"elichika/enum"

	"fmt"
)

type Content struct {
	ContentType   int   `xorm:"'content_type'" json:"content_type"`
	ContentID     int   `xorm:"'content_id'" json:"content_id"`
	ContentAmount int64 `xorm:"'content_amount'" json:"content_amount"`
}

type RewardDrop struct { // unused
	DropColor int     `json:"drop_color"`
	Content   Content `json:"content"`
	IsRare    bool    `json:"is_rare"`
	BonusType *int    `json:"bonus_type"`
}

type GachaPoint struct {
	PointMasterID int   `json:"point_master_id"`
	Amount        int64 `json:"amount"`
}

func (gp *GachaPoint) ID() int64 {
	return int64(gp.PointMasterID)
}
func (gp *GachaPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGachaPoint { // 5
		panic(fmt.Sprintln("Wrong content for GachaPoint: ", content))
	}
	gp.PointMasterID = content.ContentID
	gp.Amount = content.ContentAmount
}

type LessonEnhancingItem struct {
	EnhancingItemID int   `json:"enhancing_item_id"`
	Amount          int64 `json:"amount"`
}

func (lei *LessonEnhancingItem) ID() int64 {
	return int64(lei.EnhancingItemID)
}
func (lei *LessonEnhancingItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeLessonEnhancingItem { // 6
		panic(fmt.Sprintln("Wrong content for LessonEnhancingItem: ", content))
	}
	lei.EnhancingItemID = content.ContentID
	lei.Amount = content.ContentAmount
}

type TrainingMaterial struct {
	TrainingMaterialMasterID int   `json:"training_material_master_id"`
	Amount                   int64 `json:"amount"`
}

func (tm *TrainingMaterial) ID() int64 {
	return int64(tm.TrainingMaterialMasterID)
}
func (tm *TrainingMaterial) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeTrainingMaterial { // 12
		panic(fmt.Sprintln("Wrong content for TrainingMaterial: ", content))
	}
	tm.TrainingMaterialMasterID = content.ContentID
	tm.Amount = content.ContentAmount
}

type GradeUpItem struct {
	ItemMasterID int   `json:"item_master_id"`
	Amount       int64 `json:"amount"`
}

func (gui *GradeUpItem) ID() int64 {
	return int64(gui.ItemMasterID)
}
func (gui *GradeUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeCardExchange { // 13
		panic(fmt.Sprintln("Wrong content for GradeUpItem: ", content))
	}
	gui.ItemMasterID = content.ContentID
	gui.Amount = content.ContentAmount
}

type RecoverAp struct {
	RecoveryApMasterID int   `json:"recovery_ap_master_id"`
	Amount             int64 `json:"amount"`
}

func (ra *RecoverAp) ID() int64 {
	return int64(ra.RecoveryApMasterID)
}
func (ra *RecoverAp) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeSheetRecoveryAP { // 16
		panic(fmt.Sprintln("Wrong content for RecoverAp: ", content))
	}
	ra.RecoveryApMasterID = content.ContentID
	ra.Amount = content.ContentAmount
}

type RecoverLp struct {
	RecoveryLpMasterID int   `json:"recovery_lp_master_id"`
	Amount             int64 `json:"amount"`
}

func (rl *RecoverLp) ID() int64 {
	return int64(rl.RecoveryLpMasterID)
}
func (rl *RecoverLp) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeRecoveryLP { // 17
		panic(fmt.Sprintln("Wrong content for RecoverLp: ", content))
	}
	rl.RecoveryLpMasterID = content.ContentID
	rl.Amount = content.ContentAmount
}

type ExchangeEventPoint struct {
	PointID int   `json:"-"`
	Amount  int64 `json:"amount"`
}

func (eep *ExchangeEventPoint) ID() int64 {
	return int64(eep.PointID)
}
func (eep *ExchangeEventPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeExchangeEventPoint { // 21
		panic(fmt.Sprintln("Wrong content for ExchangeEventPoint: ", content))
	}
	eep.PointID = content.ContentID
	eep.Amount = content.ContentAmount
}

type AccessoryLevelUpItem struct {
	AccessoryLevelUpItemMasterID int   `json:"accessory_level_up_item_master_id"`
	Amount                       int64 `json:"amount"`
}

func (arui *AccessoryLevelUpItem) ID() int64 {
	return int64(arui.AccessoryLevelUpItemMasterID)
}
func (arui *AccessoryLevelUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeAccessoryLevelUpItem { // 24
		panic(fmt.Sprintln("Wrong content for AccessoryLevelUpItem: ", content))
	}
	arui.AccessoryLevelUpItemMasterID = content.ContentID
	arui.Amount = content.ContentAmount
}

type AccessoryRarityUpItem struct {
	AccessoryRarityUpItemMasterID int   `json:"accessory_rarity_up_item_master_id"`
	Amount                        int64 `json:"amount"`
}

func (arui *AccessoryRarityUpItem) ID() int64 {
	return int64(arui.AccessoryRarityUpItemMasterID)
}
func (arui *AccessoryRarityUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeAccessoryRarityUpItem { // 25
		panic(fmt.Sprintln("Wrong content for AccessoryRarityUpItem: ", content))
	}
	arui.AccessoryRarityUpItemMasterID = content.ContentID
	arui.Amount = content.ContentAmount
}

type LiveSkipTicket struct {
	TicketMasterID int   `json:"ticket_master_id"`
	Amount         int64 `json:"amount"`
}

func (lst *LiveSkipTicket) ID() int64 {
	return int64(lst.TicketMasterID)
}
func (lst *LiveSkipTicket) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeSkipTicket { // 28
		panic(fmt.Sprintln("Wrong content for LiveSkipTicket: ", content))
	}
	lst.TicketMasterID = content.ContentID
	lst.Amount = content.ContentAmount
}

type StoryEventUnlockItem struct {
	StoryEventUnlockItemMasterID int   `json:"story_event_unlock_item_master_id"`
	Amount                       int64 `json:"amount"`
}

func (seui *StoryEventUnlockItem) ID() int64 {
	return int64(seui.StoryEventUnlockItemMasterID)
}
func (seui *StoryEventUnlockItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeStoryEventUnLock { // 30
		panic(fmt.Sprintln("Wrong content for StoryEventUnlockItem: ", content))
	}
	seui.StoryEventUnlockItemMasterID = content.ContentID
	seui.Amount = content.ContentAmount
}
