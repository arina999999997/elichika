package user_mission

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func ReceiveReward(session *userdata.Session, missionIds []int32) any {
	// the behavior is as follow:
	// - mission master id list get filled like fetch_mission
	// - user model mission list only has relevant mission (expired or updated or received or ...)
	// - received_item are stacked, each item only have one appreance
	// - if there is an expired item, return an error response instead
	session.SendMissionDetail = true
	expired := generic.List[int32]{}
	itemGain := map[int32]map[int32]int32{}
	for _, missionId := range missionIds {
		masterMission := session.Gamedata.Mission[missionId]

		for _, content := range masterMission.Rewards {
			_, exist := itemGain[content.ContentType]
			if !exist {
				itemGain[content.ContentType] = map[int32]int32{}
			}
			itemGain[content.ContentType][content.ContentId] = itemGain[content.ContentType][content.ContentId] + content.ContentAmount
		}

		switch masterMission.Term {
		case enum.MissionTermDaily:
			userDailyMission := getUserDailyMission(session, missionId)
			if !userDailyMission.IsCleared { // expired
				expired.Append(missionId)
			} else {
				userDailyMission.IsReceivedReward = true
			}
			session.UserModel.UserDailyMissionByMissionId.Set(missionId, userDailyMission)
		case enum.MissionTermWeekly:
			userWeeklyMission := getUserWeeklyMission(session, missionId)
			if !userWeeklyMission.IsCleared { // expired
				expired.Append(missionId)
			} else {
				userWeeklyMission.IsReceivedReward = true
			}
			session.UserModel.UserWeeklyMissionByMissionId.Set(missionId, userWeeklyMission)
		default:
			userMission := getUserMission(session, missionId)
			userMission.IsReceivedReward = true
			session.UserModel.UserMissionByMissionId.Set(missionId, userMission)
		}
	}
	if expired.Size() > 0 { // has expired
		return response.MissionReceiveErrorResponse{
			MissionMasterIdList: FetchMissionIds(session),
			UserModel:           &session.UserModel,
			ExpiredMissionIds:   expired,
		}
	} else {
		items := generic.List[client.Content]{}
		for contentType, contents := range itemGain {
			for contentId, contentAmount := range contents {
				items.Append(client.Content{
					ContentType:   contentType,
					ContentId:     contentId,
					ContentAmount: contentAmount,
				})
			}
		}
		for _, content := range items.Slice {
			user_content.AddContent(session, content)
		}

		return response.MissionReceiveResponse{
			MissionMasterIdList:  FetchMissionIds(session),
			UserModel:            &session.UserModel,
			ReceivedPresentItems: items,
			LimitExceeded:        len(session.UnreceivedContent) > 0,
		}
	}
}
