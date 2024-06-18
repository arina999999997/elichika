package user_beginner_challenge

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/userdata"

	"sort"
)

func FetchChallengeBeginner(session *userdata.Session) response.FetchChallengeBeginnerResponse {
	cellState := GetBeginnerChallengeCells(session)
	challengeId := int32((1 << 31) - 1)
	resp := response.FetchChallengeBeginnerResponse{}
	// check the response and pick the earliest one not completed
	// if everything is completed, show the last one
	maxId := int32(0)
	for _, challenge := range session.Gamedata.BeginnerChallenge {
		isCompleted := true
		for _, cell := range challenge.ChallengeCells {
			if !cellState[cell.Id].IsRewardReceived { // not finished
				isCompleted = false
				break
			}
		}
		if (!isCompleted) && (challengeId > challenge.Id) {
			challengeId = challenge.Id
		} else if maxId < challenge.Id {
			maxId = challenge.Id
		}
		if isCompleted {
			resp.CompletedIds.Append(challenge.Id)
		}
	}
	sort.Slice(resp.CompletedIds.Slice, func(i, j int) bool {
		return resp.CompletedIds.Slice[i] < resp.CompletedIds.Slice[j]
	})
	if challengeId == (1<<31)-1 {
		challengeId = maxId
	}
	resp.ChallengeBeginner.ChallengeId = challengeId
	for _, cell := range session.Gamedata.BeginnerChallenge[challengeId].ChallengeCells {
		cellValue, exist := cellState[cell.Id]
		if exist {
			resp.ChallengeBeginner.Cells.Append(*cellValue)
		} else {
			resp.ChallengeBeginner.Cells.Append(client.ChallengeCell{
				CellId: cell.Id,
			})
		}
	}
	return resp
}
