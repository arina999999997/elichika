package client

type EventTopicReward struct {
	DisplayOrder            int32           `json:"display_order"`
	RewardContent           Content         `json:"reward_content"`
	MainNameTopAssetPath    TextureStruktur `json:"main_name_top_asset_path"`
	MainNameBottomAssetPath TextureStruktur `json:"main_name_bottom_asset_path"`
	SubNameTopAssetPath     TextureStruktur `json:"sub_name_top_asset_path"`
	SubNameBottomAssetPath  TextureStruktur `json:"sub_name_bottom_asset_path"`
}
