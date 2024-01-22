package client

type UserStoryLinkage struct {
	StoryLinkageCellMasterId int32 `xorm:"pk 'story_linkage_cell_master_id'" json:"story_linkage_cell_master_id"`
}
