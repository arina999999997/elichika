package request

import (
	"elichika/generic"
)

type SaveLiveDeckCardsRequest struct {
	DeckId        int32                            `json:"deck_id"`
	CardMasterIds generic.Dictionary[int32, int32] `json:"card_master_ids"`
}
