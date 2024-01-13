package client

import (
	"elichika/generic"
)

type UserInfoTrigger struct {
	UserInfoTriggerGachacPointExchangeRows          generic.List[UserInfoTriggerGachaPointExchangeRow]           `json:"user_info_trigger_gacha_point_exchange_rows"`
	UserInfoTriggerExpiredGiftBoxRows               generic.List[UserInfoTriggerExpiredGiftBox]                  `json:"user_info_trigger_expired_gift_box_rows"`
	UserInfoTriggerEventMarathonShowResultRows      generic.List[UserInfoTriggerEventMarathonShowResultRow]      `json:"user_info_trigger_event_marathon_show_result_rows"`
	UserInfoTriggerEventMiningShowResultRows        generic.List[UserInfoTriggerEventMiningShowResultRow]        `json:"user_info_trigger_event_mining_show_result_rows"`
	UserInfoTriggerEventCoopShowResultRows          generic.List[UserInfoTriggerEventCoopShowResultRow]          `json:"user_info_trigger_event_coop_show_result_rows"`
	UserInfoTriggerSubscriptionTrialEndRows         generic.List[UserInfoTriggerSubscriptionTrialEndRow]         `json:"user_info_trigger_subscription_trial_end_rows"`
	UserInfoTriggerSubscriptionEndRows              generic.List[UserInfoTriggerSubscriptionEndRow]              `json:"user_info_trigger_subscription_end_rows"`
	UserInfoTriggerMemberGuildRankingShowResultRows generic.List[UserInfoTriggerMemberGuildRankingShowResultRow] `json:"user_info_trigger_member_guild_ranking_show_result_rows"`
}
