package client

import (
	"elichika/generic"
)

type LiveSquad struct {
	CardMasterIds    generic.Array[int32]                   `json:"card_master_ids"`
	UserAccessoryIds generic.Array[generic.Nullable[int64]] `json:"user_accessory_ids"`
}
