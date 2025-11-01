package client

import (
	"elichika/generic"
)

type EventMiningTopAnimationCellMasterRow struct {
	EventMiningMasterId int32                           `json:"event_mining_master_id"`
	ThumbnailCellId     int32                           `json:"thumbnail_cell_id"`
	AddStoryNumber      int32                           `json:"add_story_number"`
	Priority            int32                           `json:"priority"`
	MovieAssetPath      MovieStruktur                   `json:"movie_asset_path"`
	IsPopup             bool                            `json:"is_popup"`
	PopupMovieAssetPath MovieStruktur                   `json:"popup_movie_asset_path"`
	PopupComment        generic.Nullable[LocalizedText] `json:"popup_comment"`
	IsBig               bool                            `json:"is_big"`
	StartAt             int64                           `json:"start_at"`
	EndAt               int64                           `json:"end_at"`
}
