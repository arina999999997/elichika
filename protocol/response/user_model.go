package response

import (
	"elichika/generic"
	"elichika/model"
)

// user_model and user_model_diff both use this, just that the outer key is different
// might be possible to mod the client to only use one thing, but maybe there will be some part where it matter
type UserModel struct {
	UserStatus                                                                 model.UserStatus                                             `json:"user_status"`
	UserMemberByMemberID                                                       generic.ObjectByObjectID[model.UserMember]                   `json:"user_member_by_member_id"`
	UserCardByCardID                                                           generic.ObjectByObjectID[model.UserCard]                     `json:"user_card_by_card_id"`
	UserSuitBySuitID                                                           generic.ObjectByObjectID[model.UserSuit]                     `json:"user_suit_by_suit_id"`
	UserLiveDeckByID                                                           generic.ObjectByObjectID[model.UserLiveDeck]                 `json:"user_live_deck_by_id"`
	UserLivePartyByID                                                          generic.ObjectByObjectID[model.UserLiveParty]                `json:"user_live_party_by_id"`
	UserLessonDeckByID                                                         generic.ObjectByObjectID[model.UserLessonDeck]               `json:"user_lesson_deck_by_id"`
	UserLiveMvDeckByID                                                         generic.ObjectByObjectID[model.UserLiveMvDeck]               `json:"user_live_mv_deck_by_id"`        // TODO: not properly handled
	UserLiveMvDeckCustomByID                                                   generic.ObjectByObjectID[model.UserLiveMvDeck]               `json:"user_live_mv_deck_custom_by_id"` // TODO: not properly handled
	UserLiveDifficultyByDifficultyID                                           generic.ObjectByObjectID[model.UserLiveDifficulty]           `json:"user_live_difficulty_by_difficulty_id"`
	UserStoryMainByStoryMainID                                                 generic.ObjectByObjectID[model.UserStoryMain]                `json:"user_story_main_by_story_main_id"`
	UserStoryMainSelectedByStoryMainCellID                                     generic.ObjectByObjectID[model.UserStoryMainSelected]        `json:"user_story_main_selected_by_story_main_cell_id"`
	UserVoiceByVoiceID                                                         generic.ObjectByObjectID[model.UserVoice]                    `json:"user_voice_by_voice_id"`
	UserEmblemByEmblemID                                                       generic.ObjectByObjectID[model.UserEmblem]                   `json:"user_emblem_by_emblem_id"`       // TODO: not properly handled
	UserGachaTicketByTicketID                                                  generic.ObjectByObjectID[int]                                `json:"user_gacha_ticket_by_ticket_id"` // TODO: not properly handled
	UserGachaPointByPointID                                                    generic.ObjectByObjectID[model.GachaPoint]                   `json:"user_gacha_point_by_point_id"`
	UserLessonEnhancingItemByItemID                                            generic.ObjectByObjectID[model.LessonEnhancingItem]          `json:"user_lesson_enhancing_item_by_item_id"`
	UserTrainingMaterialByItemID                                               generic.ObjectByObjectID[model.TrainingMaterial]             `json:"user_training_material_by_item_id"`
	UserGradeUpItemByItemID                                                    generic.ObjectByObjectID[model.GradeUpItem]                  `json:"user_grade_up_item_by_item_id"`
	UserCustomBackgroundByID                                                   generic.ObjectByObjectID[model.UserCustomBackground]         `json:"user_custom_background_by_id"` // TODO: not properly handled
	UserStorySideByID                                                          generic.ObjectByObjectID[model.UserStorySide]                `json:"user_story_side_by_id"`
	UserStoryMemberByID                                                        generic.ObjectByObjectID[model.UserStoryMember]              `json:"user_story_member_by_id"`
	UserCommunicationMemberDetailBadgeByID                                     generic.ObjectByObjectID[int]                                `json:"user_communication_member_detail_badge_by_id"` // TODO: not properly handled
	UserStoryEventHistoryByID                                                  generic.ObjectByObjectID[model.UserStoryEventHistory]        `json:"user_story_event_history_by_id"`
	UserRecoveryLpByID                                                         generic.ObjectByObjectID[model.RecoverLp]                    `json:"user_recovery_lp_by_id"`
	UserRecoveryApByID                                                         generic.ObjectByObjectID[model.RecoverAp]                    `json:"user_recovery_ap_by_id"`
	UserMissionByMissionID                                                     generic.ObjectByObjectID[int]                                `json:"user_mission_by_mission_id"`        // TODO: not properly handled
	UserDailyMissionByMissionID                                                generic.ObjectByObjectID[int]                                `json:"user_daily_mission_by_mission_id"`  // TODO: not properly handled
	UserWeeklyMissionByMissionID                                               generic.ObjectByObjectID[int]                                `json:"user_weekly_mission_by_mission_id"` // TODO: not properly handled
	UserInfoTriggerBasicByTriggerID                                            generic.ObjectByObjectID[model.TriggerBasic]                 `json:"user_info_trigger_basic_by_trigger_id"`
	UserInfoTriggerCardGradeUpByTriggerID                                      generic.ObjectByObjectID[model.TriggerCardGradeUp]           `json:"user_info_trigger_card_grade_up_by_trigger_id"`
	UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID                    generic.ObjectByObjectID[int]                                `json:"user_info_trigger_member_guild_support_item_expired_by_trigger_id"` // TODO: not properly handled
	UserInfoTriggerMemberLoveLevelUpByTriggerID                                generic.ObjectByObjectID[model.TriggerMemberLoveLevelUp]     `json:"user_info_trigger_member_love_level_up_by_trigger_id"`
	UserAccessoryByUserAccessoryID                                             generic.ObjectByObjectID[model.UserAccessory]                `json:"user_accessory_by_user_accessory_id"`
	UserAccessoryLevelUpItemByID                                               generic.ObjectByObjectID[model.AccessoryLevelUpItem]         `json:"user_accessory_level_up_item_by_id"`
	UserAccessoryRarityUpItemByID                                              generic.ObjectByObjectID[model.AccessoryRarityUpItem]        `json:"user_accessory_rarity_up_item_by_id"`
	UserUnlockScenesByEnum                                                     generic.ObjectByObjectID[model.UserUnlockScene]              `json:"user_unlock_scenes_by_enum"`
	UserSceneTipsByEnum                                                        generic.ObjectByObjectID[model.UserSceneTips]                `json:"user_scene_tips_by_enum"`
	UserRuleDescriptionByID                                                    generic.ObjectByObjectID[model.UserRuleDescription]          `json:"user_rule_description_by_id"` // TODO: not properly handled
	UserExchangeEventPointByID                                                 generic.ObjectByObjectID[model.ExchangeEventPoint]           `json:"user_exchange_event_point_by_id"`
	UserSchoolIdolFestivalIDRewardMissionByID                                  generic.ObjectByObjectID[int]                                `json:"user_school_idol_festival_id_reward_mission_by_id"` // TODO: not properly handled
	UserGpsPresentReceivedByID                                                 generic.ObjectByObjectID[int]                                `json:"user_gps_present_received_by_id"`                   // TODO: not properly handled
	UserEventMarathonByEventMasterID                                           generic.ObjectByObjectID[int]                                `json:"user_event_marathon_by_event_master_id"`            // TODO: not properly handled
	UserEventMiningByEventMasterID                                             generic.ObjectByObjectID[int]                                `json:"user_event_mining_by_event_master_id"`              // TODO: not properly handled
	UserEventCoopByEventMasterID                                               generic.ObjectByObjectID[int]                                `json:"user_event_coop_by_event_master_id"`                // TODO: not properly handled
	UserLiveSkipTicketByID                                                     generic.ObjectByObjectID[model.LiveSkipTicket]               `json:"user_live_skip_ticket_by_id"`
	UserStoryEventUnlockItemByID                                               generic.ObjectByObjectID[model.StoryEventUnlockItem]         `json:"user_story_event_unlock_item_by_id"`
	UserEventMarathonBoosterByID                                               generic.ObjectByObjectID[int]                                `json:"user_event_marathon_booster_by_id"`  // TODO: not properly handled
	UserReferenceBookByID                                                      generic.ObjectByObjectID[model.UserReferenceBook]            `json:"user_reference_book_by_id"`
	UserReviewRequestProcessFlowByID                                           generic.ObjectByObjectID[int]                                `json:"user_review_request_process_flow_by_id"`                                                    // TODO: not properly handled
	UserTowerByTowerID                                                         generic.ObjectByObjectID[int]                                `json:"user_tower_by_tower_id"`                                                                    // TODO: not properly handled
	UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterID generic.ObjectByObjectID[int]                                `json:"user_recovery_tower_card_used_count_item_by_recovery_tower_card_used_count_item_master_id"` // TODO: not properly handled
	UserStoryLinkageByID                                                       generic.ObjectByObjectID[model.UserStoryLinkage]             `json:"user_story_linkage_by_id"`
	UserSubscriptionStatusByID                                                 generic.ObjectByObjectID[int]                                `json:"user_subscription_status_by_id"` // TODO: not properly handled
	UserStoryMainPartDigestMovieByID                                           generic.ObjectByObjectID[model.UserStoryMainPartDigestMovie] `json:"user_story_main_part_digest_movie_by_id"`
	UserMemberGuildByID                                                        generic.ObjectByObjectID[int]                                `json:"user_member_guild_by_id"`                // TODO: not properly handled
	UserMemberGuildSupportItemByID                                             generic.ObjectByObjectID[int]                                `json:"user_member_guild_support_item_by_id"`   // TODO: not properly handled
	UserDailyTheaterByDailyTheaterID                                           generic.ObjectByObjectID[int]                                `json:"user_daily_theater_by_daily_theater_id"` // TODO: not properly handled
	UserPlayListByID                                                           generic.ObjectByObjectID[model.UserPlayListItem]             `json:"user_play_list_by_id"`                   
	UserSetProfileByID                                                         generic.ObjectByObjectID[model.UserSetProfile]               `json:"user_set_profile_by_id"`
	UserSteadyVoltageRankingByID                                               generic.ObjectByObjectID[int]                                `json:"user_steady_voltage_ranking_by_id"`      // TODO: not properly handled
	UserSif2DataLinkByID                                                       generic.ObjectByObjectID[int]                                `json:"user_sif_2_data_link_by_id"`             // TODO: not properly handled
}
