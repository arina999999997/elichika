package model

import (
	"elichika/generic"
)

type UserGpsPresentReceived struct {
	CampaignId int `xorm:"pk 'campaign_id'" json:"campaign_id"`
}

func (gpr *UserGpsPresentReceived) Id() int64 {
	return int64(gpr.CampaignId)
}

func init() {
	TableNameToInterface["u_gps_present_received"] = generic.UserIdWrapper[UserGpsPresentReceived]{}
}
