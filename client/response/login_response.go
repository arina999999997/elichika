package response

import (
	"elichika/client"
	"elichika/generic"
)

type LoginResponse struct {
	SessionKey              string                                `xorm:"-" json:"session_key"`
	UserModel               *client.UserModel                     `xorm:"-" json:"user_model"`
	IsPlatformServiceLinked bool                                  `json:"is_platform_service_linked"`
	LastTimestamp           int64                                 `json:"last_timestamp"`
	Cautions                generic.List[client.Caution]          `xorm:"-" json:"cautions"`
	ShowHomeCaution         bool                                  `xorm:"-" json:"show_home_caution"`
	LiveResume              generic.Nullable[client.LiveResume]   `xorm:"-" json:"live_resume"` // pointer
	FromEea                 bool                                  `json:"from_eea"`
	GdprConsentedInfo       client.UserGdprConsentedInfo          `xorm:"extends" json:"gdpr_consented_info"` // pointer
	MemberLovePanels        generic.Array[client.MemberLovePanel] `xorm:"-" json:"member_love_panels"`
	CheckMaintenance        bool                                  `xorm:"-" json:"check_maintenance"`
	ReproInfo               client.ReproInfo                      `xorm:"extends" json:"repro_info"`
}
