package model

import (
	"elichika/client"
)

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

// type GachaDraw struct { // s_gacha_draw
// 	GachaMasterId        int    `xorm:"gacha_master_id" json:"-"` // client only send gacha_draw_master_id
// 	GachaDrawMasterId    int    `xorm:"pk 'gacha_draw_master_id'" json:"gacha_draw_master_id"`
// 	RecoveryType         int    `json:"recover_type"`
// 	RecoverAt            *int64 `json:"recover_at"`
// 	DrawCount            int    `json:"draw_count"`
// 	GachaPaymentType     int    `json:"gacha_payment_type"`
// 	GachaPaymentMasterId int    `xorm:"gacha_payment_master_id" json:"gacha_payment_master_id"`
// 	GachaPaymentAmount   int    `json:"gacha_payment_amount"`
// 	// doesn't seem to be used? maybe for paid festival?
// 	GachaPointAmount int `json:"gacha_point_amount"`
// 	Description      struct {
// 		DotUnderText string `xorm:"'description'" json:"dot_under_text"`
// 	} `xorm:"extends" json:"description"`
// 	IsBonus         bool `json:"is_bonus"`
// 	BonusAppealText struct {
// 		DotUnderText string `xorm:"'bonus_appeal_text'" json:"dot_under_text"`
// 	} `xorm:"extends" json:"bonus_appeal_text"`
// 	RetryCount           *int   `json:"retry_count"`
// 	GachaRetryMasterRows [0]int `json:"gacha_retry_master_rows"`
// 	DailyLimit           int    `json:"daily_limit"`    // 0 for unlimited
// 	DailyInterval        int    `json:"daily_interval"` // always 1
// 	TermLimit            int    `json:"term_limit"`     // 0 for unlimited
// 	// number of remaining pull in this day
// 	RemainDayCount *int `json:"remain_day_count"`
// 	// number of remaining pull in this banner
// 	RemainTermCount *int  `json:"remain_term_count"`
// 	PerformanceId   int   `xorm:"performance_id" json:"performance_id"` // gacha performance
// 	IsSubscription  bool  `json:"is_subscription"`                      // is subscription only gacha
// 	Guarantees      []int `xorm:"'guarantees'" json:"-"`                // ids to gacha.GachaGuarantee
// }

type GachaDrawStepupDetail struct { // u_gacha_draw_stepup
	GachaMasterId int  `json:"-"`
	CurrentStep   int  `json:"current_step"`
	LoopCount     int  `json:"loop_count"`
	MaxLoop       int  `json:"max_loop"`
	MaxStep       int  `json:"max_step"`
	IsMaxNextStep bool `json:"is_max_next_step"`
}

type ResultCard struct {
	GachaLotType         int             `json:"gacha_lot_type"` // 1 for normal, 2 for guaranteed
	CardMasterId         int             `json:"card_master_id"`
	Level                int             `json:"level"`                   // always 1
	BeforeGrade          int             `json:"before_grade"`            // 0 for new
	AfterGrade           int             `json:"after_grade"`             // 0 for new, 5 for maxed
	Content              *client.Content `json:"content"`                 // if maxed then award radiance. Technically we can award anything, but this is only the display value
	LimitExceeded        bool            `json:"limit_exceeded"`          // always false
	BeforeLoveLevelLimit int             `json:"before_love_level_limit"` // always correct
	AfterLoveLevelLimit  int             `json:"after_love_level_limit"`  // 0 for max level
}
