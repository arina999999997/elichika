package request

import (
	"elichika/client"
	"elichika/generic"
)

type SaveLiveDeckAllRequest struct {
	DeckId       int32                                              `json:"deck_id"`
	CardWithSuit generic.Dictionary[int32, generic.Nullable[int32]] `json:"card_with_suit"`
	SquadDict    generic.Dictionary[int32, client.LiveSquad]        `json:"squad_dict"`
}
