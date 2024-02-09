package client

type NoticeDetail struct {
	NoticeId   int32         `json:"notice_id"`
	Category   int32         `json:"category" enum:"NoticeTagCategory"`
	Title      LocalizedText `json:"title"`
	DetailText LocalizedText `json:"detail_text"`
	Date       int64         `json:"date"`
}
