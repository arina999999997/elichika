package request

import (
	"elichika/generic"
)

type FinishUserStoryMemberRequest struct {
	StoryMemberMasterId int32                  `json:"story_member_master_id"`
	IsAutoMode        generic.Nullable[bool] `json:"is_auto_mode"`
}

