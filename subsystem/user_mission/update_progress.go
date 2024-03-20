package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

// how mission progress work:
// - because of many different types of mission condition, all the relevant data and context, ..., it's very hard to keep everything in one place
// - so mission checking logic is decentralized, similar to many other system
// - furthermore, there are thousands of missions, so we have to narrow it down in some way or it will take forever to calculate
// - so the actual interface we have is the following:
//   - in the relevant handler, call UpdateProgress
//   - UpdateProgress will filter out the missions that are tracked, then call the actual handler on the avaiable missions
//   - The handler then Update the mission using user_mission system

// condition checking use function based on the ConditionType
// - depend on how we choose to handle things, once of the following is most common:
//   - total match: the params must match for it to be accepted
//   - null match all: the params can be provided, but it will match null as well
// - the first type is when we have to track the info more multiple things at once, and then the whole thing too.
//   - for example, when we clear a song, then we have to mark the clear for each member
//   - but we also have to mark the clear for the whole song, so the 2nd type will overcount it
// - the second type is nice when we do things like clear a bond episode, the bond episode count toward the
// member's mission but it also count toward the general mission too

// session, then the list of mission, finally the forwarded params
type Handler = func(*userdata.Session, []any, ...any)

func UpdateProgress(session *userdata.Session, conditionType int32, conditionParam1, conditionParam2 *int32,
	handler Handler, handlerParams ...any) {
	var missionList []any
	for _, mission := range session.Gamedata.MissionByClearConditionType[conditionType] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}
		// condition type must be the same, if a request is applicable to multiple condition type, split it to multiple conditions
		if mission.MissionClearConditionType != conditionType {
			continue
		}
		// if one side is nil, we have a match
		// this might not be the correct way for all case, so check again in the handler if necessary
		if (mission.MissionClearConditionParam1 != nil) && (conditionParam1 != nil) && (*mission.MissionClearConditionParam1 != *conditionParam1) {
			continue
		}
		if (mission.MissionClearConditionParam2 != nil) && (conditionParam2 != nil) && (*mission.MissionClearConditionParam2 != *conditionParam2) {
			continue
		}

		switch mission.Term {
		case enum.MissionTermDaily:
			userDailyMission := getUserDailyMission(session, mission.Id)
			if userDailyMission.MissionMId == 0 {
				continue
			}
			if userDailyMission.IsReceivedReward {
				continue
			}
			missionList = append(missionList, userDailyMission)
		case enum.MissionTermWeekly:
			userWeeklyMission := getUserWeeklyMission(session, mission.Id)
			if userWeeklyMission.MissionMId == 0 {
				continue
			}
			if userWeeklyMission.IsReceivedReward {
				continue
			}
			missionList = append(missionList, userWeeklyMission)
		default:
			userMission := getUserMission(session, mission.Id)
			if userMission.MissionMId == 0 {
				continue
			}
			if userMission.IsReceivedReward {
				continue
			}
			missionList = append(missionList, userMission)
		}
	}
	handler(session, missionList, handlerParams...)
}
