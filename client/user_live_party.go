package client

import (
	"elichika/generic"
)

type UserLiveParty struct {
	PartyId          int32                   `xorm:"pk 'party_id'" json:"party_id"`
	UserLiveDeckId   int32                   `xorm:"'user_live_deck_id'" json:"user_live_deck_id"`
	Name             LocalizedText           `xorm:"'name'" json:"name"` // deck name
	IconMasterId     int32                   `xorm:"'icon_master_id'" json:"icon_master_id"`
	CardMasterId1    generic.Nullable[int32] `xorm:"json 'card_master_id_1'" json:"card_master_id_1"`
	CardMasterId2    generic.Nullable[int32] `xorm:"json 'card_master_id_2'" json:"card_master_id_2"`
	CardMasterId3    generic.Nullable[int32] `xorm:"json 'card_master_id_3'" json:"card_master_id_3"`
	UserAccessoryId1 generic.Nullable[int64] `xorm:"json 'user_accessory_id_1'" json:"user_accessory_id_1"`
	UserAccessoryId2 generic.Nullable[int64] `xorm:"json 'user_accessory_id_2'" json:"user_accessory_id_2"`
	UserAccessoryId3 generic.Nullable[int64] `xorm:"json 'user_accessory_id_3'" json:"user_accessory_id_3"`
}

func (uld *UserLiveParty) Id() int64 {
	return int64(uld.PartyId)
}
