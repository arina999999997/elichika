package request

import (
	"elichika/generic"
)

type ChallengeBeginnerRewardRequest struct {
	ChallengeId     int32                   `json:"challenge_id"`
	ChallengeCellId generic.Nullable[int32] `json:"challenge_cell_id"`
}
