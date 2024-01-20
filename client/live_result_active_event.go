package client

// These types are implemented but we don't really have a good idea what should be null or not
// it's probably possible to reconstruct an event and see
type LiveResultActiveEvent struct {
	EventId               int32                            `json:"event_id"`
	EventType             int32                            `json:"event_type"`
	EventLogoAssetPath    TextureStruktur                  `json:"event_logo_asset_path"`     // pointer
	ReceivePoint          LiveResultActiveEventPoint       `json:"receive_point"`             // pointer
	TotalPoint            LiveResultActiveEventPoint       `json:"total_point"`               // pointer
	BonusPoint            LiveResultActiveEventPoint       `json:"bonus_point"`               // pointer
	OpenedEventStory      EventResultOpenedNewStory        `json:"opened_event_story"`        // pointer
	LiveEventDropItemInfo LiveEventDropItemInfo            `json:"live_event_drop_item_info"` // pointer
	PointReward           LiveResultActiveEventPointReward `json:"point_reward"`              // pointer
	IsStartLoopReward     bool                             `json:"is_start_loop_reward"`
}
