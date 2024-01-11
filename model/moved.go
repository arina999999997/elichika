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
	TableNameToInterface["u_unlock_scenes"] = generic.UserIdWrapper[client.UserUnlockScene]{}
	TableNameToInterface["u_scene_tips"] = generic.UserIdWrapper[client.UserSceneTips]{}
	TableNameToInterface["u_rule_description"] = generic.UserIdWrapper[client.UserRuleDescription]{}
	TableNameToInterface["u_reference_book"] = generic.UserIdWrapper[client.UserReferenceBook]{}
	TableNameToInterface["u_story_linkage"] = generic.UserIdWrapper[client.UserStoryLinkage]{}
	TableNameToInterface["u_story_main_part_digest_movie"] = generic.UserIdWrapper[client.UserStoryMainPartDigestMovie]{}
	TableNameToInterface["u_communication_member_detail_badge"] = generic.UserIdWrapper[client.UserCommunicationMemberDetailBadge]{}
	// TODO(mission): Not handled
	TableNameToInterface["u_mission"] = generic.UserIdWrapper[client.UserMission]{}
	TableNameToInterface["u_daily_mission"] = generic.UserIdWrapper[client.UserDailyMission]{}
	TableNameToInterface["u_weekly_mission"] = generic.UserIdWrapper[client.UserWeeklyMission]{}

	TableNameToInterface["u_school_idol_festival_id_reward_mission"] = generic.UserIdWrapper[client.UserSchoolIdolFestivalIdRewardMission]{}
	TableNameToInterface["u_sif_2_data_link"] = generic.UserIdWrapper[client.UserSif2DataLink]{}
	TableNameToInterface["u_gps_present_received"] = generic.UserIdWrapper[client.UserGpsPresentReceived]{}

	TableNameToInterface["u_event_marathon"] = generic.UserIdWrapper[client.UserEventMarathon]{}
	TableNameToInterface["u_event_mining"] = generic.UserIdWrapper[client.UserEventMining]{}
	TableNameToInterface["u_event_coop"] = generic.UserIdWrapper[client.UserEventCoop]{}
	TableNameToInterface["u_review_request_process_flow"] = generic.UserIdWrapper[client.UserReviewRequestProcessFlow]{}

	TableNameToInterface["u_member_guild"] = generic.UserIdWrapper[client.UserMemberGuild]{}
	TableNameToInterface["u_member_guild_support_item"] = generic.UserIdWrapper[client.UserMemberGuildSupportItem]{}

	TableNameToInterface["u_daily_theater"] = generic.UserIdWrapper[client.UserDailyTheater]{}

	// TODO(refactor): change this database name
	TableNameToInterface["u_custom_set_profile"] = generic.UserIdWrapper[client.UserSetProfile]{}

	TableNameToInterface["u_steady_voltage_ranking"] = generic.UserIdWrapper[client.UserSteadyVoltageRanking]{}
	TableNameToInterface["u_play_list"] = generic.UserIdWrapper[client.UserPlayList]{}

	TableNameToInterface["u_tower"] = generic.UserIdWrapper[client.UserTower]{}
	TableNameToInterface["u_subscription_status"] = generic.UserIdWrapper[client.UserSubscriptionStatus]{}

	TableNameToInterface["u_info_trigger_basic"] = generic.UserIdWrapper[client.UserInfoTriggerBasic]{}
	TableNameToInterface["u_info_trigger_card_grade_up"] = generic.UserIdWrapper[client.UserInfoTriggerCardGradeUp]{}
	TableNameToInterface["u_info_trigger_member_love_level_up"] = generic.UserIdWrapper[client.UserInfoTriggerMemberLoveLevelUp]{}
	TableNameToInterface["u_info_trigger_member_guild_support_item_expired"] = generic.UserIdWrapper[client.UserInfoTriggerMemberGuildSupportItemExpired]{}
}
