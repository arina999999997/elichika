package event

import (
	"elichika/client/response"
	"elichika/generic"
	"elichika/log"
	"elichika/subsystem/user_event/mining"
	"elichika/userdata"
)

func FinishEventMiningStory(session *userdata.Session, storyEventMasterId int32, isAutoMode generic.Nullable[bool]) *response.UserModelResponse {
	if isAutoMode.HasValue {
		session.UserStatus.IsAutoMode = isAutoMode.Value
	}
	eventStory := session.Gamedata.EventStory[storyEventMasterId]
	userEvent := mining.GetUserEventMining(session)
	if eventStory.EventMasterId != userEvent.EventMasterId {
		log.Panic("event changed")
	}
	if userEvent.ReadStoryNumber < eventStory.StoryNumber {
		userEvent.ReadStoryNumber = eventStory.StoryNumber
		session.UserModel.UserEventMiningByEventMasterId.Set(userEvent.EventMasterId, userEvent)
		userEventStatus := GetUserEventStatus(session, eventStory.EventMasterId)
		userEventStatus.IsNew = true
		UpdateUserEventStatus(session, userEventStatus)
	}
	return &response.UserModelResponse{
		UserModel: &session.UserModel,
	}
}
