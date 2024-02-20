package database

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/generic"
	"elichika/utils"

	"reflect"

	"xorm.io/xorm"
)

// initialise the engine and add the tables that is constructed from client types
func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", config.UserdataPath)
	utils.CheckErr(err)
	Engine.SetMaxOpenConns(50)
	Engine.SetMaxIdleConns(10)
	// Engine.ShowSQL(true)

	AddTable("u_status", generic.UserIdWrapper[client.UserStatus]{})
	AddTable("u_content", generic.InterfaceWithAddedKey[int](
		// user_id can't be pk because we can't mark client content as pk to fetch into map
		// that won't be a problems once gamedata is updated
		client.Content{},
		[]string{"UserId"},
		[]reflect.StructTag{`xorm:"'user_id'"`},
	))
	AddTable("u_member", generic.UserIdWrapper[client.UserMember]{})
	AddTable("u_suit", generic.UserIdWrapper[client.UserSuit]{})
	AddTable("u_card", generic.UserIdWrapper[client.UserCard]{})
	AddTable("u_lesson_deck", generic.UserIdWrapper[client.UserLessonDeck]{})
	AddTable("u_accessory", generic.UserIdWrapper[client.UserAccessory]{})
	AddTable("u_live_deck", generic.UserIdWrapper[client.UserLiveDeck]{})
	AddTable("u_live_party", generic.UserIdWrapper[client.UserLiveParty]{})
	AddTable("u_live_mv_deck", generic.UserIdWrapper[client.UserLiveMvDeck]{})
	AddTable("u_live_mv_deck_custom", generic.UserIdWrapper[client.UserLiveMvDeck]{})
	AddTable("u_story_main", generic.UserIdWrapper[client.UserStoryMain]{})
	AddTable("u_story_main_selected", generic.UserIdWrapper[client.UserStoryMainSelected]{})
	AddTable("u_voice", generic.UserIdWrapper[client.UserVoice]{})
	AddTable("u_emblem", generic.UserIdWrapper[client.UserEmblem]{})
	AddTable("u_custom_background", generic.UserIdWrapper[client.UserCustomBackground]{})
	AddTable("u_story_side", generic.UserIdWrapper[client.UserStorySide]{})
	AddTable("u_story_member", generic.UserIdWrapper[client.UserStoryMember]{})
	AddTable("u_story_event_history", generic.UserIdWrapper[client.UserStoryEventHistory]{})
	AddTable("u_unlock_scenes", generic.UserIdWrapper[client.UserUnlockScene]{})
	AddTable("u_scene_tips", generic.UserIdWrapper[client.UserSceneTips]{})
	type UserRuleDescriptionDbInterface struct {
		RuleDescriptionId   int32                      `xorm:"pk 'rule_description_id'"`
		UserRuleDescription client.UserRuleDescription `xorm:"extends"`
	}
	AddTable("u_rule_description", generic.UserIdWrapper[UserRuleDescriptionDbInterface]{})
	AddTable("u_reference_book", generic.UserIdWrapper[client.UserReferenceBook]{})
	AddTable("u_story_linkage", generic.UserIdWrapper[client.UserStoryLinkage]{})
	AddTable("u_story_main_part_digest_movie", generic.UserIdWrapper[client.UserStoryMainPartDigestMovie]{})
	AddTable("u_communication_member_detail_badge", generic.UserIdWrapper[client.UserCommunicationMemberDetailBadge]{})

	AddTable("u_mission", generic.UserIdWrapper[client.UserMission]{})
	AddTable("u_daily_mission", generic.UserIdWrapper[client.UserDailyMission]{})
	AddTable("u_weekly_mission", generic.UserIdWrapper[client.UserWeeklyMission]{})

	AddTable("u_school_idol_festival_id_reward_mission", generic.UserIdWrapper[client.UserSchoolIdolFestivalIdRewardMission]{})
	AddTable("u_sif_2_data_link", generic.UserIdWrapper[client.UserSif2DataLink]{})
	AddTable("u_gps_present_received", generic.UserIdWrapper[client.UserGpsPresentReceived]{})

	AddTable("u_event_marathon", generic.UserIdWrapper[client.UserEventMarathon]{})
	AddTable("u_event_mining", generic.UserIdWrapper[client.UserEventMining]{})
	AddTable("u_event_coop", generic.UserIdWrapper[client.UserEventCoop]{})

	type UserReviewRequestProcessFlowDbInterface struct {
		ReviewRequestId              int64                               `xorm:"pk 'review_request_id'"`
		UserReviewRequestProcessFlow client.UserReviewRequestProcessFlow `xorm:"extends"`
	}
	AddTable("u_review_request_process_flow", generic.UserIdWrapper[UserReviewRequestProcessFlowDbInterface]{})

	AddTable("u_member_guild", generic.UserIdWrapper[client.UserMemberGuild]{})
	AddTable("u_member_guild_support_item", generic.UserIdWrapper[client.UserMemberGuildSupportItem]{})

	AddTable("u_daily_theater", generic.UserIdWrapper[client.UserDailyTheater]{})

	AddTable("u_set_profile", generic.UserIdWrapper[client.UserSetProfile]{})

	AddTable("u_steady_voltage_ranking", generic.UserIdWrapper[client.UserSteadyVoltageRanking]{})
	AddTable("u_play_list", generic.UserIdWrapper[client.UserPlayList]{})

	AddTable("u_tower", generic.UserIdWrapper[client.UserTower]{})
	AddTable("u_subscription_status", generic.UserIdWrapper[client.UserSubscriptionStatus]{})

	AddTable("u_info_trigger_basic", generic.UserIdWrapper[client.UserInfoTriggerBasic]{})
	AddTable("u_info_trigger_card_grade_up", generic.UserIdWrapper[client.UserInfoTriggerCardGradeUp]{})
	AddTable("u_info_trigger_member_love_level_up", generic.UserIdWrapper[client.UserInfoTriggerMemberLoveLevelUp]{})
	AddTable("u_info_trigger_member_guild_support_item_expired", generic.UserIdWrapper[client.UserInfoTriggerMemberGuildSupportItemExpired]{})

	AddTable("u_member_love_panel", generic.UserIdWrapper[client.MemberLovePanel]{})

	AddTable("u_live_difficulty", generic.UserIdWrapper[client.UserLiveDifficulty]{})
	AddTable("u_last_play_live_difficulty_deck", generic.UserIdWrapper[client.LastPlayLiveDifficultyDeck]{})

	AddTable("u_login", generic.UserIdWrapper[response.LoginResponse]{})

	AddTable("u_live", generic.UserIdWrapper[client.Live]{})
	AddTable("u_start_live_request", generic.UserIdWrapper[request.StartLiveRequest]{})

	AddTable("u_lesson", generic.UserIdWrapper[response.LessonResultResponse]{})

	AddTable("u_card_training_tree_cell", generic.InterfaceWithAddedKey[int](
		client.UserCardTrainingTreeCell{},
		[]string{"UserId", "CardMasterId"},
		[]reflect.StructTag{`xorm:"'user_id'"`, `xorm:"'card_master_id'"`},
	))
}
