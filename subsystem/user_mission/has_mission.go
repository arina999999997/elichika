package user_mission

import (
	"elichika/enum"
	"elichika/subsystem/user_subscription_status"
	"elichika/userdata"
)

func hasMissionPickupType(session *userdata.Session, pickupType *int32) bool {
	if pickupType == nil {
		return true
	}
	if *pickupType == enum.MissionPickupTypeSubscription {
		return user_subscription_status.HasSubscription(session)
	}
	return true
}

func hasTriggerCondition(session *userdata.Session, missionId int32) bool {
	masterMissionId := session.Gamedata.Mission[missionId]
	if masterMissionId == nil {
		return false
	}

	switch masterMissionId.TriggerType {
	case enum.MissionTriggerGameStart:
		return true
	case enum.MissionTriggerClearMission:
		return finishedMission(session, masterMissionId.TriggerCondition1)
	default:
		panic("unsuported mission trigger type")
	}
}

func hasMission(session *userdata.Session, missionId int32) bool {
	masterMissionId := session.Gamedata.Mission[missionId]
	if masterMissionId == nil {
		return false
	}
	return hasMissionPickupType(session, masterMissionId.PickupType) && hasTriggerCondition(session, missionId)
}
