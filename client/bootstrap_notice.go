package client

import (
	"elichika/generic"
)

type BootstrapNotice struct {
	SuperNotices        generic.List[BootstrapSuperNotice] `json:"super_notices"`
	FetchedAt           int64                              `json:"fetched_at"`
	ReviewSuperNoticeAt int64                              `json:"review_super_notice_at"`
	ForceViewNoticeIds  generic.List[int32]                `json:"force_view_notice_ids"`
}
