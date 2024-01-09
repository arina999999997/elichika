package model

import (
	"elichika/client"
	"elichika/generic"
)

type UserStatus struct {
	PassWord string `xorm:"'pass_word'" json:"-"`

	Name                                      client.LocalizedText `xorm:"'name'" json:"name"`                                         // player name
	Nickname                                  client.LocalizedText `xorm:"'nickname'" json:"nickname"`                                 // nickname in story
	LastLoginAt                               int64                `json:"last_login_at"`                                              // in unix second
	Rank                                      int32                `json:"rank"`                                                       // rank
	Exp                                       int32                `json:"exp"`                                                        // total exp
	Message                                   client.LocalizedText `xorm:"'message'" json:"message"`                                   // introduction message
	RecommendCardMasterId                     int32                `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"` // featured / partner card
	MaxFriendNum                              int32                `json:"max_friend_num"`
	LivePointFullAt                           int64                `json:"live_point_full_at"`                              // in unix second
	LivePointBroken                           int32                `json:"live_point_broken"`                               // max LP
	LivePointSubscriptionRecoveryDailyCount   int32                `json:"live_point_subscription_recovery_daily_count"`    // membership LP 1 or 0?
	LivePointSubscriptionRecoveryDailyResetAt int64                `json:"live_point_subscription_recovery_daily_reset_at"` // in unix second
	ActivityPointCount                        int32                `json:"activity_point_count"`
	ActivityPointResetAt                      int64                `json:"activity_point_reset_at"`                        // in unix time
	ActivityPointPaymentRecoveryDailyCount    int32                `json:"activity_point_payment_recovery_daily_count"`    // how many AP recover used today?
	ActivityPointPaymentRecoveryDailyResetAt  int32                `json:"activity_point_payment_recovery_daily_reset_at"` // in unix time 1684854000?
	GameMoney                                 int32                `json:"game_money"`                                     // Gold
	CardExp                                   int32                `json:"card_exp"`                                       // exp currency
	FreeSnsCoin                               int32                `json:"free_sns_coin"`                                  // free gem
	AppleSnsCoin                              int32                `json:"apple_sns_coin"`                                 // paid gem ios
	GoogleSnsCoin                             int32                `json:"google_sns_coin"`                                // paid gem android
	SubscriptionCoin                          int32                `json:"subscription_coin"`                              // member coin
	BirthDate                                 *int32               `json:"birth_date"`                                     // yyyymm ?
	BirthMonth                                *int32               `json:"birth_month"`
	BirthDay                                  *int32               `json:"birth_day"`
	LatestLiveDeckId                          int32                `xorm:"'latest_live_deck_id'" json:"latest_live_deck_id"` // last used live formation
	MainLessonDeckId                          int32                `xorm:"'main_lesson_deck_id'" json:"main_lesson_deck_id"` // last used training formation
	FavoriteMemberId                          int32                `xorm:"'favorite_member_id'" json:"favorite_member_id"`   // partner id
	LastLiveDifficultyId                      int32                `xorm:"'last_live_difficulty_id'" json:"last_live_difficulty_id"`
	LpMagnification                           int32                `json:"lp_magnification"`                                                                 // unused feature, always 1
	EmblemId                                  int32                `xorm:"'emblem_id' "json:"emblem_id"`                                                     // title
	DeviceToken                               string               `json:"device_token"`                                                                     //  some sort of salted encryption?, used to prevent using multiple device
	TutorialPhase                             int32                `json:"tutorial_phase"`                                                                   // 99 = done
	TutorialEndAt                             int64                `json:"tutorial_end_at"`                                                                  // in unix second
	LoginDays                                 int32                `json:"login_days"`                                                                       // amount of days logged in
	NaviTapCount                              int32                `json:"navi_tap_count"`                                                                   // number of partner tap that will increase bond
	NaviTapRecoverAt                          int64                `json:"navi_tap_recover_at"`                                                              // in unix time
	IsAutoMode                                bool                 `json:"is_auto_mode"`                                                                     // is autoplay enabled, for restarting app without finishing lives
	MaxScoreLiveDifficultyMasterId            *int32               `xorm:"'max_score_live_difficulty_master_id'" json:"max_score_live_difficulty_master_id"` // not the one featured in profile?
	LiveMaxScore                              int32                `json:"live_max_score"`
	MaxComboLiveDifficultyMasterId            *int32               `xorm:"'max_combo_live_difficulty_master_id'" json:"max_combo_live_difficulty_master_id"`
	LiveMaxCombo                              int32                `json:"live_max_combo"`
	LessonResumeStatus                        int32                `json:"lesson_resume_status"`                                                 // for quitting while training, the number probably mean the phase of the training
	AccessoryBoxAdditional                    int32                `json:"accessory_box_additional"`                                             // additional accessory slot, max is 400 in official
	TermsOfUseVersion                         int32                `json:"terms_of_use_version"`                                                 // 3 mean nothing to accept
	BootstrapSifidCheckAt                     int64                `json:"bootstrap_sifid_check_at"`                                             // not really what it sound like, probably safe to ignore it
	GdprVersion                               int32                `json:"gdpr_version"`                                                         // set to 0 for outside the EU
	MemberGuildMemberMasterId                 *int32               `xorm:"'member_guild_member_master_id'" json:"member_guild_member_master_id"` // member id of the channel joined
	MemberGuildLastUpdatedAt                  int64                `json:"member_guild_last_updated_at"`                                         // unix time stamp, last time joining the channel (used to allow only 1 channel changing)
}

// this is not stored, constructed from main db
// partially loaded from u_info, then load from u_card
type UserBasicInfo struct {
	UserId int `xorm:"pk 'user_id'" json:"user_id"`
	Name   struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // player name
	Rank                  int   `json:"rank"` // rank
	LastPlayedAt          int64 `xorm:"'last_login_at'" json:"last_played_at"`
	RecommendCardMasterId int   `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"` // featured / partner card

	RecommendCardLevel                  int  `xorm:"-" json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool `xorm:"-" json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool `xorm:"-" json:"is_recommend_card_all_training_activated"`

	EmblemId            int  `xorm:"'emblem_id' "json:"emblem_id"` // title
	IsNew               bool `xorm:"-" json:"is_new"`              // not sure what this thing is about, maybe new friend?
	IntroductionMessage struct {
		DotUnderText string `xorm:"message" json:"dot_under_text"`
	} `xorm:"extends" json:"introduction_message"` // introduction message
	FriendApprovedAt *int64 `xorm:"-" json:"friend_approved_at"`
	RequestStatus    int    `xorm:"-" json:"request_status"`
	IsRequestPending bool   `xorm:"-" json:"is_request_pending"`
}

type UserProfileLiveStats struct {
	LivePlayCount  [5]int `xorm:"'live_play_count'"`
	LiveClearCount [5]int `xorm:"'live_clear_count'"`
}

type UserProfileInfo struct {
	BasicInfo      UserBasicInfo `xorm:"extends" json:"basic_info"`
	TotalLovePoint int           `xorm:"-" json:"total_love_point"`
	LoveMembers    [3]struct {
		MemberMasterId int `json:"member_master_id"`
		LovePoint      int `json:"love_point"`
	} `xorm:"-" json:"love_members"`
	MemberGuildMemberMasterId int `xorm:"'member_guild_member_master_id'" json:"member_guild_member_master_id"`
}

type LivePartnerCard struct {
	LivePartnerCategoryMasterId int             `json:"live_partner_category_master_id"`
	PartnerCard                 PartnerCardInfo `json:"partner_card"`
}

type Profile struct {
	ProfileInfo UserProfileInfo `xorm:"extends" json:"profile_info"`
	GuestInfo   struct {
		LivePartnersCards []LivePartnerCard `json:"live_partner_cards"`
	} `xorm:"-" json:"guest_info"`
	PlayInfo struct {
		LivePlayCount          []int          `xorm:"-" json:"live_play_count"`
		LiveClearCount         []int          `xorm:"-" json:"live_clear_count"`
		JoinedLiveCardRanking  []CardPlayInfo `xorm:"-" json:"joined_live_card_ranking"`
		PlaySkillCardRanking   []CardPlayInfo `xorm:"-" json:"play_skill_card_ranking"`
		MaxScoreLiveDifficulty struct {
			LiveDifficultyMasterId int32 `xorm:"'max_score_live_difficulty_master_id'" json:"live_difficulty_master_id"`
			Score                  int32 `xorm:"'live_max_score'" json:"score"`
		} `xorm:"extends" json:"max_score_live_difficulty"`
		MaxComboLiveDifficulty struct {
			LiveDifficultyMasterId int32 `xorm:"'max_combo_live_difficulty_master_id'" json:"live_difficulty_master_id"`
			Score                  int32 `xorm:"'live_max_combo'" json:"score"`
		} `xorm:"extends" json:"max_combo_live_difficulty"`
	} `xorm:"extends" json:"play_info"`
	MemberInfo struct {
		UserMembers []MemberPublicInfo `json:"user_members"`
	} `xorm:"-" json:"member_info"`
}

func init() {

	type DbUser struct {
		UserStatus           `xorm:"extends"`
		UserProfileLiveStats `xorm:"extends"`
	}
	TableNameToInterface["u_info"] = generic.UserIdWrapper[DbUser]{}


}
