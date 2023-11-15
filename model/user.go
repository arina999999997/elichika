package model

type UserStatus struct {
	UserID   int    `xorm:"pk 'user_id'" json:"-"`
	PassWord string `xorm:"'pass_word'" json:"-"`
	Name     struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // player name
	Nickname struct {
		DotUnderText string `xorm:"nickname" json:"dot_under_text"`
	} `xorm:"extends" json:"nickname"` // nickname in story
	LastLoginAt int64 `json:"last_login_at"` // in unix second
	Rank        int   `json:"rank"`          // rank
	Exp         int   `json:"exp"`           // total exp
	Message     struct {
		DotUnderText string `xorm:"message" json:"dot_under_text"`
	} `xorm:"extends" json:"message"` // introduction message
	RecommendCardMasterID                     int    `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"` // featured / partner card
	MaxFriendNum                              int    `json:"max_friend_num"`
	LivePointFullAt                           int64  `json:"live_point_full_at"`                              // in unix second
	LivePointBroken                           int    `json:"live_point_broken"`                               // max LP
	LivePointSubscriptionRecoveryDailyCount   int    `json:"live_point_subscription_recovery_daily_count"`    // membership LP 1 or 0?
	LivePointSubscriptionRecoveryDailyResetAt int    `json:"live_point_subscription_recovery_daily_reset_at"` // in unix second
	ActivityPointCount                        int    `json:"activity_point_count"`
	ActivityPointResetAt                      int    `json:"activity_point_reset_at"`                        // in unix time
	ActivityPointPaymentRecoveryDailyCount    int    `json:"activity_point_payment_recovery_daily_count"`    // how many AP recover used today?
	ActivityPointPaymentRecoveryDailyResetAt  int    `json:"activity_point_payment_recovery_daily_reset_at"` // in unix time 1684854000?
	GameMoney                                 int64  `json:"game_money"`                                     // Gold
	CardExp                                   int64  `json:"card_exp"`                                       // exp currency
	FreeSnsCoin                               int    `json:"free_sns_coin"`                                  // free gem
	AppleSnsCoin                              int    `json:"apple_sns_coin"`                                 // paid gem ios
	GoogleSnsCoin                             int    `json:"google_sns_coin"`                                // paid gem android
	SubscriptionCoin                          int    `json:"subscription_coin"`                              // member coin
	BirthDate                                 int    `json:"birth_date"`                                     // yyyymm ?
	BirthMonth                                int    `json:"birth_month"`
	BirthDay                                  int    `json:"birth_day"`
	LatestLiveDeckID                          int    `xorm:"'latest_live_deck_id'" json:"latest_live_deck_id"` // last used live formation
	MainLessonDeckID                          int    `xorm:"'main_lesson_deck_id'" json:"main_lesson_deck_id"` // last used training formation
	FavoriteMemberID                          int    `xorm:"'favorite_member_id'" json:"favorite_member_id"`   // partner id
	LastLiveDifficultyID                      int    `xorm:"'last_live_difficulty_id'" json:"last_live_difficulty_id"`
	LpMagnification                           int    `json:"lp_magnification"`                                                                 // unused feature, always 1
	EmblemID                                  int    `xorm:"'emblem_id' "json:"emblem_id"`                                                     // title
	DeviceToken                               string `json:"device_token"`                                                                     //  some sort of salted encryption?, used to prevent using multiple device
	TutorialPhase                             int    `json:"tutorial_phase"`                                                                   // 99 = done
	TutorialEndAt                             int    `json:"tutorial_end_at"`                                                                  // in unix second
	LoginDays                                 int    `json:"login_days"`                                                                       // amount of days logged in
	NaviTapCount                              int    `json:"navi_tap_count"`                                                                   // number of partner tap that will increase bond
	NaviTapRecoverAt                          int    `json:"navi_tap_recover_at"`                                                              // in unix time
	IsAutoMode                                bool   `json:"is_auto_mode"`                                                                     // is autoplay enabled, for restarting app without finishing lives
	MaxScoreLiveDifficultyMasterID            int    `xorm:"'max_score_live_difficulty_master_id'" json:"max_score_live_difficulty_master_id"` // not the one featured in profile?
	LiveMaxScore                              int    `json:"live_max_score"`
	MaxComboLiveDifficultyMasterID            int    `xorm:"'max_combo_live_difficulty_master_id'" json:"max_combo_live_difficulty_master_id"`
	LiveMaxCombo                              int    `json:"live_max_combo"`
	LessonResumeStatus                        int    `json:"lesson_resume_status"`                                                 // for quitting while training, the number probably mean the phase of the training
	AccessoryBoxAdditional                    int    `json:"accessory_box_additional"`                                             // additional accessory slot, max is 400 in official
	TermsOfUseVersion                         int    `json:"terms_of_use_version"`                                                 // 3 mean nothing to accept
	BootstrapSifidCheckAt                     int64  `json:"bootstrap_sifid_check_at"`                                             // not really what it sound like, probably safe to ignore it
	GdprVersion                               int    `json:"gdpr_version"`                                                         // set to 0 for outside the EU
	MemberGuildMemberMasterID                 int    `xorm:"'member_guild_member_master_id'" json:"member_guild_member_master_id"` // member id of the channel joined
	MemberGuildLastUpdatedAt                  int64  `json:"member_guild_last_updated_at"`                                         // unix time stamp, last time joining the channel (used to allow only 1 channel changing)
	Cash                                      int    `json:"cash"`
}

// this is not stored, constructed from main db
// partially loaded from u_info, then load from u_card
type UserBasicInfo struct {
	UserID int `xorm:"pk 'user_id'" json:"user_id"`
	Name   struct {
		DotUnderText string `xorm:"name" json:"dot_under_text"`
	} `xorm:"extends" json:"name"` // player name
	Rank                  int   `json:"rank"` // rank
	LastPlayedAt          int64 `xorm:"'last_login_at'" json:"last_played_at"`
	RecommendCardMasterID int   `xorm:"'recommend_card_master_id'" json:"recommend_card_master_id"` // featured / partner card

	RecommendCardLevel                  int  `xorm:"-" json:"recommend_card_level"`
	IsRecommendCardImageAwaken          bool `xorm:"-" json:"is_recommend_card_image_awaken"`
	IsRecommendCardAllTrainingActivated bool `xorm:"-" json:"is_recommend_card_all_training_activated"`

	EmblemID            int  `xorm:"'emblem_id' "json:"emblem_id"` // title
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
		MemberMasterID int `json:"member_master_id"`
		LovePoint      int `json:"love_point"`
	} `xorm:"-" json:"love_members"`
	MemberGuildMemberMasterID int `xorm:"'member_guild_member_master_id'" json:"member_guild_member_master_id"`
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
			LiveDifficultyMasterID int `xorm:"'max_score_live_difficulty_master_id'" json:"live_difficulty_master_id"`
			Score                  int `xorm:"'live_max_score'" json:"score"`
		} `xorm:"extends" json:"max_score_live_difficulty"`
		MaxComboLiveDifficulty struct {
			LiveDifficultyMasterID int `xorm:"'max_combo_live_difficulty_master_id'" json:"live_difficulty_master_id"`
			Score                  int `xorm:"'live_max_combo'" json:"score"`
		} `xorm:"extends" json:"max_combo_live_difficulty"`
	} `xorm:"extends" json:"play_info"`
	MemberInfo struct {
		UserMembers []MemberPublicInfo `json:"user_members"`
	} `xorm:"-" json:"member_info"`
}

type UserSetProfile struct {
	UserID                  int `xorm:"'user_id'" json:"-"`
	UserSetProfileID        int `xorm:"-" json:"user_set_profile_id"` // always 0
	VoltageLiveDifficultyID int `xorm:"'voltage_live_difficulty_id'" json:"voltage_live_difficulty_id"`
	ComboLiveDifficultyID   int `xorm:"'commbo_live_difficulty_id'" json:"commbo_live_difficulty_id"`
}

func (usp *UserSetProfile) ID() int64 {
	return int64(usp.UserSetProfileID)
}
