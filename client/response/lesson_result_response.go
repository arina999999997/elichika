package response

import (
	"elichika/client"
	"elichika/generic"
)

type LessonResultResponse struct {
	UserModelDiff  *client.UserModel                                 `xorm:"-" json:"user_model_diff"`
	SelectedDeckId int32                                             `json:"selected_deck_id"`
	DropItemList   generic.List[client.LessonDropItem]               `xorm:"json" json:"drop_item_list"`
	DropSkillList  generic.List[client.LessonResultDropPassiveSkill] `xorm:"json" json:"drop_skill_list"`
}
