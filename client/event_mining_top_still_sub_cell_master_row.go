package client

type EventMiningTopStillSubCellMasterRow struct {
	EventMiningMasterId          int32           `json:"event_mining_master_id"`
	PanelSetId                   int32           `json:"panel_set_id"`
	ThumbnailCellId              int32           `json:"thumbnail_cell_id"`
	Priority                     int32           `json:"priority"`
	ImageThumbnailAssetPath      TextureStruktur `json:"image_thumbnail_asset_path"`
	IsPopup                      bool            `json:"is_popup"`
	PopupImageThumbnailAssetPath TextureStruktur `json:"popup_image_thumbnail_asset_path"`
}
