package user_beginner_challenge

import (
	"elichika/userdata"
)

// this work similar to the mission system
type Handler = func(*userdata.Session, []any, ...any)

func UpdateProgress(session *userdata.Session, conditionType int32, conditionParam1, conditionParam2 *int32,
	handler Handler, handlerParams ...any) {

	challengeState := GetBeginnerChallengeCells(session)
	earliestUncleared := int32((1 << 31) - 1)
	for _, cell := range challengeState {
		if cell.IsRewardReceived {
			continue
		}
		challengeId := session.Gamedata.BeginnerChallengeCell[cell.CellId].ChallengeId
		if earliestUncleared > challengeId {
			earliestUncleared = challengeId
		}
	}
	if earliestUncleared == int32((1<<31)-1) { // all cleared
		return
	}
	var missionList []any
	// only update the current level progress
	for _, cell := range session.Gamedata.BeginnerChallenge[earliestUncleared].ChallengeCells {
		if challengeState[cell.Id].IsRewardReceived {
			continue
		}
		if cell.MissionClearConditionType != conditionType {
			continue
		}
		// if one side is nil, we have a match
		// this might not be the correct way for all case, so check again in the handler if necessary
		if (cell.MissionClearConditionParam1 != nil) && (conditionParam1 != nil) && (*cell.MissionClearConditionParam1 != *conditionParam1) {
			continue
		}
		if (cell.MissionClearConditionParam2 != nil) && (conditionParam2 != nil) && (*cell.MissionClearConditionParam2 != *conditionParam2) {
			continue
		}

		missionList = append(missionList, *challengeState[cell.Id])
	}
	handler(session, missionList, handlerParams...)
}
