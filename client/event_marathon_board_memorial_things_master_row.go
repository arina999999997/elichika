package client

type EventMarathonBoardMemorialThingsMasterRow struct {
	EventMarathonBoardPositionType int32           `xorm:"'event_marathon_board_position_type'" json:"event_marathon_board_position_type" enum:"EventMarathonBoardPositionType"`
	Position                       int32           `xorm:"'position'" json:"position"`
	AddStoryNumber                 int32           `xorm:"'add_story_number'" json:"add_story_number"`
	Priority                       int32           `xorm:"'priority'" json:"priority"`
	ImageThumbnailAssetPath        TextureStruktur `xorm:"'image_thumbnail_asset_path'" json:"image_thumbnail_asset_path"`
	IsEffect                       bool            `xorm:"-" json:"is_effect"`
}
