package client

import (
	"elichika/generic"
)

type LiveResultContentPack struct {
	StandardDrops   generic.Array[LiveDropContent] `json:"standard_drops"`
	AdditionalDrops generic.Array[LiveDropContent] `json:"additional_drops"`
	GimmickDrops    generic.Array[LiveDropContent] `json:"gimmick_drops"`
}
