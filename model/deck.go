package model

import (
	"elichika/generic"
)

type UserLessonDeck struct {
	UserLessonDeckId int    `xorm:"pk 'user_lesson_deck_id'" json:"user_lesson_deck_id"`
	Name             string `json:"name"`
	CardMasterId1    *int   `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterId2    *int   `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterId3    *int   `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	CardMasterId4    *int   `xorm:"'card_master_id_4'" json:"card_master_id_4"`
	CardMasterId5    *int   `xorm:"'card_master_id_5'" json:"card_master_id_5"`
	CardMasterId6    *int   `xorm:"'card_master_id_6'" json:"card_master_id_6"`
	CardMasterId7    *int   `xorm:"'card_master_id_7'" json:"card_master_id_7"`
	CardMasterId8    *int   `xorm:"'card_master_id_8'" json:"card_master_id_8"`
	CardMasterId9    *int   `xorm:"'card_master_id_9'" json:"card_master_id_9"`
}

func (uld *UserLessonDeck) Id() int64 {
	return int64(uld.UserLessonDeckId)
}
func (uld *UserLessonDeck) SetId(id int64) {
	uld.UserLessonDeckId = int(id)
}

func init() {

	TableNameToInterface["u_lesson_deck"] = generic.UserIdWrapper[UserLessonDeck]{}
}
