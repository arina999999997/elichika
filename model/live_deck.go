package model

// TODO(refactor): Figure out this one, it should be LiveSquad
type DeckSquadDict struct {
	CardMasterIds    []int32  `json:"card_master_ids"`
	UserAccessoryIds []*int64 `json:"user_accessory_ids"`
	// UserAccessoryIds []generic.Nullable[int64] `json:"user_accessory_ids"`
}
