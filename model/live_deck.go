package model

import (
	"elichika/client"
)

type UserLiveDeck struct {
	UserLiveDeckId int32                `xorm:"pk 'user_live_deck_id'" json:"user_live_deck_id"`
	Name           client.LocalizedText `xorm:"name" json:"name"` // deck name
	// TODO(refactor, 20240108): make these nullable
	CardMasterId1 int32 `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterId2 int32 `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterId3 int32 `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	CardMasterId4 int32 `xorm:"'card_master_id_4'" json:"card_master_id_4"`
	CardMasterId5 int32 `xorm:"'card_master_id_5'" json:"card_master_id_5"`
	CardMasterId6 int32 `xorm:"'card_master_id_6'" json:"card_master_id_6"`
	CardMasterId7 int32 `xorm:"'card_master_id_7'" json:"card_master_id_7"`
	CardMasterId8 int32 `xorm:"'card_master_id_8'" json:"card_master_id_8"`
	CardMasterId9 int32 `xorm:"'card_master_id_9'" json:"card_master_id_9"`
	SuitMasterId1 int32 `xorm:"'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterId2 int32 `xorm:"'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterId3 int32 `xorm:"'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterId4 int32 `xorm:"'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterId5 int32 `xorm:"'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterId6 int32 `xorm:"'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterId7 int32 `xorm:"'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterId8 int32 `xorm:"'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterId9 int32 `xorm:"'suit_master_id_9'" json:"suit_master_id_9"`
}

func (uld *UserLiveDeck) Id() int64 {
	return int64(uld.UserLiveDeckId)
}

type UserLiveParty struct {
	PartyId        int32                `xorm:"pk 'party_id'" json:"party_id"`
	UserLiveDeckId int32                `xorm:"'user_live_deck_id'" json:"user_live_deck_id"`
	Name           client.LocalizedText `xorm:"name" json:"name"` // deck name
	IconMasterId   int32                `xorm:"'icon_master_id'" json:"icon_master_id" `

	// TODO(refactor, 20240108): make these nullable
	CardMasterId1    int32  `xorm:"'card_master_id_1'" json:"card_master_id_1"`
	CardMasterId2    int32  `xorm:"'card_master_id_2'" json:"card_master_id_2"`
	CardMasterId3    int32  `xorm:"'card_master_id_3'" json:"card_master_id_3"`
	UserAccessoryId1 *int64 `xorm:"'user_accessory_id_1'" json:"user_accessory_id_1"` // null for empty
	UserAccessoryId2 *int64 `xorm:"'user_accessory_id_2'" json:"user_accessory_id_2"`
	UserAccessoryId3 *int64 `xorm:"'user_accessory_id_3'" json:"user_accessory_id_3"`
}

func (uld *UserLiveParty) Id() int64 {
	return int64(uld.PartyId)
}

// DeckSquadDict ...
type DeckSquadDict struct {
	CardMasterIds    []int    `json:"card_master_ids"`
	UserAccessoryIds []*int64 `json:"user_accessory_ids"`
}
