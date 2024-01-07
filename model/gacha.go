package model

type GachaAppeal struct { // s_gacha_appeal
	GachaAppealMasterId int `xorm:"pk 'gacha_appeal_master_id'" json:"-"`
	CardMasterId        int `xorm:"'card_master_id'" json:"card_master_id"`
	AppearanceType      int `json:"appearance_type"`
	MainImageAsset      struct {
		V string `xorm:"'main_image_asset'" json:"v"`
	} `xorm:"extends" json:"main_image_asset"`
	SubImageAsset struct {
		V *string `xorm:"'sub_image_asset'" json:"v"`
	} `xorm:"extends" json:"sub_image_asset"`
	TextImageAsset struct {
		V *string `xorm:"'text_image_asset'" json:"v"`
	} `xorm:"extends" json:"text_image_asset"`
}

type GachaDraw struct { // s_gacha_draw
	GachaMasterId        int    `xorm:"gacha_master_id" json:"-"` // client only send gacha_draw_master_id
	GachaDrawMasterId    int    `xorm:"pk 'gacha_draw_master_id'" json:"gacha_draw_master_id"`
	RecoveryType         int    `json:"recover_type"`
	RecoverAt            *int64 `json:"recover_at"`
	DrawCount            int    `json:"draw_count"`
	GachaPaymentType     int    `json:"gacha_payment_type"`
	GachaPaymentMasterId int    `xorm:"gacha_payment_master_id" json:"gacha_payment_master_id"`
	GachaPaymentAmount   int    `json:"gacha_payment_amount"`
	// doesn't seem to be used? maybe for paid festival?
	GachaPointAmount int `json:"gacha_point_amount"`
	Description      struct {
		DotUnderText string `xorm:"'description'" json:"dot_under_text"`
	} `xorm:"extends" json:"description"`
	IsBonus         bool `json:"is_bonus"`
	BonusAppealText struct {
		DotUnderText string `xorm:"'bonus_appeal_text'" json:"dot_under_text"`
	} `xorm:"extends" json:"bonus_appeal_text"`
	RetryCount           *int   `json:"retry_count"`
	GachaRetryMasterRows [0]int `json:"gacha_retry_master_rows"`
	DailyLimit           int    `json:"daily_limit"`    // 0 for unlimited
	DailyInterval        int    `json:"daily_interval"` // always 1
	TermLimit            int    `json:"term_limit"`     // 0 for unlimited
	// number of remaining pull in this day
	RemainDayCount *int `json:"remain_day_count"`
	// number of remaining pull in this banner
	RemainTermCount *int  `json:"remain_term_count"`
	PerformanceId   int   `xorm:"performance_id" json:"performance_id"` // gacha performance
	IsSubscription  bool  `json:"is_subscription"`                      // is subscription only gacha
	Guarantees      []int `xorm:"'guarantees'" json:"-"`                // ids to gacha.GachaGuarantee
}

type GachaDrawStepupDetail struct { // u_gacha_draw_stepup
	UserId        int  `json:"-"`
	GachaMasterId int  `json:"-"`
	CurrentStep   int  `json:"current_step"`
	LoopCount     int  `json:"loop_count"`
	MaxLoop       int  `json:"max_loop"`
	MaxStep       int  `json:"max_step"`
	IsMaxNextStep bool `json:"is_max_next_step"`
}

type Gacha struct { // s_gacha
	GachaMasterId int `xorm:"pk 'gacha_master_id'" json:"gacha_master_id"`
	GachaType     int `json:"gacha_type"`
	GachaDrawType int `json:"gacha_draw_type"`
	Title         struct {
		DotUnderText string `xorm:"'title'" json:"dot_under_text"`
	} `xorm:"extends" json:"title"`
	BannerImageAsset struct {
		// stand for visual?
		V string `xorm:"'banner_image_asset' "json:"v"`
	} `xorm:"extends" json:"banner_image_asset"`
	IsTimeLimited       bool                   `json:"is_time_limited"`
	EndAt               int64                  `json:"end_at"` // 1924873200
	PointMasterId       *int                   `xorm:"'point_master_id'" json:"point_master_id"`
	PointExchangeExpire *int64                 `json:"point_exchange_expire_at"`
	AppealAt            int64                  `json:"appeal_at"`
	NoticeId            int                    `xorm:"'notice_id'" json:"notice_id"`
	AppealView          int                    `json:"appeal_view"`
	GachaAppeals        []GachaAppeal          `xorm:"-" json:"gacha_appeals"`     // member details button
	DbGachaAppeals      []int                  `xorm:"'gacha_appeals'" json:"-"`   // member details button
	GachaDraws          []GachaDraw            `xorm:"-" json:"gacha_draws"`       // scout buttons
	DbGachaDraws        []int                  `xorm:"'gacha_draws'" json:"-"`     // scout buttons
	GachaDrawStepup     *GachaDrawStepupDetail `xorm:"-" json:"gacha_draw_stepup"` // not stored since this is user exlusive
	DbGachaDrawStepup   *int                   `xorm:"gacha_draw_stepup" json:"-"` // store the default stepup in db if there is one
	DbGachaGroups       []int                  `xorm:"gacha_groups" json:"-"`
}

type GachaDrawReq struct {
	GachaDrawMasterId int `json:"gacha_draw_master_id"`
	ButtonDrawCount   int `json:"button_draw_count"`
}

type ResultCard struct {
	GachaLotType         int      `json:"gacha_lot_type"` // 1 for normal, 2 for guaranteed
	CardMasterId         int      `json:"card_master_id"`
	Level                int      `json:"level"`                   // always 1
	BeforeGrade          int      `json:"before_grade"`            // 0 for new
	AfterGrade           int      `json:"after_grade"`             // 0 for new, 5 for maxed
	Content              *Content `json:"content"`                 // if maxed then award radiance. Technically we can award anything, but this is only the display value
	LimitExceeded        bool     `json:"limit_exceeded"`          // always false
	BeforeLoveLevelLimit int      `json:"before_love_level_limit"` // always correct
	AfterLoveLevelLimit  int      `json:"after_love_level_limit"`  // 0 for max level
}

// How this server implement (back end) gacha (for now):
// - Each banner need to implement a set of cards, each cards have group Id.
// - Each banner need to implement the weight of all the group Id.
// - To get a card from a pile, the system first roll the group based on the weight (empty group excluded).
// - Then the system choose equally randomly from within the chosen group.
// - Some common groups used in actual game is R, SR, UR, Featured UR and so on.
// - Multi-pull guaranteed cards are defined using a condition system.
// - Condition system is built using callable.
// - Each condition decide for exactly 1 card only.
// - The system roll the guaranteed cards first, then it roll all other card randomly.
// - This is more simple because randomly roll everything and check might take a long time.
// - And randomly choose and fix run into the problem of which card to replace.
// - If the guarantee function can't find a candidate, then the gurantee is forfeited and you get a random card instead.
// - The weight inside the guaranteed cards pile stay the same. I.e. Still choose the group, then choose random between group.
// - There is no random seeding effort, math/rand is used as is.

// different gacha can share groups
type GachaGroup struct { // s_gacha_group
	GroupMasterId int   `xorm:"pk 'group_master_id'"`
	GroupWeight   int64 `xorm:"'group_weight'"`
}

type GachaCard struct { // s_gacha_card
	GroupMasterId int `xorm:"pk 'group_master_id'"`
	CardMasterId  int `xorm:"pk 'card_master_id'"`
}

// can be shared depending on impl
// GuaranteedCardSet is not stored, built by gamedata when loaded, if applicable
// - (static) CardSet can be used to specify almost if not all the guaranteed form official version had
// - More exotic form of guarantee should be built into the handler itself
// - If CardSetSQL is not empty, GuaranteedCardSet would contain the relevant cards Id
type GachaGuarantee struct { // s_gacha_guarantee
	GachaGuaranteeMasterId int          `xorm:"pk 'gacha_guarantee_master_id'"` // unique id
	GuaranteeHandler       string       `xorm:"'handler'"`
	CardSetSQL             string       `xorm:"card_set_sql"`
	GuaranteedCardSet      map[int]bool `xorm:"-"`
}
