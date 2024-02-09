package response

import (
	"elichika/client"
)

type FetchNoticeDetailResponse struct {
	Notice client.NoticeDetail `json:"notice"`
}
