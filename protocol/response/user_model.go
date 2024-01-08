package response

import (
	"elichika/client"
	"elichika/generic"
	"elichika/model"

	"fmt"
	"reflect"
)

// user_model and user_model_diff both use this, just that the outer key is different
// might be possible to mod the client to only use one thing, but maybe there will be some part where it matter
type UserModel struct {
	UserStatus                                                                 model.UserStatus                                                          `json:"user_status"`
	UserMemberByMemberId                                                       generic.ObjectByObjectIdList[client.UserMember]                           `json:"user_member_by_member_id"`
	UserCardByCardId                                                           generic.ObjectByObjectIdList[client.UserCard]                             `json:"user_card_by_card_id"`
	UserSuitBySuitId                                                           generic.ObjectByObjectIdList[model.UserSuit]                              `json:"user_suit_by_suit_id"`
	UserLiveDeckById                                                           generic.ObjectByObjectIdList[model.UserLiveDeck]                          `json:"user_live_deck_by_id"`
	UserLivePartyById                                                          generic.ObjectByObjectIdList[model.UserLiveParty]                         `json:"user_live_party_by_id"`
	UserLessonDeckById                                                         generic.ObjectByObjectIdList[model.UserLessonDeck]                        `json:"user_lesson_deck_by_id"`
	UserLiveMvDeckById                                                         generic.ObjectByObjectIdList[model.UserLiveMvDeck]                        `json:"user_live_mv_deck_by_id"`
	UserLiveMvDeckCustomById                                                   generic.ObjectByObjectIdList[model.UserLiveMvDeck]                        `json:"user_live_mv_deck_custom_by_id"`
	UserLiveDifficultyByDifficultyId                                           generic.ObjectByObjectIdList[model.UserLiveDifficulty]                    `json:"user_live_difficulty_by_difficulty_id"`
	UserStoryMainByStoryMainId                                                 generic.ObjectByObjectIdList[model.UserStoryMain]                         `json:"user_story_main_by_story_main_id"`
	UserStoryMainSelectedByStoryMainCellId                                     generic.ObjectByObjectIdList[model.UserStoryMainSelected]                 `json:"user_story_main_selected_by_story_main_cell_id"`
	UserVoiceByVoiceId                                                         generic.ObjectByObjectIdList[model.UserVoice]                             `json:"user_voice_by_voice_id"`
	UserEmblemByEmblemId                                                       generic.ObjectByObjectIdList[model.UserEmblem]                            `json:"user_emblem_by_emblem_id"` // TODO: not properly handled
	UserGachaTicketByTicketId                                                  generic.ObjectByObjectIdList[model.GachaTicket]                           `json:"user_gacha_ticket_by_ticket_id"`
	UserGachaPointByPointId                                                    generic.ObjectByObjectIdList[model.GachaPoint]                            `json:"user_gacha_point_by_point_id"`
	UserLessonEnhancingItemByItemId                                            generic.ObjectByObjectIdList[model.LessonEnhancingItem]                   `json:"user_lesson_enhancing_item_by_item_id"`
	UserTrainingMaterialByItemId                                               generic.ObjectByObjectIdList[model.TrainingMaterial]                      `json:"user_training_material_by_item_id"`
	UserGradeUpItemByItemId                                                    generic.ObjectByObjectIdList[model.GradeUpItem]                           `json:"user_grade_up_item_by_item_id"`
	UserCustomBackgroundById                                                   generic.ObjectByObjectIdList[model.UserCustomBackground]                  `json:"user_custom_background_by_id"`
	UserStorySideById                                                          generic.ObjectByObjectIdList[model.UserStorySide]                         `json:"user_story_side_by_id"`
	UserStoryMemberById                                                        generic.ObjectByObjectIdList[model.UserStoryMember]                       `json:"user_story_member_by_id"`
	UserCommunicationMemberDetailBadgeById                                     generic.ObjectByObjectIdList[model.UserCommunicationMemberDetailBadge]    `json:"user_communication_member_detail_badge_by_id"` // TODO: not properly handled
	UserStoryEventHistoryById                                                  generic.ObjectByObjectIdList[model.UserStoryEventHistory]                 `json:"user_story_event_history_by_id"`
	UserRecoveryLpById                                                         generic.ObjectByObjectIdList[model.RecoverLp]                             `json:"user_recovery_lp_by_id"`
	UserRecoveryApById                                                         generic.ObjectByObjectIdList[model.RecoverAp]                             `json:"user_recovery_ap_by_id"`
	UserMissionByMissionId                                                     generic.ObjectByObjectIdList[model.UserMission]                           `json:"user_mission_by_mission_id"`        // TODO: not properly handled
	UserDailyMissionByMissionId                                                generic.ObjectByObjectIdList[model.UserDailyMission]                      `json:"user_daily_mission_by_mission_id"`  // TODO: not properly handled
	UserWeeklyMissionByMissionId                                               generic.ObjectByObjectIdList[model.UserWeeklyMission]                     `json:"user_weekly_mission_by_mission_id"` // TODO: not properly handled
	UserInfoTriggerBasicByTriggerId                                            generic.ObjectByObjectIdList[model.TriggerBasic]                          `json:"user_info_trigger_basic_by_trigger_id"`
	UserInfoTriggerCardGradeUpByTriggerId                                      generic.ObjectByObjectIdList[model.TriggerCardGradeUp]                    `json:"user_info_trigger_card_grade_up_by_trigger_id"`
	UserInfoTriggerMemberGuildSupportItemExpiredByTriggerId                    generic.ObjectByObjectIdList[model.TriggerMemberGuildSupportItemExpired]  `json:"user_info_trigger_member_guild_support_item_expired_by_trigger_id"` // TODO: not properly handled
	UserInfoTriggerMemberLoveLevelUpByTriggerId                                generic.ObjectByObjectIdList[model.TriggerMemberLoveLevelUp]              `json:"user_info_trigger_member_love_level_up_by_trigger_id"`
	UserAccessoryByUserAccessoryId                                             generic.ObjectByObjectIdList[model.UserAccessory]                         `json:"user_accessory_by_user_accessory_id"`
	UserAccessoryLevelUpItemById                                               generic.ObjectByObjectIdList[model.AccessoryLevelUpItem]                  `json:"user_accessory_level_up_item_by_id"`
	UserAccessoryRarityUpItemById                                              generic.ObjectByObjectIdList[model.AccessoryRarityUpItem]                 `json:"user_accessory_rarity_up_item_by_id"`
	UserUnlockScenesByEnum                                                     generic.ObjectByObjectIdList[model.UserUnlockScene]                       `json:"user_unlock_scenes_by_enum"`
	UserSceneTipsByEnum                                                        generic.ObjectByObjectIdList[model.UserSceneTips]                         `json:"user_scene_tips_by_enum"`
	UserRuleDescriptionById                                                    generic.ObjectByObjectIdList[model.UserRuleDescription]                   `json:"user_rule_description_by_id"` // TODO: not properly handled
	UserExchangeEventPointById                                                 generic.ObjectByObjectIdList[model.ExchangeEventPoint]                    `json:"user_exchange_event_point_by_id"`
	UserSchoolIdolFestivalIdRewardMissionById                                  generic.ObjectByObjectIdList[model.UserSchoolIdolFestivalIdRewardMission] `json:"user_school_idol_festival_id_reward_mission_by_id"` // not handled
	UserGpsPresentReceivedById                                                 generic.ObjectByObjectIdList[model.UserGpsPresentReceived]                `json:"user_gps_present_received_by_id"`                   // not handled
	UserEventMarathonByEventMasterId                                           generic.ObjectByObjectIdList[model.UserEventMarathon]                     `json:"user_event_marathon_by_event_master_id"`            // TODO: not properly handled
	UserEventMiningByEventMasterId                                             generic.ObjectByObjectIdList[model.UserEventMining]                       `json:"user_event_mining_by_event_master_id"`              // TODO: not properly handled
	UserEventCoopByEventMasterId                                               generic.ObjectByObjectIdList[model.UserEventCoop]                         `json:"user_event_coop_by_event_master_id"`                // TODO: not properly handled
	UserLiveSkipTicketById                                                     generic.ObjectByObjectIdList[model.LiveSkipTicket]                        `json:"user_live_skip_ticket_by_id"`
	UserStoryEventUnlockItemById                                               generic.ObjectByObjectIdList[model.StoryEventUnlockItem]                  `json:"user_story_event_unlock_item_by_id"`
	UserEventMarathonBoosterById                                               generic.ObjectByObjectIdList[model.EventMarathonBooster]                  `json:"user_event_marathon_booster_by_id"` // TODO: not properly handled
	UserReferenceBookById                                                      generic.ObjectByObjectIdList[model.UserReferenceBook]                     `json:"user_reference_book_by_id"`
	UserReviewRequestProcessFlowById                                           generic.ObjectByObjectIdList[model.UserReviewRequestProcessFlow]          `json:"user_review_request_process_flow_by_id"`                                                    // TODO: not properly handled
	UserTowerByTowerId                                                         generic.ObjectByObjectIdList[model.UserTower]                             `json:"user_tower_by_tower_id"`                                                                    // TODO: not properly handled
	UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterId generic.ObjectByObjectIdList[model.RecoveryTowerCardUsedCountItem]        `json:"user_recovery_tower_card_used_count_item_by_recovery_tower_card_used_count_item_master_id"` // TODO: not properly handled
	UserStoryLinkageById                                                       generic.ObjectByObjectIdList[model.UserStoryLinkage]                      `json:"user_story_linkage_by_id"`
	UserSubscriptionStatusById                                                 generic.ObjectByObjectIdList[model.UserSubscriptionStatus]                `json:"user_subscription_status_by_id"` // TODO: not properly handled
	UserStoryMainPartDigestMovieById                                           generic.ObjectByObjectIdList[model.UserStoryMainPartDigestMovie]          `json:"user_story_main_part_digest_movie_by_id"`
	UserMemberGuildById                                                        generic.ObjectByObjectIdList[model.UserMemberGuild]                       `json:"user_member_guild_by_id"`                // TODO: not properly handled
	UserMemberGuildSupportItemById                                             generic.ObjectByObjectIdList[model.UserMemberGuildSupportItem]            `json:"user_member_guild_support_item_by_id"`   // TODO: not properly handled
	UserDailyTheaterByDailyTheaterId                                           generic.ObjectByObjectIdList[model.UserDailyTheater]                      `json:"user_daily_theater_by_daily_theater_id"` // TODO: not properly handled
	UserPlayListById                                                           generic.ObjectByObjectIdList[model.UserPlayListItem]                      `json:"user_play_list_by_id"`
	UserSetProfileById                                                         generic.ObjectByObjectIdList[model.UserSetProfile]                        `json:"user_set_profile_by_id"`
	UserSteadyVoltageRankingById                                               generic.ObjectByObjectIdList[model.UserSteadyVoltageRanking]              `json:"user_steady_voltage_ranking_by_id"` // TODO: not properly handled
	UserSif2DataLinkById                                                       generic.ObjectByObjectIdList[model.UserSif2DataLink]                      `json:"user_sif_2_data_link_by_id"`        // TODO: not properly handled
}

func (userModel *UserModel) SetUserId(userId int) {
	rUserModel := reflect.ValueOf(userModel)
	for i := 0; i < rUserModel.Elem().NumField(); i++ {
		rSetUserId := reflect.Indirect(rUserModel).Field(i).Addr().MethodByName("SetUserId")
		if rSetUserId.IsValid() {
			rSetUserId.Call([]reflect.Value{reflect.ValueOf(userId)})
		} else {
			rUserId := reflect.Indirect(rUserModel).Field(i).FieldByName("UserId")
			if rUserId.IsValid() {
				rUserId.Set(reflect.ValueOf(userId))
			} else {
				fmt.Println("skipping field: ", reflect.Indirect(rUserModel).Field(i))
			}
		}
	}
}
