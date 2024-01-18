package request

import (
	"elichika/generic"
)

type FetchNoticeListRequest struct {
	NoticeCategory int32                   `json:"notice_category" enum:"NoticeSubCategory"`
	Page           int32                   `json:"page"`
	FetchedAt      generic.Nullable[int64] `json:"fetched_at"`
	PreFetchedAt   generic.Nullable[int64] `json:"pre_fetched_at"`
}
