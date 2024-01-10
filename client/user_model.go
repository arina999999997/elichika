package client

import (
	"elichika/generic"
)

// user_model and user_model_diff both use this, just that the outer key is different
// might be possible to mod the client to only use one thing, but maybe there will be some part where it matter
type UserModel struct {
	UserStatus                                                                 UserStatus                                                                 `json:"user_status"` // is a pointer
	UserMemberByMemberId                                                       generic.Dictionary[int32, UserMember]                                      `json:"user_member_by_member_id" table:"u_member" key:"member_master_id"`
	UserCardByCardId                                                           generic.ObjectByObjectIdList[UserCard]                                     `json:"user_card_by_card_id"`
	UserSuitBySuitId                                                           generic.ObjectByObjectIdList[UserSuit]                                     `json:"user_suit_by_suit_id"`
	UserLiveDeckById                                                           generic.ObjectByObjectIdList[UserLiveDeck]                                 `json:"user_live_deck_by_id"`
	UserLivePartyById                                                          generic.ObjectByObjectIdList[UserLiveParty]                                `json:"user_live_party_by_id"`
	UserLessonDeckById                                                         generic.ObjectByObjectIdList[UserLessonDeck]                               `json:"user_lesson_deck_by_id"`
	UserLiveMvDeckById                                                         generic.ObjectByObjectIdList[UserLiveMvDeck]                               `json:"user_live_mv_deck_by_id"`
	UserLiveMvDeckCustomById                                                   generic.ObjectByObjectIdList[UserLiveMvDeck]                               `json:"user_live_mv_deck_custom_by_id"`
	UserLiveDifficultyByDifficultyId                                           generic.ObjectByObjectIdList[UserLiveDifficulty]                           `json:"user_live_difficulty_by_difficulty_id"`
	UserStoryMainByStoryMainId                                                 generic.ObjectByObjectIdList[UserStoryMain]                                `json:"user_story_main_by_story_main_id"`
	UserStoryMainSelectedByStoryMainCellId                                     generic.ObjectByObjectIdList[UserStoryMainSelected]                        `json:"user_story_main_selected_by_story_main_cell_id"`
	UserVoiceByVoiceId                                                         generic.ObjectByObjectIdList[UserVoice]                                    `json:"user_voice_by_voice_id"`
	UserEmblemByEmblemId                                                       generic.ObjectByObjectIdList[UserEmblem]                                   `json:"user_emblem_by_emblem_id"` // TODO: not properly handled
	UserGachaTicketByTicketId                                                  generic.ObjectByObjectIdList[UserGachaTicket]                              `json:"user_gacha_ticket_by_ticket_id"`
	UserGachaPointByPointId                                                    generic.ObjectByObjectIdList[UserGachaPoint]                               `json:"user_gacha_point_by_point_id"`
	UserLessonEnhancingItemByItemId                                            generic.ObjectByObjectIdList[UserLessonEnhancingItem]                      `json:"user_lesson_enhancing_item_by_item_id"`
	UserTrainingMaterialByItemId                                               generic.ObjectByObjectIdList[UserTrainingMaterial]                         `json:"user_training_material_by_item_id"`
	UserGradeUpItemByItemId                                                    generic.ObjectByObjectIdList[UserGradeUpItem]                              `json:"user_grade_up_item_by_item_id"`
	UserCustomBackgroundById                                                   generic.ObjectByObjectIdList[UserCustomBackground]                         `json:"user_custom_background_by_id"`
	UserStorySideById                                                          generic.ObjectByObjectIdList[UserStorySide]                                `json:"user_story_side_by_id"`
	UserStoryMemberById                                                        generic.ObjectByObjectIdList[UserStoryMember]                              `json:"user_story_member_by_id"`
	UserCommunicationMemberDetailBadgeById                                     generic.ObjectByObjectIdList[UserCommunicationMemberDetailBadge]           `json:"user_communication_member_detail_badge_by_id"` // TODO: not properly handled
	UserStoryEventHistoryById                                                  generic.ObjectByObjectIdList[UserStoryEventHistory]                        `json:"user_story_event_history_by_id"`
	UserRecoveryLpById                                                         generic.ObjectByObjectIdList[UserRecoveryLp]                               `json:"user_recovery_lp_by_id"`
	UserRecoveryApById                                                         generic.ObjectByObjectIdList[UserRecoveryAp]                               `json:"user_recovery_ap_by_id"`
	UserMissionByMissionId                                                     generic.ObjectByObjectIdList[UserMission]                                  `json:"user_mission_by_mission_id"`        // TODO: not properly handled
	UserDailyMissionByMissionId                                                generic.ObjectByObjectIdList[UserDailyMission]                             `json:"user_daily_mission_by_mission_id"`  // TODO: not properly handled
	UserWeeklyMissionByMissionId                                               generic.ObjectByObjectIdList[UserWeeklyMission]                            `json:"user_weekly_mission_by_mission_id"` // TODO: not properly handled
	UserInfoTriggerBasicByTriggerId                                            generic.ObjectByObjectIdList[UserInfoTriggerBasic]                         `json:"user_info_trigger_basic_by_trigger_id"`
	UserInfoTriggerCardGradeUpByTriggerId                                      generic.ObjectByObjectIdList[UserInfoTriggerCardGradeUp]                   `json:"user_info_trigger_card_grade_up_by_trigger_id"`
	UserInfoTriggerMemberGuildSupportItemExpiredByTriggerId                    generic.ObjectByObjectIdList[UserInfoTriggerMemberGuildSupportItemExpired] `json:"user_info_trigger_member_guild_support_item_expired_by_trigger_id"` // TODO: not properly handled
	UserInfoTriggerMemberLoveLevelUpByTriggerId                                generic.ObjectByObjectIdList[UserInfoTriggerMemberLoveLevelUp]             `json:"user_info_trigger_member_love_level_up_by_trigger_id"`
	UserAccessoryByUserAccessoryId                                             generic.ObjectByObjectIdList[UserAccessory]                                `json:"user_accessory_by_user_accessory_id"`
	UserAccessoryLevelUpItemById                                               generic.ObjectByObjectIdList[UserAccessoryLevelUpItem]                     `json:"user_accessory_level_up_item_by_id"`
	UserAccessoryRarityUpItemById                                              generic.ObjectByObjectIdList[UserAccessoryRarityUpItem]                    `json:"user_accessory_rarity_up_item_by_id"`
	UserUnlockScenesByEnum                                                     generic.ObjectByObjectIdList[UserUnlockScene]                              `json:"user_unlock_scenes_by_enum"`
	UserSceneTipsByEnum                                                        generic.ObjectByObjectIdList[UserSceneTips]                                `json:"user_scene_tips_by_enum"`
	UserRuleDescriptionById                                                    generic.ObjectByObjectIdList[UserRuleDescription]                          `json:"user_rule_description_by_id"` // TODO: not properly handled
	UserExchangeEventPointById                                                 generic.ObjectByObjectIdList[UserExchangeEventPoint]                       `json:"user_exchange_event_point_by_id"`
	UserSchoolIdolFestivalIdRewardMissionById                                  generic.ObjectByObjectIdList[UserSchoolIdolFestivalIdRewardMission]        `json:"user_school_idol_festival_id_reward_mission_by_id"` // not handled
	UserGpsPresentReceivedById                                                 generic.ObjectByObjectIdList[UserGpsPresentReceived]                       `json:"user_gps_present_received_by_id"`                   // not handled
	UserEventMarathonByEventMasterId                                           generic.ObjectByObjectIdList[UserEventMarathon]                            `json:"user_event_marathon_by_event_master_id"`            // TODO: not properly handled
	UserEventMiningByEventMasterId                                             generic.ObjectByObjectIdList[UserEventMining]                              `json:"user_event_mining_by_event_master_id"`              // TODO: not properly handled
	UserEventCoopByEventMasterId                                               generic.ObjectByObjectIdList[UserEventCoop]                                `json:"user_event_coop_by_event_master_id"`                // TODO: not properly handled
	UserLiveSkipTicketById                                                     generic.ObjectByObjectIdList[UserLiveSkipTicket]                           `json:"user_live_skip_ticket_by_id"`
	UserStoryEventUnlockItemById                                               generic.ObjectByObjectIdList[UserStoryEventUnlockItem]                     `json:"user_story_event_unlock_item_by_id"`
	UserEventMarathonBoosterById                                               generic.ObjectByObjectIdList[UserEventMarathonBooster]                     `json:"user_event_marathon_booster_by_id"` // TODO: not properly handled
	UserReferenceBookById                                                      generic.ObjectByObjectIdList[UserReferenceBook]                            `json:"user_reference_book_by_id"`
	UserReviewRequestProcessFlowById                                           generic.ObjectByObjectIdList[UserReviewRequestProcessFlow]                 `json:"user_review_request_process_flow_by_id"`                                                    // TODO: not properly handled
	UserTowerByTowerId                                                         generic.ObjectByObjectIdList[UserTower]                                    `json:"user_tower_by_tower_id"`                                                                    // TODO: not properly handled
	UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterId generic.ObjectByObjectIdList[UserRecoveryTowerCardUsedCountItem]           `json:"user_recovery_tower_card_used_count_item_by_recovery_tower_card_used_count_item_master_id"` // TODO: not properly handled
	UserStoryLinkageById                                                       generic.ObjectByObjectIdList[UserStoryLinkage]                             `json:"user_story_linkage_by_id"`
	UserSubscriptionStatusById                                                 generic.ObjectByObjectIdList[UserSubscriptionStatus]                       `json:"user_subscription_status_by_id"` // TODO: not properly handled
	UserStoryMainPartDigestMovieById                                           generic.ObjectByObjectIdList[UserStoryMainPartDigestMovie]                 `json:"user_story_main_part_digest_movie_by_id"`
	UserMemberGuildById                                                        generic.ObjectByObjectIdList[UserMemberGuild]                              `json:"user_member_guild_by_id"`                // TODO: not properly handled
	UserMemberGuildSupportItemById                                             generic.ObjectByObjectIdList[UserMemberGuildSupportItem]                   `json:"user_member_guild_support_item_by_id"`   // TODO: not properly handled
	UserDailyTheaterByDailyTheaterId                                           generic.ObjectByObjectIdList[UserDailyTheater]                             `json:"user_daily_theater_by_daily_theater_id"` // TODO: not properly handled
	UserPlayListById                                                           generic.ObjectByObjectIdList[UserPlayList]                                 `json:"user_play_list_by_id"`
	UserSetProfileById                                                         generic.ObjectByObjectIdList[UserSetProfile]                               `json:"user_set_profile_by_id"`
	UserSteadyVoltageRankingById                                               generic.ObjectByObjectIdList[UserSteadyVoltageRanking]                     `json:"user_steady_voltage_ranking_by_id"` // TODO: not properly handled
	UserSif2DataLinkById                                                       generic.ObjectByObjectIdList[UserSif2DataLink]                             `json:"user_sif_2_data_link_by_id"`        // TODO: not properly handled
}
