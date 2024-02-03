package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchBootstrapResponse struct {
	UserModelDiff                      *client.UserModel                             `json:"user_model_diff"`
	UserInfoTrigger                    client.UserInfoTrigger                        `json:"user_info_trigger"`                     // TODO(not_really_handled)
	BillingStateInfo                   client.BillingStateInfo                       `json:"billing_state_info"`                    // TODO(not_really_handled)
	FetchBootstrapBannerResponse       client.BootstrapBanner                        `json:"fetch_bootstrap_banner_response"`       // TODO(not_really_handled)
	FetchBootstrapNewBadgeResponse     generic.Nullable[client.BootstrapNewBadge]    `json:"fetch_bootstrap_new_badge_response"`    // TODO(not_really_handled)
	FetchBootstrapPickupInfoResponse   client.BootstrapPickupInfo                    `json:"fetch_bootstrap_pickup_info_response"`  // TODO(not_really_handled)
	FetchBootstrapExpiredItemResponse  generic.Nullable[client.BootstrapExpiredItem] `json:"fetch_bootstrap_expired_item_response"` // TODO(not_really_handled)
	FetchBootstrapLoginBonusResponse   generic.Nullable[client.BootstrapLoginBonus]  `json:"fetch_bootstrap_login_bonus_response"`
	FetchBootstrapNoticeResponse       generic.Nullable[client.BootstrapNotice]      `json:"fetch_bootstrap_notice_response"`       // TODO(not_really_handled)
	FetchBootstrapSubscriptionResponse client.BootstrapSubscription                  `json:"fetch_bootstrap_subscription_response"` // TODO(not_really_handled)
	MissionBeginnerMasterId            generic.Nullable[int32]                       `json:"mission_beginner_master_id"`
	ShowChallengeBeginnerButton        bool                                          `json:"show_challenge_beginner_button"`
	ChallengeBeginnerCompletedIds      generic.List[int32]                           `json:"challenge_beginner_completed_ids"`
}
