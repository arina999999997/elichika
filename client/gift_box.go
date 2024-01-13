package client

import (
	"elichika/generic"
)

type GiftBox struct {
	IsInPeriodGiftBox bool                          `json:"is_in_period_gift_box"`
	GiftBoxContent    generic.Array[GiftBoxContent] `json:"gift_box_content"`
}
