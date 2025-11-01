package request

type LikeEventMiningPanelRequest struct {
	EventId         int32 `json:"event_id"`
	ThumbnailCellId int32 `json:"thumbnail_cell_id"`
}
