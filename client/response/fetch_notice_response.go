package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchNoticeResponse struct {
	NoticeLists     generic.Dictionary[int32, client.NoticeList] `json:"notice_lists" enum:""`
	NoticeNoCheckAt int64                                        `json:"notice_no_check_at"`
}
