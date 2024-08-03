package event

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_event/marathon"
	"elichika/subsystem/user_info_trigger"
	"elichika/userdata"
)

func FetchEventMarathon(session *userdata.Session, eventId int32) (*response.FetchEventMarathonResponse, *response.RecoverableExceptionResponse) {
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if (event == nil) || (event.EventId != eventId) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeEventMarathonOutOfDate,
		}
	}

	eventMarathon := session.Gamedata.EventActive.GetEventMarathon()
	resp := &response.FetchEventMarathonResponse{
		EventMarathonTopStatus: eventMarathon.TopStatus,
		UserModelDiff:          &session.UserModel,
	}
	resp.EventMarathonTopStatus.StartAt = event.StartAt
	resp.EventMarathonTopStatus.EndAt = event.EndAt
	resp.EventMarathonTopStatus.ResultAt = event.ResultAt
	resp.EventMarathonTopStatus.ExpiredAt = event.ExpiredAt
	userEventStatus := GetUserEventStatus(session, eventId)

	resp.EventMarathonTopStatus.IsFirstAccess = userEventStatus.IsFirstAccess

	if resp.EventMarathonTopStatus.IsFirstAccess {
		user_info_trigger.AddTriggerBasic(session,
			client.UserInfoTriggerBasic{
				InfoTriggerType: enum.InfoTriggerTypeEventMarathonFirstRuleDescription,
				ParamInt:        generic.NewNullable(eventId),
			})
	}

	userEventMarathon := marathon.GetUserEventMarathon(session)

	switch userEventMarathon.ReadStoryNumber {
	case 0:
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(eventMarathon.BoardMemos[0])
	case 7:
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(eventMarathon.BoardMemos[2])
	default:
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(eventMarathon.BoardMemos[1])
	}

	resp.EventMarathonTopStatus.StoryStatus.ReadStoryNumber = userEventMarathon.ReadStoryNumber
	for i := int32(0); i < userEventMarathon.ReadStoryNumber; i++ {
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Append(eventMarathon.BoardPictures[i])
	}
	if userEventMarathon.ReadStoryNumber > 0 {
		resp.EventMarathonTopStatus.BoardStatus.BoardThingMasterRows.Slice[userEventMarathon.ReadStoryNumber].IsEffect = userEventStatus.IsNew
		resp.EventMarathonTopStatus.BoardStatus.IsEffect = userEventStatus.IsNew
	}
	if resp.EventMarathonTopStatus.BoardStatus.IsEffect || resp.EventMarathonTopStatus.IsFirstAccess {
		userEventStatus.IsFirstAccess = false
		userEventStatus.IsNew = false
		UpdateUserEventStatus(session, userEventStatus)
	}

	nextRewardPoint, nextRewardContent := eventMarathon.GetNextReward(userEventMarathon.EventPoint)

	resp.EventMarathonTopStatus.UserRankingStatus = client.EventMarathonUserRanking{
		Order:           marathon.GetUserEventMarathonRanking(session, event.EventId),
		TotalPoint:      generic.NewNullable(userEventMarathon.EventPoint),
		NextRewardPoint: nextRewardPoint,
		RewardContent:   nextRewardContent,
	}
	return resp, nil
}
