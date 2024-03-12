package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

// upgrade the progress of missions based on clear condition
// this will only update uncleared missions (unclaimed reward), or daily/weekly missions
// this is a very general functions, and should only be used for commonly triggered changes
// for items that are harder to check, we might want to only check them when the mission is still there
// so we should pass a lazily evaluated function

// condition checking use function based on the ConditionType
// - depend on how we choose to handle things, once of the following is most common:
//   - total match: the params must match for it to be accepted
//   - null match all: the params can be provided, but it will match null as well
// - the first type is when we have to track the info more multiple things at once, and then the whole thing too.
//   - for example, when we clear a song, then we have to mark the clear for each member
//   - but we also have to mark the clear for the whole song, so the 2nd type will overcount it
// - the second type is nice when we do things like clear a bond episode, the bond episode count toward the
// member's mission but it also count toward the general mission too

// TODO(now): let's make it so there is a checking function per conditionType, and this function only get called when it's relevant
// - The relevant subsystem register a list of possible calls with relevant data, ideally it doesn't do any computation yet
// - Then this system figure out the relevant missions that those calls apply to, so not cleared stuff.
// - And then that function is invoked with the original data and the list of relevant mission
// - When a mission is finished, we can get new mission or trigger the progress of another mission, it should be fine to just call the relevant handler again
func UpdateProgress(session *userdata.Session, count, conditionType int32, conditionParam1, conditionParam2 *int32) {
	// TODO(optimisation): For now we iterate through the missions, but it's might be good to have a map
	for _, mission := range session.Gamedata.MissionByClearConditionType[conditionType] {
		switch mission.Term {
		case enum.MissionTermDaily:
			userDailyMission := getUserDailyMission(session, mission.Id)
			if userDailyMission.MissionMId == 0 {
				continue
			}
			updateCountByConditionType(conditionType, &userDailyMission.MissionCount, count)
			if userDailyMission.MissionCount >= userDailyMission.MissionStartCount+mission.MissionClearConditionCount {
				userDailyMission.IsCleared = true
			}
			// updateUserDailyMission(session, userDailyMission)
		case enum.MissionTermWeekly:
			userWeeklyMission := getUserWeeklyMission(session, mission.Id)
			if userWeeklyMission.MissionMId == 0 {
				continue
			}
			updateCountByConditionType(conditionType, &userWeeklyMission.MissionCount, count)
			if userWeeklyMission.MissionCount >= userWeeklyMission.MissionStartCount+mission.MissionClearConditionCount {
				userWeeklyMission.IsCleared = true
			}
			// updateUserWeeklyMission(session, userWeeklyMission)
		default:
			userMission := getUserMission(session, mission.Id)
			if userMission.MissionMId == 0 {
				continue
			}
			updateCountByConditionType(conditionType, &userMission.MissionCount, count)
			if userMission.MissionCount > mission.MissionClearConditionCount {
				userMission.IsCleared = true
			}
			// updateUserMission(session, userMission)
		}
	}
}
