package client

import (
	"elichika/generic"
)

type EventMiningTopStillCellMasterRow struct {
	EventMiningMasterId          int32                           `json:"event_mining_master_id"`
	ThumbnailCellId              int32                           `json:"thumbnail_cell_id"`
	AddStoryNumber               int32                           `json:"add_story_number"`
	Priority                     int32                           `json:"priority"`
	ImageThumbnailAssetPath      TextureStruktur                 `json:"image_thumbnail_asset_path"`
	IsPopup                      bool                            `json:"is_popup"`
	PopupImageThumbnailAssetPath TextureStruktur                 `json:"popup_image_thumbnail_asset_path"`
	PopupComment                 generic.Nullable[LocalizedText] `json:"popup_comment"`
	IsBig                        bool                            `json:"is_big"`
}
