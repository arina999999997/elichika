// content that can only be obtained once like story, voice, emblem, ...
// basically everything that is unlocked once then you have access to it from then on, but you can't modify it except for some small flag
// listed in the order of appearing in user model
package model

import (
	"elichika/generic"
)

var (
	TableNameToInterface = map[string]interface{}{}
)

type UserStoryMain struct {
	StoryMainMasterId int `xorm:"pk 'story_main_master_id'" json:"story_main_master_id"`
}

func (usm *UserStoryMain) Id() int64 {
	return int64(usm.StoryMainMasterId)
}

// for that one part in the main story where you select an idol
type UserStoryMainSelected struct {
	StoryMainCellId int `xorm:"pk 'story_main_cell_id'" json:"story_main_cell_id"`
	SelectedId      int `xorm:"'selected_id'" json:"selected_id"`
}

func (usms *UserStoryMainSelected) Id() int64 {
	return int64(usms.StoryMainCellId)
}

type UserVoice struct {
	NaviVoiceMasterId int  `xorm:"pk 'navi_voice_master_id'" json:"navi_voice_master_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
}

func (uv *UserVoice) Id() int64 {
	return int64(uv.NaviVoiceMasterId)
}

type UserEmblem struct {
	EmblemMId   int     `xorm:"pk 'emblem_m_id'" json:"emblem_m_id"`
	IsNew       bool    `xorm:"'is_new'" json:"is_new"`
	EmblemParam *string `xorm:"'emblem_param'" json:"emblem_param"`
	AcquiredAt  int64   `xorm:"'acquired_at'" json:"acquired_at"`
}

func (ue *UserEmblem) Id() int64 {
	return int64(ue.EmblemMId)
}

type UserCustomBackground struct {
	CustomBackgroundMasterId int  `xorm:"pk 'custom_background_master_id'" json:"custom_background_master_id"`
	IsNew                    bool `xorm:"'is_new'" json:"is_new"`
}

func (ucb *UserCustomBackground) Id() int64 {
	return int64(ucb.CustomBackgroundMasterId)
}

type UserStorySide struct {
	StorySideMasterId int   `xorm:"pk 'story_side_master_id'" json:"story_side_master_id"`
	IsNew             bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt        int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (uss *UserStorySide) Id() int64 {
	return int64(uss.StorySideMasterId)
}

type UserStoryMember struct {
	StoryMemberMasterId int   `xorm:"pk 'story_member_master_id'" json:"story_member_master_id"`
	IsNew               bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt          int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (usm *UserStoryMember) Id() int64 {
	return int64(usm.StoryMemberMasterId)
}

type UserStoryEventHistory struct {
	StoryEventId int `xorm:"pk 'story_event_id'" json:"story_event_id"`
}

func (useh *UserStoryEventHistory) Id() int64 {
	return int64(useh.StoryEventId)
}

type UserUnlockScene struct {
	UnlockSceneType int `xorm:"pk 'unlock_scene_type'" json:"unlock_scene_type"`
	Status          int `xorm:"'status'" json:"status"`
}

func (uus *UserUnlockScene) Id() int64 {
	return int64(uus.UnlockSceneType)
}

type UserSceneTips struct {
	SceneTipsType int `xorm:"pk 'scene_tips_type'" json:"scene_tips_type"`
}

func (ust *UserSceneTips) Id() int64 {
	return int64(ust.SceneTipsType)
}

type UserRuleDescription struct {
	RuleDescriptionId int `xorm:"pk 'rule_description_id'" json:"-"`
	DisplayStatus     int `xorm:"'display_status'" json:"display_status"`
}

func (urd *UserRuleDescription) Id() int64 {
	return int64(urd.RuleDescriptionId)
}
func (urd *UserRuleDescription) SetId(id int64) {
	urd.RuleDescriptionId = int(id)
}

type UserReferenceBook struct {
	ReferenceBookId int `xorm:"pk 'reference_book_id'" json:"reference_book_id"`
}

func (urb *UserReferenceBook) Id() int64 {
	return int64(urb.ReferenceBookId)
}

type UserStoryLinkage struct {
	StoryLinkageCellMasterId int `xorm:"pk 'story_linkage_cell_master_id'" json:"story_linkage_cell_master_id"`
}

func (usl *UserStoryLinkage) Id() int64 {
	return int64(usl.StoryLinkageCellMasterId)
}

type UserStoryMainPartDigestMovie struct {
	StoryMainPartMasterId int `xorm:"pk 'story_main_part_master_id'" json:"story_main_part_master_id"`
}

func (usmpdm *UserStoryMainPartDigestMovie) Id() int64 {
	return int64(usmpdm.StoryMainPartMasterId)
}

func init() {

	TableNameToInterface["u_story_main"] = generic.UserIdWrapper[UserStoryMain]{}
	TableNameToInterface["u_story_main_selected"] = generic.UserIdWrapper[UserStoryMainSelected]{}
	TableNameToInterface["u_voice"] = generic.UserIdWrapper[UserVoice]{}
	TableNameToInterface["u_emblem"] = generic.UserIdWrapper[UserEmblem]{}
	TableNameToInterface["u_custom_background"] = generic.UserIdWrapper[UserCustomBackground]{}
	TableNameToInterface["u_story_side"] = generic.UserIdWrapper[UserStorySide]{}
	TableNameToInterface["u_story_member"] = generic.UserIdWrapper[UserStoryMember]{}
	TableNameToInterface["u_story_event_history"] = generic.UserIdWrapper[UserStoryEventHistory]{}
	TableNameToInterface["u_unlock_scene"] = generic.UserIdWrapper[UserUnlockScene]{}
	TableNameToInterface["u_scene_tips"] = generic.UserIdWrapper[UserSceneTips]{}
	TableNameToInterface["u_rule_description"] = generic.UserIdWrapper[UserRuleDescription]{}
	TableNameToInterface["u_reference_book"] = generic.UserIdWrapper[UserReferenceBook]{}
	TableNameToInterface["u_story_linkage"] = generic.UserIdWrapper[UserStoryLinkage]{}
	TableNameToInterface["u_story_main_part_digest_movie"] = generic.UserIdWrapper[UserStoryMainPartDigestMovie]{}
}
