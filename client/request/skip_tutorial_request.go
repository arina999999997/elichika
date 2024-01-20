package request

import (
	"elichika/client"
	"elichika/generic"
)

type SkipTutorialRequest struct {
	CardWithSuitDict generic.Dictionary[int32, generic.Nullable[int32]] `json:"card_with_suit_dict"`
	SquadDict        generic.Dictionary[int32, client.LiveSquad]        `json:"squad_dict"`
}
