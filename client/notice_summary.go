package client

type NoticeSummary struct {
	NoticeId        int32           `json:"notice_id"`
	Category        int32           `json:"category"`
	IsNew           bool            `json:"is_new"`
	Title           LocalizedText   `json:"title"`
	Date            int64           `json:"date"`
	BannerThumbnail TextureStruktur `json:"banner_thumbnail"`
}
