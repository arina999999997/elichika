package client

import (
	"elichika/generic"
)

type LiveEventDropItemInfo struct {
	BonusRate             int32                                `json:"bonus_rate"`
	IsSendToPresentBox    bool                                 `json:"is_send_to_present_box"`
	LiveEventDropContents generic.Array[LiveEventDropContents] `json:"live_event_drop_contents"`
}
