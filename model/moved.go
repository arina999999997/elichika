package model

import (
	"elichika/client"
	"elichika/generic"
)

var (
	TableNameToInterface = map[string]interface{}{}
)

func init() {
	type DbMember struct {
		client.UserMember         `xorm:"extends"`
		LovePanelLevel            int   `xorm:"'love_panel_level' default 1"`
		LovePanelLastLevelCellIds []int `xorm:"'love_panel_last_level_cell_ids' default '[]'"`
	}
	TableNameToInterface["u_member"] = generic.UserIdWrapper[DbMember]{}
	TableNameToInterface["u_suit"] = generic.UserIdWrapper[client.UserSuit]{}
	TableNameToInterface["u_card"] = generic.UserIdWrapper[client.UserCard]{}
	TableNameToInterface["u_lesson_deck"] = generic.UserIdWrapper[client.UserLessonDeck]{}
	TableNameToInterface["u_accessory"] = generic.UserIdWrapper[client.UserAccessory]{}
	TableNameToInterface["u_live_deck"] = generic.UserIdWrapper[client.UserLiveDeck]{}
	TableNameToInterface["u_live_party"] = generic.UserIdWrapper[client.UserLiveParty]{}
	TableNameToInterface["u_live_mv_deck"] = generic.UserIdWrapper[client.UserLiveMvDeck]{}
	TableNameToInterface["u_live_mv_deck_custom"] = generic.UserIdWrapper[client.UserLiveMvDeck]{}
	TableNameToInterface["u_story_main"] = generic.UserIdWrapper[client.UserStoryMain]{}
	TableNameToInterface["u_story_main_selected"] = generic.UserIdWrapper[client.UserStoryMainSelected]{}
	TableNameToInterface["u_voice"] = generic.UserIdWrapper[client.UserVoice]{}
	TableNameToInterface["u_emblem"] = generic.UserIdWrapper[client.UserEmblem]{}
	TableNameToInterface["u_custom_background"] = generic.UserIdWrapper[client.UserCustomBackground]{}
	TableNameToInterface["u_story_side"] = generic.UserIdWrapper[client.UserStorySide]{}
	TableNameToInterface["u_story_member"] = generic.UserIdWrapper[client.UserStoryMember]{}
	TableNameToInterface["u_story_event_history"] = generic.UserIdWrapper[client.UserStoryEventHistory]{}
	TableNameToInterface["u_unlock_scene"] = generic.UserIdWrapper[client.UserUnlockScene]{}
	TableNameToInterface["u_scene_tips"] = generic.UserIdWrapper[client.UserSceneTips]{}
	TableNameToInterface["u_rule_description"] = generic.UserIdWrapper[client.UserRuleDescription]{}
	TableNameToInterface["u_reference_book"] = generic.UserIdWrapper[client.UserReferenceBook]{}
	TableNameToInterface["u_story_linkage"] = generic.UserIdWrapper[client.UserStoryLinkage]{}
	TableNameToInterface["u_story_main_part_digest_movie"] = generic.UserIdWrapper[client.UserStoryMainPartDigestMovie]{}
}
