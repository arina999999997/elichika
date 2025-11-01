package event

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/subsystem/user_event/mining"
	"elichika/userdata"
)

func FetchEventMining(session *userdata.Session, eventId int32) (*response.FetchEventMiningResponse, *response.RecoverableExceptionResponse) {
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if (event == nil) || (event.EventId != eventId) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeEventMiningOutOfDate,
		}
	}

	eventMining := session.Gamedata.EventActive.GetEventMining()
	resp := &response.FetchEventMiningResponse{
		EventMiningTopStatus: eventMining.TopStatus,
		UserModelDiff:        &session.UserModel,
	}
	resp.EventMiningTopStatus.StartAt = event.StartAt
	resp.EventMiningTopStatus.EndAt = event.EndAt
	resp.EventMiningTopStatus.ResultAt = event.ResultAt
	resp.EventMiningTopStatus.ExpiredAt = event.ExpiredAt
	userEventStatus := GetUserEventStatus(session, eventId)
	if userEventStatus.IsNew {
		resp.EventMiningTopStatus.IsAddNewPanel = true
		userEventStatus.IsNew = false
		UpdateUserEventStatus(session, userEventStatus)
	}
	userEventMining := mining.GetUserEventMining(session)
	n := 0
	for ; n < len(resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice); n++ {
		if resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice[n].AddStoryNumber > userEventMining.ReadStoryNumber {
			break
		}
	}
	resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice = resp.EventMiningTopStatus.EventMiningTopStillCellMasterRows.Slice[:n]
	resp.EventMiningTopStatus.StoryStatus.ReadStoryNumber = userEventMining.ReadStoryNumber
	resp.EventMiningTopStatus.UserRankingStatus = mining.GetEventMiningUserRanking(session)
	resp.EventMiningTopStatus.EventMiningTopCellStateList = mining.GetEventMiningTopCellStateList(session, eventId)
	resp.EventMiningTopStatus.EventMiningMusicBestScoreList = mining.GetUserVoltageRankingSummary(session, eventId).GetMusicBestScoreList()
	return resp, nil
}
