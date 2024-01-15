package response

import (
	"elichika/client"
	"elichika/generic"
)

type ExecuteLessonResponse struct {
	UserModelDiff        *client.UserModel                                                `json:"user_model_diff"`
	LessonMenuActions    generic.Dictionary[int32, generic.List[client.LessonMenuAction]] `json:"lesson_menu_actions"`
	LessonDropRarityList generic.Dictionary[int32, generic.List[int32]]                   `json:"lesson_drop_rarity_list" enum:"LessonDropRarityType"`
	IsSubscription       bool                                                             `json:"is_subscription"`
}
