package client

import (
	"elichika/generic"
)

// This is the correct name for data transfer. Fyi Banner is an internal type to handle banners
type Banner1 struct {
	BannerMasterId       int32                   `json:"banner_master_id"`
	BannerImageAssetPath TextureStruktur         `json:"banner_image_asset_path"`
	BannerType           int32                   `json:"banner_type" enum:"BannerType"`
	ExpireAt             generic.Nullable[int64] `json:"expire_at"`
	TransitionId         int32                   `json:"transition_id"`
	TransitionParameter  generic.Nullable[int32] `json:"transition_parameter"`
}
