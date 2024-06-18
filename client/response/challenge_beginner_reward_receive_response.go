package response

import (
	"elichika/client"
	"elichika/generic"
)

type ChallengeBeginnerRewardReceiveResponse struct {
	UserModel                *client.UserModel            `json:"user_model"`
	ReceivedCellContents     generic.List[client.Content] `json:"received_cell_contents"`
	ReceivedCompleteContents generic.List[client.Content] `json:"received_complete_contents"`
	CellLimitExceeded        bool                         `json:"cell_limit_exceeded"`
	CompleteLimitExceeded    bool                         `json:"complete_limit_exceeded"`
	ChallengeBeginner        client.ChallengeBeginner     `json:"challenge_beginner"`
	CompletedIds             generic.List[int32]          `json:"completed_ids"`
}
