package database

import (
	"elichika/generic"
)

// These are some info not stored by user model
type UserEventStatus struct {
	EventId int32 `xorm:"pk 'event_id'"`

	IsFirstAccess bool `xorm:"'is_first_access'"`
	// whether the event top state need to be shown as new.
	IsNew bool `xorm:"'is_new'"`

	IsRewardReceived bool `xorm:"'is_reward_received'"`
}

type UserEventMiningTopCellState struct {
	EventId         int32 `xorm:"'event_id'"`
	ThumbnailCellId int32 `xorm:"pk 'thumbnail_cell_id'"`
	IsLike          bool  `xorm:"'is_like'"`
}

type UserEventMiningMusicBestScore struct {
	UserId           int32 `xorm:"'user_id'"`
	LiveId           int32 `xorm:"'live_id'"`
	LiveDifficultyId int32 `xorm:"'live_difficulty_id'"`
	Score            int32 `xorm:"'score'"`
}

type UserEventMiningTradeProduct struct {
	ProductId   int32 `xorm:"pk 'product_id'"`
	TradedCount int32 `xorm:"'traded_count'"`
}

func init() {
	AddTable("u_event_status", generic.UserIdWrapper[UserEventStatus]{})
	AddTable("u_event_mining_top_cell_state", generic.UserIdWrapper[UserEventMiningTopCellState]{})
	AddTable("u_event_mining_music_best_score", UserEventMiningMusicBestScore{})
	AddTable("u_event_mining_trade_product", generic.UserIdWrapper[UserEventMiningTradeProduct]{})
}
