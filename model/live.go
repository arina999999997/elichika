package model

// LiveDaily ...
type LiveDaily struct {
	LiveDailyMasterId      int `json:"live_daily_master_id" xorm:"id"`
	LiveMasterId           int `json:"live_master_id" xorm:"live_id"`
	EndAt                  int `json:"end_at"`
	RemainingPlayCount     int `json:"remaining_play_count"`
	RemainingRecoveryCount int `json:"remaining_recovery_count"`
}

// LivePartnerInfo ...
type LivePartnerInfo UserBasicInfo

// guests before live start
type LiveStartLivePartner struct {
	UserId int `xorm:"'user_id' "json:"user_id"`
	Name   struct {
		DotUnderText string `xorm:"'name'" json:"dot_under_text"`
	} `xorm:"extends" json:"name"`
	Rank                int   `json:"rank"`
	LastLoginAt         int64 `json:"last_login_at"`
	CardByCategory      []any `xorm:"-" json:"card_by_category"`
	EmblemId            int   `xorm:"'emblem_id'" json:"emblem_id"`
	IsFriend            bool  `xorm:"-" json:"is_friend"`
	IntroductionMessage struct {
		DotUnderText string `xorm:"'message'" json:"dot_under_text"`
	} `xorm:"extends" json:"introduction_message"`
}

type LiveUpdatePlayListReq struct {
	LiveMasterId int32 `json:"live_master_id"`
	GroupNum     int32 `json:"group_num"`
	IsSet        bool  `json:"is_set"`
}
