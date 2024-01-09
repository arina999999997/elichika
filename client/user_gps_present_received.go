package client

type UserGpsPresentReceived struct {
	CampaignId int32 `xorm:"pk 'campaign_id'" json:"campaign_id"`
}

func (ugpr *UserGpsPresentReceived) Id() int64 {
	return int64(ugpr.CampaignId)
}
