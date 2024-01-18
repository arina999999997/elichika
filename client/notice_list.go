package client

import (
	"elichika/generic"
)

type NoticeList struct {
	Category      int32                        `json:"category" enum:"NoticeSubCategory"`
	NewArrivalIds generic.List[int32]          `json:"new_arrival_ids"`
	CurrentPage   int32                        `json:"current_page"`
	MaxPage       int32                        `json:"max_page"`
	MaxIndex      int32                        `json:"max_index"`
	Notices       generic.Array[NoticeSummary] `json:"notices"`
}
