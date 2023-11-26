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
func (gp *GachaPoint) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeGachaPoint,
		ContentID:     gp.PointMasterID,
		ContentAmount: gp.Amount,
	}
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
func (lei *LessonEnhancingItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeLessonEnhancingItem,
		ContentID:     lei.EnhancingItemID,
		ContentAmount: lei.Amount,
	}
}

// normally this would need its own table for the specific amounts
// but we just combine everything into normal amount because there's no use for other amount anyway
type GachaTicket struct {
	TicketMasterID int `json:"ticket_master_id"`
	NormalAmount   int `json:"normal_amount"`
	AppleAmount    int `json:"apple_amount"`
	GoogleAmount   int `json:"google_amount"`
}

func (gt *GachaTicket) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeGachaTicket { // 9
		panic(fmt.Sprintln("Wrong content for GachaTicket: ", content))
	}
	gt.TicketMasterID = content.ContentID
	gt.NormalAmount = int(content.ContentAmount)
	gt.AppleAmount = 0
	gt.GoogleAmount = 0
}
func (gt *GachaTicket) ID() int64 {
	return int64(gt.TicketMasterID)
}
func (gt *GachaTicket) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentID:     gt.TicketMasterID,
		ContentAmount: int64(gt.NormalAmount + gt.AppleAmount + gt.GoogleAmount),
	}
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
func (tm *TrainingMaterial) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentID:     tm.TrainingMaterialMasterID,
		ContentAmount: tm.Amount,
	}
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
func (gui *GradeUpItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeCardExchange,
		ContentID:     gui.ItemMasterID,
		ContentAmount: gui.Amount,
	}
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
func (ra *RecoverAp) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeSheetRecoveryAP,
		ContentID:     ra.RecoveryApMasterID,
		ContentAmount: ra.Amount,
	}
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
func (rl *RecoverLp) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeRecoveryLP,
		ContentID:     rl.RecoveryLpMasterID,
		ContentAmount: rl.Amount,
	}
}

type ExchangeEventPoint struct {
	PointID int   `json:"-"`
	Amount  int64 `json:"amount"`
}

func (eep *ExchangeEventPoint) ID() int64 {
	return int64(eep.PointID)
}
func (eep *ExchangeEventPoint) SetID(id int64) {
	eep.PointID = int(id)
}
func (eep *ExchangeEventPoint) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeExchangeEventPoint { // 21
		panic(fmt.Sprintln("Wrong content for ExchangeEventPoint: ", content))
	}
	eep.PointID = content.ContentID
	eep.Amount = content.ContentAmount
}
func (eep *ExchangeEventPoint) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeExchangeEventPoint,
		ContentID:     eep.PointID,
		ContentAmount: eep.Amount,
	}
}

type AccessoryLevelUpItem struct {
	AccessoryLevelUpItemMasterID int   `json:"accessory_level_up_item_master_id"`
	Amount                       int64 `json:"amount"`
}

func (alui *AccessoryLevelUpItem) ID() int64 {
	return int64(alui.AccessoryLevelUpItemMasterID)
}
func (alui *AccessoryLevelUpItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeAccessoryLevelUpItem { // 24
		panic(fmt.Sprintln("Wrong content for AccessoryLevelUpItem: ", content))
	}
	alui.AccessoryLevelUpItemMasterID = content.ContentID
	alui.Amount = content.ContentAmount
}
func (alui *AccessoryLevelUpItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeAccessoryLevelUpItem,
		ContentID:     alui.AccessoryLevelUpItemMasterID,
		ContentAmount: alui.Amount,
	}
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
func (arui *AccessoryRarityUpItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeAccessoryRarityUpItem,
		ContentID:     arui.AccessoryRarityUpItemMasterID,
		ContentAmount: arui.Amount,
	}
}

type EventMarathonBooster struct {
	EventItemID int   `json:"event_item_id"`
	Amount      int64 `json:"amount"`
}

func (emb *EventMarathonBooster) ID() int64 {
	return int64(emb.EventItemID)
}
func (emb *EventMarathonBooster) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeEventMarathonBooster { // 27
		panic(fmt.Sprintln("Wrong content for EventMarathonBooster: ", content))
	}
	emb.EventItemID = content.ContentID
	emb.Amount = content.ContentAmount
}
func (emb *EventMarathonBooster) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeEventMarathonBooster,
		ContentID:     emb.EventItemID,
		ContentAmount: emb.Amount,
	}
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
func (lst *LiveSkipTicket) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeSkipTicket,
		ContentID:     lst.TicketMasterID,
		ContentAmount: lst.Amount,
	}
}

type StoryEventUnlockItem struct {
	StoryEventUnlockItemMasterID int   `json:"story_event_unlock_item_master_id"`
	Amount                       int64 `json:"amount"`
}

func (seui *StoryEventUnlockItem) ID() int64 {
	return int64(seui.StoryEventUnlockItemMasterID)
}
func (seui *StoryEventUnlockItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeStoryEventUnlock { // 30
		panic(fmt.Sprintln("Wrong content for StoryEventUnlockItem: ", content))
	}
	seui.StoryEventUnlockItemMasterID = content.ContentID
	seui.Amount = content.ContentAmount
}
func (seui *StoryEventUnlockItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeStoryEventUnlock,
		ContentID:     seui.StoryEventUnlockItemMasterID,
		ContentAmount: seui.Amount,
	}
}

type RecoveryTowerCardUsedCountItem struct {
	RecoveryTowerCardUsedCountTtemMasterID int   `json:"recovery_tower_card_used_count_item_master_id"`
	Amount                                 int64 `json:"amount"`
}

func (rtcuci *RecoveryTowerCardUsedCountItem) ID() int64 {
	return int64(rtcuci.RecoveryTowerCardUsedCountTtemMasterID)
}
func (rtcuci *RecoveryTowerCardUsedCountItem) FromContent(content Content) {
	if content.ContentType != enum.ContentTypeTowerRecoveryItem { // 31
		panic(fmt.Sprintln("Wrong content for RecoveryTowerCardUsedCountItem: ", content))
	}
	rtcuci.RecoveryTowerCardUsedCountTtemMasterID = content.ContentID
	rtcuci.Amount = content.ContentAmount
}
func (rtcuci *RecoveryTowerCardUsedCountItem) ToContent() Content {
	return Content{
		ContentType:   enum.ContentTypeTowerRecoveryItem,
		ContentID:     rtcuci.RecoveryTowerCardUsedCountTtemMasterID,
		ContentAmount: rtcuci.Amount,
	}
}
