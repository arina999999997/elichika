package request

type SaveLiveDeckMemberSuitRequest struct {
	DeckId       int32 `json:"deck_id"`
	CardIndex    int32 `json:"card_index"`
	SuitMasterId int32 `json:"suit_master_id"`
	ViewStatus   int32 `json:"view_status" enum:"MemberViewStatus"`
}
