package client

type EventResultOpenedNewStory struct {
	Title                 string          `json:"title"`
	PreviewImageAssetPath TextureStruktur `json:"preview_image_asset_path"`
}
