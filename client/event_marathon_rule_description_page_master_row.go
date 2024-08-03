package client

type EventMarathonRuleDescriptionPageMasterRow struct {
	Page           int32           `json:"page"`
	Title          LocalizedText   `json:"title"`
	ImageAssetPath TextureStruktur `json:"image_asset_path"`
}
