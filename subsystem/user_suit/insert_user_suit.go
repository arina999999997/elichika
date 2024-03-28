package user_suit

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

// suit are inserted when the function is called as suit is unique and doesn't change

func InsertUserSuits(session *userdata.Session, suits []client.UserSuit) {
	for _, suit := range suits {
		if !userdata.GenericDatabaseExist(session, "u_suit", suit) {
			user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountSuit, nil, nil,
				user_mission.AddProgressHandler, int32(1))
			session.UserModel.UserSuitBySuitId.Set(suit.SuitMasterId, suit)
		}
	}
}

func InsertUserSuit(session *userdata.Session, suitMasterId int32) {
	suit := client.UserSuit{
		SuitMasterId: suitMasterId,
		IsNew:        true,
	}
	if !userdata.GenericDatabaseExist(session, "u_suit", suit) {
		user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountSuit, nil, nil,
			user_mission.AddProgressHandler, int32(1))
		session.UserModel.UserSuitBySuitId.Set(suit.SuitMasterId, suit)
	}
}
