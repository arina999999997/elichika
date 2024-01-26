package client

import (
	"elichika/generic"
)

type OtherUserDeck struct {
	Id            int32                         `json:"id"` // this field is irrelevant
	Name          LocalizedText                 `json:"name"`
	Parties       generic.Array[OtherUserParty] `json:"parties"`
	Cards         generic.Array[OtherUserCard]  `json:"cards"`
	CardIds       generic.Array[int32]          `json:"card_ids"`
	SuitMasterIds generic.Array[int32]          `json:"suit_master_ids"`
}
