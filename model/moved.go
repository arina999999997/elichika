package model

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/generic"
)

var (
	TableNameToInterface = map[string]interface{}{}
)

func init() {
	TableNameToInterface["u_member"] = generic.UserIdWrapper[client.UserMember]{}
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
	type UserRuleDescriptionDbInterface struct {
		RuleDescriptionId   int32                      `xorm:"pk 'rule_description_id'"`
		UserRuleDescription client.UserRuleDescription `xorm:"extends"`
	}
	TableNameToInterface["u_rule_description"] = generic.UserIdWrapper[UserRuleDescriptionDbInterface]{}
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

	type UserReviewRequestProcessFlowDbInterface struct {
		ReviewRequestId              int64                               `xorm:"pk 'review_request_id'"`
		UserReviewRequestProcessFlow client.UserReviewRequestProcessFlow `xorm:"extends"`
	}
	TableNameToInterface["u_review_request_process_flow"] = generic.UserIdWrapper[UserReviewRequestProcessFlowDbInterface]{}

	TableNameToInterface["u_member_guild"] = generic.UserIdWrapper[client.UserMemberGuild]{}
	TableNameToInterface["u_member_guild_support_item"] = generic.UserIdWrapper[client.UserMemberGuildSupportItem]{}

	TableNameToInterface["u_daily_theater"] = generic.UserIdWrapper[client.UserDailyTheater]{}

	TableNameToInterface["u_set_profile"] = generic.UserIdWrapper[client.UserSetProfile]{}

	TableNameToInterface["u_steady_voltage_ranking"] = generic.UserIdWrapper[client.UserSteadyVoltageRanking]{}
	TableNameToInterface["u_play_list"] = generic.UserIdWrapper[client.UserPlayList]{}

	TableNameToInterface["u_tower"] = generic.UserIdWrapper[client.UserTower]{}
	TableNameToInterface["u_subscription_status"] = generic.UserIdWrapper[client.UserSubscriptionStatus]{}

	TableNameToInterface["u_info_trigger_basic"] = generic.UserIdWrapper[client.UserInfoTriggerBasic]{}
	TableNameToInterface["u_info_trigger_card_grade_up"] = generic.UserIdWrapper[client.UserInfoTriggerCardGradeUp]{}
	TableNameToInterface["u_info_trigger_member_love_level_up"] = generic.UserIdWrapper[client.UserInfoTriggerMemberLoveLevelUp]{}
	TableNameToInterface["u_info_trigger_member_guild_support_item_expired"] = generic.UserIdWrapper[client.UserInfoTriggerMemberGuildSupportItemExpired]{}

	TableNameToInterface["u_member_love_panel"] = generic.UserIdWrapper[client.MemberLovePanel]{}

	TableNameToInterface["u_live_difficulty"] = generic.UserIdWrapper[client.UserLiveDifficulty]{}
	TableNameToInterface["u_last_play_live_difficulty_deck"] = generic.UserIdWrapper[client.LastPlayLiveDifficultyDeck]{}

	TableNameToInterface["u_login"] = generic.UserIdWrapper[response.LoginResponse]{}

}
