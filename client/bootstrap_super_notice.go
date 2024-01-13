package client

type BootstrapSuperNotice struct {
	SuperNoticeId int32         `json:"super_notice_id"`
	MoviePath     MovieStruktur `json:"movie_path"`
	IsOnce        bool          `json:"is_once"`
	LastUpdatedAt int64         `json:"last_updated_at"`
}
