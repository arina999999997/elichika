package response

import (
	"elichika/client"
)

type FetchNoticeListResponse struct {
	NoticeList client.NoticeList `json:"notice_list"`
}
