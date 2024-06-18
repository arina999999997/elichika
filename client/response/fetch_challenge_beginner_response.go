package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchChallengeBeginnerResponse struct {
	ChallengeBeginner client.ChallengeBeginner `json:"challenge_beginner"`
	CompletedIds      generic.List[int32]      `json:"completed_ids"`
}
