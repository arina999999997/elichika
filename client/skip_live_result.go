package client

import (
	"elichika/generic"
)

type SkipLiveResult struct {
	LiveDifficultyMasterId        int32                                                 `json:"live_difficulty_master_id"`
	LiveDeckId                    int32                                                 `json:"live_deck_id"`
	Drops                         generic.Array[LiveResultContentPack]                  `json:"drops"`
	MemberLoveStatuses            generic.Dictionary[int32, LiveResultMemberLoveStatus] `json:"member_love_statuses"`
	GainUserExp                   int32                                                 `json:"gain_user_exp"`
	IsRewardAccessoryInPresentBox bool                                                  `json:"is_reward_accessory_in_present_box"`
	ActiveEventResult             generic.Nullable[LiveResultActiveEvent]               `json:"active_event_result"`
	LiveResultMemberGuild         generic.Nullable[LiveResultMemberGuild]               `json:"live_result_member_guild"`
}
