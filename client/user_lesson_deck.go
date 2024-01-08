package client

import (
	"elichika/generic"
)

type UserLessonDeck struct {
	UserLessonDeckId int32                   `xorm:"pk 'user_lesson_deck_id'" json:"user_lesson_deck_id"`
	Name             string                  `xorm:"'name'" json:"name"`
	CardMasterId1    generic.Nullable[int32] `xorm:"json 'card_master_id_1'" json:"card_master_id_1"`
	CardMasterId2    generic.Nullable[int32] `xorm:"json 'card_master_id_2'" json:"card_master_id_2"`
	CardMasterId3    generic.Nullable[int32] `xorm:"json 'card_master_id_3'" json:"card_master_id_3"`
	CardMasterId4    generic.Nullable[int32] `xorm:"json 'card_master_id_4'" json:"card_master_id_4"`
	CardMasterId5    generic.Nullable[int32] `xorm:"json 'card_master_id_5'" json:"card_master_id_5"`
	CardMasterId6    generic.Nullable[int32] `xorm:"json 'card_master_id_6'" json:"card_master_id_6"`
	CardMasterId7    generic.Nullable[int32] `xorm:"json 'card_master_id_7'" json:"card_master_id_7"`
	CardMasterId8    generic.Nullable[int32] `xorm:"json 'card_master_id_8'" json:"card_master_id_8"`
	CardMasterId9    generic.Nullable[int32] `xorm:"json 'card_master_id_9'" json:"card_master_id_9"`
}

func (uld *UserLessonDeck) Id() int64 {
	return int64(uld.UserLessonDeckId)
}

func (uld *UserLessonDeck) SetId(id int64) {
	uld.UserLessonDeckId = int32(id)
}
