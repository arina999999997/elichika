package model

type UserGpsPresentReceived struct {
	UserId     int `xorm:"pk 'user_id'" json:"user_id"`
	CampaignId int `xorm:"pk 'campaign_id'" json:"campaign_id"`
}

func (gpr *UserGpsPresentReceived) Id() int64 {
	return int64(gpr.CampaignId)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_gps_present_received"] = UserGpsPresentReceived{}
}
