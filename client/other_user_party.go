package client

import (
	"elichika/generic"
)

type OtherUserParty struct {
	Id          int32                             `json:"id"`
	CardIds     generic.Array[int32]              `json:"card_ids"`
	Accessories generic.Array[OtherUserAccessory] `json:"accessories"`
}
