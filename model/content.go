// content that use can have multiple of
// so materials, gacha points, tickets, ...
package model

import (
	"elichika/client"
	"elichika/enum"

	"fmt"
)

type RewardDrop struct { // unused
	DropColor int32          `json:"drop_color"`
	Content   client.Content `json:"content"`
	IsRare    bool           `json:"is_rare"`
	BonusType *int           `json:"bonus_type"`
}

type GachaPoint struct {
	PointMasterId int32 `json:"point_master_id"`
	Amount        int32 `json:"amount"`
}

func (gp *GachaPoint) Id() int64 {
	return int64(gp.PointMasterId)
}
func (gp *GachaPoint) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeGachaPoint { // 5
		panic(fmt.Sprintln("Wrong content for GachaPoint: ", content))
	}
	gp.PointMasterId = content.ContentId
	gp.Amount = content.ContentAmount
}
func (gp *GachaPoint) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeGachaPoint,
		ContentId:     gp.PointMasterId,
		ContentAmount: gp.Amount,
	}
}

type LessonEnhancingItem struct {
	EnhancingItemId int32 `json:"enhancing_item_id"`
	Amount          int32 `json:"amount"`
}

func (lei *LessonEnhancingItem) Id() int64 {
	return int64(lei.EnhancingItemId)
}
func (lei *LessonEnhancingItem) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeLessonEnhancingItem { // 6
		panic(fmt.Sprintln("Wrong content for LessonEnhancingItem: ", content))
	}
	lei.EnhancingItemId = content.ContentId
	lei.Amount = content.ContentAmount
}
func (lei *LessonEnhancingItem) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeLessonEnhancingItem,
		ContentId:     lei.EnhancingItemId,
		ContentAmount: lei.Amount,
	}
}

// normally this would need its own table for the specific amounts
// but we just combine everything into normal amount because there's no use for other amount anyway
type GachaTicket struct {
	TicketMasterId int32 `json:"ticket_master_id"`
	NormalAmount   int   `json:"normal_amount"`
	AppleAmount    int   `json:"apple_amount"`
	GoogleAmount   int   `json:"google_amount"`
}

func (gt *GachaTicket) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeGachaTicket { // 9
		panic(fmt.Sprintln("Wrong content for GachaTicket: ", content))
	}
	gt.TicketMasterId = content.ContentId
	gt.NormalAmount = int(content.ContentAmount)
	gt.AppleAmount = 0
	gt.GoogleAmount = 0
}
func (gt *GachaTicket) Id() int64 {
	return int64(gt.TicketMasterId)
}
func (gt *GachaTicket) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeGachaTicket,
		ContentId:     gt.TicketMasterId,
		ContentAmount: int32(gt.NormalAmount + gt.AppleAmount + gt.GoogleAmount),
	}
}

type TrainingMaterial struct {
	TrainingMaterialMasterId int32 `json:"training_material_master_id"`
	Amount                   int32 `json:"amount"`
}

func (tm *TrainingMaterial) Id() int64 {
	return int64(tm.TrainingMaterialMasterId)
}
func (tm *TrainingMaterial) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeTrainingMaterial { // 12
		panic(fmt.Sprintln("Wrong content for TrainingMaterial: ", content))
	}
	tm.TrainingMaterialMasterId = content.ContentId
	tm.Amount = content.ContentAmount
}
func (tm *TrainingMaterial) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeTrainingMaterial,
		ContentId:     tm.TrainingMaterialMasterId,
		ContentAmount: tm.Amount,
	}
}

type GradeUpItem struct {
	ItemMasterId int32 `json:"item_master_id"`
	Amount       int32 `json:"amount"`
}

func (gui *GradeUpItem) Id() int64 {
	return int64(gui.ItemMasterId)
}
func (gui *GradeUpItem) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeGradeUpper { // 13
		panic(fmt.Sprintln("Wrong content for GradeUpItem: ", content))
	}
	gui.ItemMasterId = content.ContentId
	gui.Amount = content.ContentAmount
}
func (gui *GradeUpItem) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeGradeUpper,
		ContentId:     gui.ItemMasterId,
		ContentAmount: gui.Amount,
	}
}

type RecoverAp struct {
	RecoveryApMasterId int32 `json:"recovery_ap_master_id"`
	Amount             int32 `json:"amount"`
}

func (ra *RecoverAp) Id() int64 {
	return int64(ra.RecoveryApMasterId)
}
func (ra *RecoverAp) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeRecoveryAp { // 16
		panic(fmt.Sprintln("Wrong content for RecoverAp: ", content))
	}
	ra.RecoveryApMasterId = content.ContentId
	ra.Amount = content.ContentAmount
}
func (ra *RecoverAp) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeRecoveryAp,
		ContentId:     ra.RecoveryApMasterId,
		ContentAmount: ra.Amount,
	}
}

type RecoverLp struct {
	RecoveryLpMasterId int32 `json:"recovery_lp_master_id"`
	Amount             int32 `json:"amount"`
}

func (rl *RecoverLp) Id() int64 {
	return int64(rl.RecoveryLpMasterId)
}
func (rl *RecoverLp) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeRecoveryLp { // 17
		panic(fmt.Sprintln("Wrong content for RecoverLp: ", content))
	}
	rl.RecoveryLpMasterId = content.ContentId
	rl.Amount = content.ContentAmount
}
func (rl *RecoverLp) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeRecoveryLp,
		ContentId:     rl.RecoveryLpMasterId,
		ContentAmount: rl.Amount,
	}
}

type ExchangeEventPoint struct {
	PointId int32 `json:"-"`
	Amount  int32 `json:"amount"`
}

func (eep *ExchangeEventPoint) Id() int64 {
	return int64(eep.PointId)
}
func (eep *ExchangeEventPoint) SetId(id int64) {
	eep.PointId = int32(id)
}
func (eep *ExchangeEventPoint) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeExchangeEventPoint { // 21
		panic(fmt.Sprintln("Wrong content for ExchangeEventPoint: ", content))
	}
	eep.PointId = content.ContentId
	eep.Amount = content.ContentAmount
}
func (eep *ExchangeEventPoint) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeExchangeEventPoint,
		ContentId:     eep.PointId,
		ContentAmount: eep.Amount,
	}
}

type AccessoryLevelUpItem struct {
	AccessoryLevelUpItemMasterId int32 `json:"accessory_level_up_item_master_id"`
	Amount                       int32 `json:"amount"`
}

func (alui *AccessoryLevelUpItem) Id() int64 {
	return int64(alui.AccessoryLevelUpItemMasterId)
}
func (alui *AccessoryLevelUpItem) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeAccessoryLevelUp { // 24
		panic(fmt.Sprintln("Wrong content for AccessoryLevelUpItem: ", content))
	}
	alui.AccessoryLevelUpItemMasterId = content.ContentId
	alui.Amount = content.ContentAmount
}
func (alui *AccessoryLevelUpItem) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeAccessoryLevelUp,
		ContentId:     alui.AccessoryLevelUpItemMasterId,
		ContentAmount: alui.Amount,
	}
}

type AccessoryRarityUpItem struct {
	AccessoryRarityUpItemMasterId int32 `json:"accessory_rarity_up_item_master_id"`
	Amount                        int32 `json:"amount"`
}

func (arui *AccessoryRarityUpItem) Id() int64 {
	return int64(arui.AccessoryRarityUpItemMasterId)
}
func (arui *AccessoryRarityUpItem) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeAccessoryRarityUp { // 25
		panic(fmt.Sprintln("Wrong content for AccessoryRarityUpItem: ", content))
	}
	arui.AccessoryRarityUpItemMasterId = content.ContentId
	arui.Amount = content.ContentAmount
}
func (arui *AccessoryRarityUpItem) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeAccessoryRarityUp,
		ContentId:     arui.AccessoryRarityUpItemMasterId,
		ContentAmount: arui.Amount,
	}
}

type EventMarathonBooster struct {
	EventItemId int32 `json:"event_item_id"`
	Amount      int32 `json:"amount"`
}

func (emb *EventMarathonBooster) Id() int64 {
	return int64(emb.EventItemId)
}
func (emb *EventMarathonBooster) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeEventMarathonBooster { // 27
		panic(fmt.Sprintln("Wrong content for EventMarathonBooster: ", content))
	}
	emb.EventItemId = content.ContentId
	emb.Amount = content.ContentAmount
}
func (emb *EventMarathonBooster) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeEventMarathonBooster,
		ContentId:     emb.EventItemId,
		ContentAmount: emb.Amount,
	}
}

type LiveSkipTicket struct {
	TicketMasterId int32 `json:"ticket_master_id"`
	Amount         int32 `json:"amount"`
}

func (lst *LiveSkipTicket) Id() int64 {
	return int64(lst.TicketMasterId)
}
func (lst *LiveSkipTicket) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeLiveSkipTicket { // 28
		panic(fmt.Sprintln("Wrong content for LiveSkipTicket: ", content))
	}
	lst.TicketMasterId = content.ContentId
	lst.Amount = content.ContentAmount
}
func (lst *LiveSkipTicket) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeLiveSkipTicket,
		ContentId:     lst.TicketMasterId,
		ContentAmount: lst.Amount,
	}
}

type StoryEventUnlockItem struct {
	StoryEventUnlockItemMasterId int32 `json:"story_event_unlock_item_master_id"`
	Amount                       int32 `json:"amount"`
}

func (seui *StoryEventUnlockItem) Id() int64 {
	return int64(seui.StoryEventUnlockItemMasterId)
}
func (seui *StoryEventUnlockItem) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeStoryEventUnlock { // 30
		panic(fmt.Sprintln("Wrong content for StoryEventUnlockItem: ", content))
	}
	seui.StoryEventUnlockItemMasterId = content.ContentId
	seui.Amount = content.ContentAmount
}
func (seui *StoryEventUnlockItem) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeStoryEventUnlock,
		ContentId:     seui.StoryEventUnlockItemMasterId,
		ContentAmount: seui.Amount,
	}
}

type RecoveryTowerCardUsedCountItem struct {
	RecoveryTowerCardUsedCountTtemMasterId int32 `json:"recovery_tower_card_used_count_item_master_id"`
	Amount                                 int32 `json:"amount"`
}

func (rtcuci *RecoveryTowerCardUsedCountItem) Id() int64 {
	return int64(rtcuci.RecoveryTowerCardUsedCountTtemMasterId)
}
func (rtcuci *RecoveryTowerCardUsedCountItem) FromContent(content client.Content) {
	if content.ContentType != enum.ContentTypeRecoveryTowerCardUsedCount { // 31
		panic(fmt.Sprintln("Wrong content for RecoveryTowerCardUsedCountItem: ", content))
	}
	rtcuci.RecoveryTowerCardUsedCountTtemMasterId = content.ContentId
	rtcuci.Amount = content.ContentAmount
}
func (rtcuci *RecoveryTowerCardUsedCountItem) ToContent() client.Content {
	return client.Content{
		ContentType:   enum.ContentTypeRecoveryTowerCardUsedCount,
		ContentId:     rtcuci.RecoveryTowerCardUsedCountTtemMasterId,
		ContentAmount: rtcuci.Amount,
	}
}
