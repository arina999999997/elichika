package client

import (
	"elichika/generic"
)

type GachaAppeal struct {
	CardMasterId   generic.Nullable[int32] `json:"card_master_id"`
	AppearanceType generic.Nullable[int32] `json:"appearance_type" enum:"CardAppearanceType"`
	MainImageAsset TextureStruktur         `json:"main_image_asset"`
	SubImageAsset  TextureStruktur         `json:"sub_image_asset"`
	TextImageAsset TextureStruktur         `json:"text_image_asset"`
}
