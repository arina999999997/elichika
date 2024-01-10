package client

import (
	"elichika/generic"
)

type UserStatus struct {
	// TODO(refactor): move password to another table
	PassWord string `xorm:"'pass_word'" json:"-"`

	Name                                      LocalizedText           `xorm:"'name'" json:"name"`                                         // player name
	Nickname                                  LocalizedText           `xorm:"'nickname'" json:"nickname"`                                 // nickname in story
	LastLoginAt                               int64                   `json:"last_login_at"`                                              // in unix second
	Rank                                      int32                   `json:"rank"`                                                       // rank
	Exp                                       int32                   `json:"exp"`                                                        // total exp
	Message                                   LocalizedText           `xorm:"'message'" json:"message"`                                   // introduction message
	RecommendCardMasterId                     int32                   `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"` // featured / partner card
	MaxFriendNum                              int32                   `json:"max_friend_num"`
	LivePointFullAt                           int64                   `json:"live_point_full_at"`                              // in unix second
	LivePointBroken                           int32                   `json:"live_point_broken"`                               // max LP
	LivePointSubscriptionRecoveryDailyCount   int32                   `json:"live_point_subscription_recovery_daily_count"`    // membership LP 1 or 0?
	LivePointSubscriptionRecoveryDailyResetAt int64                   `json:"live_point_subscription_recovery_daily_reset_at"` // in unix second
	ActivityPointCount                        int32                   `json:"activity_point_count"`
	ActivityPointResetAt                      int64                   `json:"activity_point_reset_at"`                        // in unix time
	ActivityPointPaymentRecoveryDailyCount    int32                   `json:"activity_point_payment_recovery_daily_count"`    // how many AP recover used today?
	ActivityPointPaymentRecoveryDailyResetAt  int64                   `json:"activity_point_payment_recovery_daily_reset_at"` // in unix time 1684854000?
	GameMoney                                 int32                   `json:"game_money"`                                     // Gold
	CardExp                                   int32                   `json:"card_exp"`                                       // exp currency
	FreeSnsCoin                               int32                   `json:"free_sns_coin"`                                  // free gem
	AppleSnsCoin                              int32                   `json:"apple_sns_coin"`                                 // paid gem ios
	GoogleSnsCoin                             int32                   `json:"google_sns_coin"`                                // paid gem android
	SubscriptionCoin                          int32                   `json:"subscription_coin"`                              // member coin
	BirthDate                                 generic.Nullable[int32] `xorm:"json" json:"birth_date"`                         // yyyymm ?
	BirthMonth                                generic.Nullable[int32] `xorm:"json" json:"birth_month"`
	BirthDay                                  generic.Nullable[int32] `xorm:"json" json:"birth_day"`
	LatestLiveDeckId                          int32                   `xorm:"'latest_live_deck_id'" json:"latest_live_deck_id"` // last used live formation
	MainLessonDeckId                          int32                   `xorm:"'main_lesson_deck_id'" json:"main_lesson_deck_id"` // last used training formation
	FavoriteMemberId                          int32                   `xorm:"'favorite_member_id'" json:"favorite_member_id"`   // partner id
	LastLiveDifficultyId                      int32                   `xorm:"'last_live_difficulty_id'" json:"last_live_difficulty_id"`
	LpMagnification                           int32                   `json:"lp_magnification"`                                                                      // unused feature, always 1
	EmblemId                                  int32                   `xorm:"'emblem_id' "json:"emblem_id"`                                                          // title
	DeviceToken                               string                  `json:"device_token"`                                                                          //  some sort of salted encryption?, used to prevent using multiple device
	TutorialPhase                             int32                   `json:"tutorial_phase"`                                                                        // 99 = done
	TutorialEndAt                             int64                   `json:"tutorial_end_at"`                                                                       // in unix second
	LoginDays                                 int32                   `json:"login_days"`                                                                            // amount of days logged in
	NaviTapCount                              int32                   `json:"navi_tap_count"`                                                                        // number of partner tap that will increase bond
	NaviTapRecoverAt                          int64                   `json:"navi_tap_recover_at"`                                                                   // in unix time
	IsAutoMode                                bool                    `json:"is_auto_mode"`                                                                          // is autoplay enabled, for restarting app without finishing lives
	MaxScoreLiveDifficultyMasterId            generic.Nullable[int32] `xorm:"json 'max_score_live_difficulty_master_id'" json:"max_score_live_difficulty_master_id"` // not the one featured in profile?
	LiveMaxScore                              int32                   `json:"live_max_score"`
	MaxComboLiveDifficultyMasterId            generic.Nullable[int32] `xorm:"json 'max_combo_live_difficulty_master_id'" json:"max_combo_live_difficulty_master_id"`
	LiveMaxCombo                              int32                   `json:"live_max_combo"`
	LessonResumeStatus                        int32                   `json:"lesson_resume_status"`                                                      // for quitting while training, the number probably mean the phase of the training
	AccessoryBoxAdditional                    int32                   `json:"accessory_box_additional"`                                                  // additional accessory slot, max is 400 in official
	TermsOfUseVersion                         int32                   `json:"terms_of_use_version"`                                                      // 3 mean nothing to accept
	BootstrapSifidCheckAt                     int64                   `json:"bootstrap_sifid_check_at"`                                                  // not really what it sound like, probably safe to ignore it
	GdprVersion                               int32                   `json:"gdpr_version"`                                                              // set to 0 for outside the EU
	MemberGuildMemberMasterId                 generic.Nullable[int32] `xorm:"json 'member_guild_member_master_id'" json:"member_guild_member_master_id"` // member id of the channel joined
	MemberGuildLastUpdatedAt                  int64                   `json:"member_guild_last_updated_at"`                                              // unix time stamp, last time joining the channel (used to allow only 1 channel changing)
}
