package client

type LiveCampaignNotice struct {
	NoticeId int32 `json:"notice_id"`
	EndAt    int64 `json:"end_at"`
}
