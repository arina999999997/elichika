package model

type UserLessonDeck struct {
	UserID           int    `xorm:"pk 'user_id'" json:"-"`
	UserLessonDeckID int    `xorm:"pk 'user_lesson_deck_id'" json:"user_lesson_deck_id"`
	Name             string `json:"name"`
	CardMasterID1    *int   `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterID2    *int   `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterID3    *int   `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	CardMasterID4    *int   `xorm:"'card_master_id_4'" json:"card_master_id_4"`
	CardMasterID5    *int   `xorm:"'card_master_id_5'" json:"card_master_id_5"`
	CardMasterID6    *int   `xorm:"'card_master_id_6'" json:"card_master_id_6"`
	CardMasterID7    *int   `xorm:"'card_master_id_7'" json:"card_master_id_7"`
	CardMasterID8    *int   `xorm:"'card_master_id_8'" json:"card_master_id_8"`
	CardMasterID9    *int   `xorm:"'card_master_id_9'" json:"card_master_id_9"`
}

func (uld *UserLessonDeck) ID() int64 {
	return int64(uld.UserLessonDeckID)
}
