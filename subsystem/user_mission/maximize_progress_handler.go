package user_mission

import (
	"elichika/userdata"
)

// take 1 param that is the count to maximize
func MaximizeProgressHandler(session *userdata.Session, missionList []any, params ...any) {
	for _, mission := range missionList {
		MaximizeMissionProgress(session, mission, params[0].(int32))
	}
}
