package user_mission

import (
	"elichika/userdata"
)

// take 1 param that is the count to add
func AddProgressHandler(session *userdata.Session, missionList []any, params ...any) {
	for _, mission := range missionList {
		AddMissionProgress(session, mission, params[0].(int32))
	}
}
