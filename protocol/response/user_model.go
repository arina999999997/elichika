package response

import (
	"elichika/generic"
	"elichika/model"

	"fmt"
	"reflect"
)

// user_model and user_model_diff both use this, just that the outer key is different
// might be possible to mod the client to only use one thing, but maybe there will be some part where it matter
type UserModel struct {
	UserStatus                                                                 model.UserStatus                                                          `json:"user_status"`
	UserMemberByMemberID                                                       generic.ObjectByObjectIDList[model.UserMember]                            `json:"user_member_by_member_id"`
	UserCardByCardID                                                           generic.ObjectByObjectIDList[model.UserCard]                              `json:"user_card_by_card_id"`
	UserSuitBySuitID                                                           generic.ObjectByObjectIDList[model.UserSuit]                              `json:"user_suit_by_suit_id"`
	UserLiveDeckByID                                                           generic.ObjectByObjectIDList[model.UserLiveDeck]                          `json:"user_live_deck_by_id"`
	UserLivePartyByID                                                          generic.ObjectByObjectIDList[model.UserLiveParty]                         `json:"user_live_party_by_id"`
	UserLessonDeckByID                                                         generic.ObjectByObjectIDList[model.UserLessonDeck]                        `json:"user_lesson_deck_by_id"`
	UserLiveMvDeckByID                                                         generic.ObjectByObjectIDList[model.UserLiveMvDeck]                        `json:"user_live_mv_deck_by_id"`
	UserLiveMvDeckCustomByID                                                   generic.ObjectByObjectIDList[model.UserLiveMvDeck]                        `json:"user_live_mv_deck_custom_by_id"`
	UserLiveDifficultyByDifficultyID                                           generic.ObjectByObjectIDList[model.UserLiveDifficulty]                    `json:"user_live_difficulty_by_difficulty_id"`
	UserStoryMainByStoryMainID                                                 generic.ObjectByObjectIDList[model.UserStoryMain]                         `json:"user_story_main_by_story_main_id"`
	UserStoryMainSelectedByStoryMainCellID                                     generic.ObjectByObjectIDList[model.UserStoryMainSelected]                 `json:"user_story_main_selected_by_story_main_cell_id"`
	UserVoiceByVoiceID                                                         generic.ObjectByObjectIDList[model.UserVoice]                             `json:"user_voice_by_voice_id"`
	UserEmblemByEmblemID                                                       generic.ObjectByObjectIDList[model.UserEmblem]                            `json:"user_emblem_by_emblem_id"` // TODO: not properly handled
	UserGachaTicketByTicketID                                                  generic.ObjectByObjectIDList[model.GachaTicket]                           `json:"user_gacha_ticket_by_ticket_id"`
	UserGachaPointByPointID                                                    generic.ObjectByObjectIDList[model.GachaPoint]                            `json:"user_gacha_point_by_point_id"`
	UserLessonEnhancingItemByItemID                                            generic.ObjectByObjectIDList[model.LessonEnhancingItem]                   `json:"user_lesson_enhancing_item_by_item_id"`
	UserTrainingMaterialByItemID                                               generic.ObjectByObjectIDList[model.TrainingMaterial]                      `json:"user_training_material_by_item_id"`
	UserGradeUpItemByItemID                                                    generic.ObjectByObjectIDList[model.GradeUpItem]                           `json:"user_grade_up_item_by_item_id"`
	UserCustomBackgroundByID                                                   generic.ObjectByObjectIDList[model.UserCustomBackground]                  `json:"user_custom_background_by_id"`
	UserStorySideByID                                                          generic.ObjectByObjectIDList[model.UserStorySide]                         `json:"user_story_side_by_id"`
	UserStoryMemberByID                                                        generic.ObjectByObjectIDList[model.UserStoryMember]                       `json:"user_story_member_by_id"`
	UserCommunicationMemberDetailBadgeByID                                     generic.ObjectByObjectIDList[model.UserCommunicationMemberDetailBadge]    `json:"user_communication_member_detail_badge_by_id"` // TODO: not properly handled
	UserStoryEventHistoryByID                                                  generic.ObjectByObjectIDList[model.UserStoryEventHistory]                 `json:"user_story_event_history_by_id"`
	UserRecoveryLpByID                                                         generic.ObjectByObjectIDList[model.RecoverLp]                             `json:"user_recovery_lp_by_id"`
	UserRecoveryApByID                                                         generic.ObjectByObjectIDList[model.RecoverAp]                             `json:"user_recovery_ap_by_id"`
	UserMissionByMissionID                                                     generic.ObjectByObjectIDList[model.UserMission]                           `json:"user_mission_by_mission_id"`        // TODO: not properly handled
	UserDailyMissionByMissionID                                                generic.ObjectByObjectIDList[model.UserDailyMission]                      `json:"user_daily_mission_by_mission_id"`  // TODO: not properly handled
	UserWeeklyMissionByMissionID                                               generic.ObjectByObjectIDList[model.UserWeeklyMission]                     `json:"user_weekly_mission_by_mission_id"` // TODO: not properly handled
	UserInfoTriggerBasicByTriggerID                                            generic.ObjectByObjectIDList[model.TriggerBasic]                          `json:"user_info_trigger_basic_by_trigger_id"`
	UserInfoTriggerCardGradeUpByTriggerID                                      generic.ObjectByObjectIDList[model.TriggerCardGradeUp]                    `json:"user_info_trigger_card_grade_up_by_trigger_id"`
	UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID                    generic.ObjectByObjectIDList[model.TriggerMemberGuildSupportItemExpired]  `json:"user_info_trigger_member_guild_support_item_expired_by_trigger_id"` // TODO: not properly handled
	UserInfoTriggerMemberLoveLevelUpByTriggerID                                generic.ObjectByObjectIDList[model.TriggerMemberLoveLevelUp]              `json:"user_info_trigger_member_love_level_up_by_trigger_id"`
	UserAccessoryByUserAccessoryID                                             generic.ObjectByObjectIDList[model.UserAccessory]                         `json:"user_accessory_by_user_accessory_id"`
	UserAccessoryLevelUpItemByID                                               generic.ObjectByObjectIDList[model.AccessoryLevelUpItem]                  `json:"user_accessory_level_up_item_by_id"`
	UserAccessoryRarityUpItemByID                                              generic.ObjectByObjectIDList[model.AccessoryRarityUpItem]                 `json:"user_accessory_rarity_up_item_by_id"`
	UserUnlockScenesByEnum                                                     generic.ObjectByObjectIDList[model.UserUnlockScene]                       `json:"user_unlock_scenes_by_enum"`
	UserSceneTipsByEnum                                                        generic.ObjectByObjectIDList[model.UserSceneTips]                         `json:"user_scene_tips_by_enum"`
	UserRuleDescriptionByID                                                    generic.ObjectByObjectIDList[model.UserRuleDescription]                   `json:"user_rule_description_by_id"` // TODO: not properly handled
	UserExchangeEventPointByID                                                 generic.ObjectByObjectIDList[model.ExchangeEventPoint]                    `json:"user_exchange_event_point_by_id"`
	UserSchoolIdolFestivalIDRewardMissionByID                                  generic.ObjectByObjectIDList[model.UserSchoolIdolFestivalIDRewardMission] `json:"user_school_idol_festival_id_reward_mission_by_id"` // not handled
	UserGpsPresentReceivedByID                                                 generic.ObjectByObjectIDList[model.UserGpsPresentReceived]                `json:"user_gps_present_received_by_id"`                   // not handled
	UserEventMarathonByEventMasterID                                           generic.ObjectByObjectIDList[model.UserEventMarathon]                     `json:"user_event_marathon_by_event_master_id"`            // TODO: not properly handled
	UserEventMiningByEventMasterID                                             generic.ObjectByObjectIDList[model.UserEventMining]                       `json:"user_event_mining_by_event_master_id"`              // TODO: not properly handled
	UserEventCoopByEventMasterID                                               generic.ObjectByObjectIDList[model.UserEventCoop]                         `json:"user_event_coop_by_event_master_id"`                // TODO: not properly handled
	UserLiveSkipTicketByID                                                     generic.ObjectByObjectIDList[model.LiveSkipTicket]                        `json:"user_live_skip_ticket_by_id"`
	UserStoryEventUnlockItemByID                                               generic.ObjectByObjectIDList[model.StoryEventUnlockItem]                  `json:"user_story_event_unlock_item_by_id"`
	UserEventMarathonBoosterByID                                               generic.ObjectByObjectIDList[model.EventMarathonBooster]                  `json:"user_event_marathon_booster_by_id"` // TODO: not properly handled
	UserReferenceBookByID                                                      generic.ObjectByObjectIDList[model.UserReferenceBook]                     `json:"user_reference_book_by_id"`
	UserReviewRequestProcessFlowByID                                           generic.ObjectByObjectIDList[model.UserReviewRequestProcessFlow]          `json:"user_review_request_process_flow_by_id"`                                                    // TODO: not properly handled
	UserTowerByTowerID                                                         generic.ObjectByObjectIDList[model.UserTower]                             `json:"user_tower_by_tower_id"`                                                                    // TODO: not properly handled
	UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterID generic.ObjectByObjectIDList[model.RecoveryTowerCardUsedCountItem]        `json:"user_recovery_tower_card_used_count_item_by_recovery_tower_card_used_count_item_master_id"` // TODO: not properly handled
	UserStoryLinkageByID                                                       generic.ObjectByObjectIDList[model.UserStoryLinkage]                      `json:"user_story_linkage_by_id"`
	UserSubscriptionStatusByID                                                 generic.ObjectByObjectIDList[model.UserSubscriptionStatus]                `json:"user_subscription_status_by_id"` // TODO: not properly handled
	UserStoryMainPartDigestMovieByID                                           generic.ObjectByObjectIDList[model.UserStoryMainPartDigestMovie]          `json:"user_story_main_part_digest_movie_by_id"`
	UserMemberGuildByID                                                        generic.ObjectByObjectIDList[model.UserMemberGuild]                       `json:"user_member_guild_by_id"`                // TODO: not properly handled
	UserMemberGuildSupportItemByID                                             generic.ObjectByObjectIDList[model.UserMemberGuildSupportItem]            `json:"user_member_guild_support_item_by_id"`   // TODO: not properly handled
	UserDailyTheaterByDailyTheaterID                                           generic.ObjectByObjectIDList[model.UserDailyTheater]                      `json:"user_daily_theater_by_daily_theater_id"` // TODO: not properly handled
	UserPlayListByID                                                           generic.ObjectByObjectIDList[model.UserPlayListItem]                      `json:"user_play_list_by_id"`
	UserSetProfileByID                                                         generic.ObjectByObjectIDList[model.UserSetProfile]                        `json:"user_set_profile_by_id"`
	UserSteadyVoltageRankingByID                                               generic.ObjectByObjectIDList[model.UserSteadyVoltageRanking]              `json:"user_steady_voltage_ranking_by_id"` // TODO: not properly handled
	UserSif2DataLinkByID                                                       generic.ObjectByObjectIDList[model.UserSif2DataLink]                      `json:"user_sif_2_data_link_by_id"`        // TODO: not properly handled
}

func (userModel *UserModel) SetUserID(uid int) {
	rUserModel := reflect.ValueOf(userModel)
	for i := 0; i < rUserModel.Elem().NumField(); i++ {
		rSetUserID := reflect.Indirect(rUserModel).Field(i).Addr().MethodByName("SetUserID")
		if rSetUserID.IsValid() {
			rSetUserID.Call([]reflect.Value{reflect.ValueOf(uid)})
		} else {
			rUserID := reflect.Indirect(rUserModel).Field(i).FieldByName("UserID")
			if rUserID.IsValid() {
				rUserID.Set(reflect.ValueOf(uid))
			} else {
				fmt.Println("skipping field: ", reflect.Indirect(rUserModel).Field(i))
			}
		}
	}
}
