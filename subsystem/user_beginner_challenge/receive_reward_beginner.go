package user_beginner_challenge

import (
	"elichika/client/response"
	"elichika/generic"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func ReceiveRewardBeginner(session *userdata.Session, challengeId int32, cellId generic.Nullable[int32]) response.ChallengeBeginnerRewardReceiveResponse {
	// if cellId isn't null, add the reward for that
	// if it is, add the reward for everything, and the reward for the full challenge if everything is cleared
	// the client also only send the challengeId, if claiming a cell would also claim the whole thing

	resp := response.ChallengeBeginnerRewardReceiveResponse{
		UserModel: &session.UserModel,
	}
	cellState := GetBeginnerChallengeCells(session)
	cellIds := []int32{}
	if cellId.HasValue { // only possible when it's not the last received
		cellIds = append(cellIds, cellId.Value)
	} else {
		allCleared := true
		for _, cell := range session.Gamedata.BeginnerChallenge[challengeId].ChallengeCells {
			if cellState[cell.Id].Progress < cell.MissionClearConditionCount {
				allCleared = false
				continue
			}
			if !cellState[cell.Id].IsRewardReceived {
				cellIds = append(cellIds, cell.Id)
			}
		}
		if allCleared {
			resp.ReceivedCompleteContents.Append(session.Gamedata.BeginnerChallenge[challengeId].CompleteReward)
			result := user_content.AddContent(session, session.Gamedata.BeginnerChallenge[challengeId].CompleteReward)
			if result != nil {
				resp.CompleteLimitExceeded = true
			}
		}
	}
	for _, cellId := range cellIds {
		for _, reward := range session.Gamedata.BeginnerChallengeCell[cellId].Rewards {
			resp.ReceivedCellContents.Append(reward)
			result := user_content.AddContent(session, reward)
			if result != nil {
				resp.CellLimitExceeded = true
			}

		}
		cellState[cellId].IsRewardReceived = true
		UpdateChallengeCell(session, *cellState[cellId])
	}

	fetchChallengeResponse := FetchChallengeBeginner(session)
	resp.ChallengeBeginner = fetchChallengeResponse.ChallengeBeginner
	resp.CompletedIds = fetchChallengeResponse.CompletedIds
	return resp
}
