package response

import (
	"elichika/client"
	"elichika/generic"
	"elichika/model"
)

type Login struct {
	SessionKey              string            `xorm:"-" json:"session_key"`
	UserModel               *client.UserModel `xorm:"-" json:"user_model"`
	IsPlatformServiceLinked bool              `xorm:"'is_platform_service_linked'" json:"is_platform_service_linked"`
	LastTimestamp           int64             `xorm:"'last_timestamp'" json:"last_timestamp"`
	Cautions                []int             `xorm:"'cautions'" json:"cautions"`
	ShowHomeCaution         bool              `xorm:"'show_home_caution'" json:"show_home_caution"`
	LiveResume              *int              `xorm:"-" json:"live_resume"`
	FromEea                 bool              `xorm:"'from_eea'" json:"from_eea"`
	GdprConsentedInfo       struct {
		HasConsentedAdPurposeOfUse bool `xorm:"'has_consented_ad_purpose_of_use'" json:"has_consented_ad_purpose_of_use"`
		HasConsentedCrashReport    bool `xorm:"'has_consented_crash_report'" json:"has_consented_crash_report"`
	} `xorm:"extends" json:"gdpr_consented_info"`
	MemberLovePanels []client.MemberLovePanel `xorm:"-" json:"member_love_panels"`
	CheckMaintenance bool                     `xorm:"-" json:"check_maintenance"`
	ReproInfo        struct {
		GroupNo int `xorm:"'group_no'" json:"group_no"`
	} `xorm:"extends" json:"repro_info"`
}

func init() {
	model.TableNameToInterface["u_login"] = generic.UserIdWrapper[Login]{}
}
