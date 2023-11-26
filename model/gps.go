package model

type UserGpsPresentReceived struct {
	UserID     int `xorm:"pk 'user_id'" json:"user_id"`
	CampaignID int `xorm:"pk 'campaign_id'" json:"campaign_id"`
}

func (gpr *UserGpsPresentReceived) ID() int64 {
	return int64(gpr.CampaignID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_gps_present_received"] = UserGpsPresentReceived{}
}
