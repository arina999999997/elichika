package response

import (
	"elichika/client"
)

// TODO(temporary types): some types or file names are different from client's
type UserInfoTrigger struct {
	UserInfoTriggerGachaPointExchangeRows           []any `json:"user_info_trigger_gacha_point_exchange_rows"`
	UserInfoTriggerExpiredGiftBoxRows               []any `json:"user_info_trigger_expired_gift_box_rows"`
	UserInfoTriggerEventMarathonShowResultRows      []any `json:"user_info_trigger_event_marathon_show_result_rows"`
	UserInfoTriggerEventMiningShowResultRows        []any `json:"user_info_trigger_event_mining_show_result_rows"`
	UserInfoTriggerEventCoopShowResultRows          []any `json:"user_info_trigger_event_coop_show_result_rows"`
	UserInfoTriggerSubscriptionTrialEndRows         []any `json:"user_info_trigger_subscription_trial_end_rows"`
	UserInfoTriggerSubscriptionEndRows              []any `json:"user_info_trigger_subscription_end_rows"`
	UserInfoTriggerMemberGuildRankingShowResultRows []any `json:"user_info_trigger_member_guild_ranking_show_result_rows"`
}

type Banner struct {
	BannerMasterId       int                    `json:"banner_master_id"`
	BannerImageAssetPath client.TextureStruktur `json:"banner_image_asset_path"`
	BannerType           int                    `json:"banner_type"`
	ExpireAt             int64                  `json:"expire_at"`
	TransitionId         int                    `json:"transition_id"`
	TransitionParameter  int                    `json:"transition_parameter"`
}

type BootstrapBanner struct {
	Banners []Banner `json:"banners"`
}

type BootstrapNewBadge struct {
	IsNewMainStory                     bool  `json:"is_new_main_story"`
	UnreceivedPresentBox               int   `json:"unreceived_present_box"`
	IsUnreceivedPresentBoxSubscription bool  `json:"is_unreceived_present_box_subscription"`
	NoticeNewArrivalsIds               []int `json:"notice_new_arrivals_ids"`
	IsUpdateFriend                     bool  `json:"is_update_friend"`
	UnreceivedMission                  int   `json:"unreceived_mission"`
	UnreceivedChallengeBeginner        int   `json:"unreceived_challenge_beginner"`
	DailyTheaterTodayId                int   `json:"daily_theater_today_id"`
}

type LiveCampaignInfo struct {
	LiveCampaignEndAt          *int64 `json:"live_campaign_end_at"`
	LiveDailyCampaignEndAt     *int64 `json:"live_daily_campaign_end_at"`
	LiveExtraCampaignEndAt     *int64 `json:"live_extra_campaign_end_at"`
	LiveChallengeCampaignEndAt *int64 `json:"live_challenge_campaign_end_at"`
}

type BootstrapPickupInfo struct {
	ActiveEvent      *any                     `json:"active_event"`
	LiveCampaignInfo LiveCampaignInfo         `json:"live_campaign_info"`
	IsLessonCampaign bool                     `json:"is_lesson_campaign"`
	AppealGachas     []client.TextureStruktur `json:"appeal_gachas"`
	IsShopSale       bool                     `json:"is_shop_sale"`
	IsSnsCoinSale    bool                     `json:"is_sns_coin_sale"`
}

type BootstrapExpiredItem struct {
	ExpiredItems []client.Content `json:"expired_items"`
}

type BootstrapNotice struct {
	SuperNotices        []any `json:"super_notices"`
	FetchedAt           int64 `json:"fetched_at"`
	ReviewSuperNoticeAt int64 `json:"review_super_notice_at"`
	ForceViewNoticeIds  []any `json:"force_view_notice_ids"`
}

type BootstrapSubscription struct {
	ContinueRewards []any `json:"continue_rewards"`
}

type FetchBootstrapResponse struct {
	UserModelDiff                      *UserModel                 `json:"user_model_diff"`
	UserInfoTrigger                    UserInfoTrigger            `json:"user_info_trigger"`                     // TODO(not_really_handled)
	BillingStateInfo                   BillingStateInfo           `json:"billing_state_info"`                    // TODO(not_really_handled)
	FetchBootstrapBannerResponse       BootstrapBanner            `json:"fetch_bootstrap_banner_response"`       // TODO(not_really_handled)
	FetchBootstrapNewBadgeResponse     BootstrapNewBadge          `json:"fetch_bootstrap_new_badge_response"`    // TODO(not_really_handled)
	FetchBootstrapPickupInfoResponse   BootstrapPickupInfo        `json:"fetch_bootstrap_pickup_info_response"`  // TODO(not_really_handled)
	FetchBootstrapExpiredItemResponse  BootstrapExpiredItem       `json:"fetch_bootstrap_expired_item_response"` // TODO(not_really_handled)
	FetchBootstrapLoginBonusResponse   client.BootstrapLoginBonus `json:"fetch_bootstrap_login_bonus_response"`
	FetchBootstrapNoticeResponse       BootstrapNotice            `json:"fetch_bootstrap_notice_response"`       // TODO(not_really_handled)
	FetchBootstrapSubscriptionResponse BootstrapSubscription      `json:"fetch_bootstrap_subscription_response"` // TODO(not_really_handled)
	MissionBeginnerMasterId            *int                       `json:"mission_beginner_master_id"`
	ShowChallengeBeginnerButton        bool                       `json:"show_challenge_beginner_button"`
	ChallengeBeginnerCompletedIds      []int                      `json:"challenge_beginner_completed_ids"`
}
