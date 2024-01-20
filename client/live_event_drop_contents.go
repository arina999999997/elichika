package client

import (
	"elichika/generic"
)

type LiveEventDropContents struct {
	StandardDrops generic.Array[Content] `json:"standard_drops"`
	BonusDrops    generic.Array[Content] `json:"bonus_drops"`
}
