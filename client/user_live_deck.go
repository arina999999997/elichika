package client

import (
	"elichika/generic"
)

type UserLiveDeck struct {
	UserLiveDeckId int32                   `xorm:"pk 'user_live_deck_id'" json:"user_live_deck_id"`
	Name           LocalizedText           `xorm:"'name'" json:"name"` // deck name
	CardMasterId1  generic.Nullable[int32] `xorm:"json 'card_master_id_1'" json:"card_master_id_1"`
	CardMasterId2  generic.Nullable[int32] `xorm:"json 'card_master_id_2'" json:"card_master_id_2"`
	CardMasterId3  generic.Nullable[int32] `xorm:"json 'card_master_id_3'" json:"card_master_id_3"`
	CardMasterId4  generic.Nullable[int32] `xorm:"json 'card_master_id_4'" json:"card_master_id_4"`
	CardMasterId5  generic.Nullable[int32] `xorm:"json 'card_master_id_5'" json:"card_master_id_5"`
	CardMasterId6  generic.Nullable[int32] `xorm:"json 'card_master_id_6'" json:"card_master_id_6"`
	CardMasterId7  generic.Nullable[int32] `xorm:"json 'card_master_id_7'" json:"card_master_id_7"`
	CardMasterId8  generic.Nullable[int32] `xorm:"json 'card_master_id_8'" json:"card_master_id_8"`
	CardMasterId9  generic.Nullable[int32] `xorm:"json 'card_master_id_9'" json:"card_master_id_9"`
	SuitMasterId1  generic.Nullable[int32] `xorm:"json 'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterId2  generic.Nullable[int32] `xorm:"json 'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterId3  generic.Nullable[int32] `xorm:"json 'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterId4  generic.Nullable[int32] `xorm:"json 'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterId5  generic.Nullable[int32] `xorm:"json 'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterId6  generic.Nullable[int32] `xorm:"json 'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterId7  generic.Nullable[int32] `xorm:"json 'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterId8  generic.Nullable[int32] `xorm:"json 'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterId9  generic.Nullable[int32] `xorm:"json 'suit_master_id_9'" json:"suit_master_id_9"`
}

func (uld *UserLiveDeck) Id() int64 {
	return int64(uld.UserLiveDeckId)
}
