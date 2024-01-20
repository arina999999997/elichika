package client

import (
	"elichika/generic"
)

type LiveResult struct {
	LiveDifficultyMasterId        int32                                                 `json:"live_difficulty_master_id"`
	LiveDeckId                    int32                                                 `json:"live_deck_id"`
	StandardDrops                 generic.Array[LiveDropContent]                        `json:"standard_drops"`
	AdditionalDrops               generic.Array[LiveDropContent]                        `json:"additional_drops"`
	GimmickDrops                  generic.Array[LiveDropContent]                        `json:"gimmick_drops"`
	MemberLoveStatuses            generic.Dictionary[int32, LiveResultMemberLoveStatus] `json:"member_love_statuses"`
	Mvp                           LiveResultMvp                                         `json:"mvp"`
	Partner                       generic.Nullable[OtherUser]                           `json:"partner"` // pointer
	LiveResultAchievements        generic.Dictionary[int32, LiveResultAchievement]      `json:"live_result_achievements"`
	LiveResultAchievementStatus   LiveResultAchievementStatus                           `json:"live_result_achievement_status"`
	Voltage                       int32                                                 `json:"voltage"`
	LastBestVoltage               int32                                                 `json:"last_best_voltage"`
	BeforeUserExp                 int32                                                 `json:"before_user_exp"`
	GainUserExp                   int32                                                 `json:"gain_user_exp"`
	IsRewardAccessoryInPresentBox bool                                                  `json:"is_reward_accessory_in_present_box"`
	ActiveEventResult             generic.Nullable[LiveResultActiveEvent]               `json:"active_event_result"`
	LiveResultTower               generic.Nullable[LiveResultTower]                     `json:"live_result_tower"`                          //pointer
	LiveResultMemberGuild         generic.Nullable[LiveResultMemberGuild]               `json:"live_result_member_guild"`                   //pointer
	LiveFinishStatus              int32                                                 `json:"live_finish_status" enum:"LiveFinishStatus"` //pointer
}
