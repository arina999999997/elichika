package client

// for that one part in the main story where you select an idol
type UserStoryMainSelected struct {
	StoryMainCellId int32 `xorm:"pk 'story_main_cell_id'" json:"story_main_cell_id"`
	SelectedId      int32 `xorm:"'selected_id'" json:"selected_id"`
}
