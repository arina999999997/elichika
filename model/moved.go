package model

import (
	"elichika/client"
	"elichika/generic"
)

func init() {
	type DbMember struct {
		client.UserMember         `xorm:"extends"`
		LovePanelLevel            int   `xorm:"'love_panel_level' default 1"`
		LovePanelLastLevelCellIds []int `xorm:"'love_panel_last_level_cell_ids' default '[]'"`
	}
	TableNameToInterface["u_member"] = generic.UserIdWrapper[DbMember]{}
	TableNameToInterface["u_suit"] = generic.UserIdWrapper[client.UserSuit]{}
	TableNameToInterface["u_card"] = generic.UserIdWrapper[client.UserCard]{}
	TableNameToInterface["u_lesson_deck"] = generic.UserIdWrapper[client.UserLessonDeck]{}
	TableNameToInterface["u_accessory"] = generic.UserIdWrapper[client.UserAccessory]{}
	TableNameToInterface["u_live_deck"] = generic.UserIdWrapper[client.UserLiveDeck]{}
	TableNameToInterface["u_live_party"] = generic.UserIdWrapper[client.UserLiveParty]{}
}
