package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchLiveDeckSelectResponse struct {
	LastPlayLiveDifficultyDeck generic.Nullable[client.LastPlayLiveDifficultyDeck] `json:"last_play_live_difficulty_deck"` // pointer
}
