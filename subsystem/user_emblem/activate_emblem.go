package user_emblem

import (
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

func ActivateEmblem(session *userdata.Session, emblemMasterId int32) {
	session.UserStatus.EmblemId = emblemMasterId
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountActiveEmblem, nil, nil,
		user_mission.AddProgressHandler, int32(1))
}
