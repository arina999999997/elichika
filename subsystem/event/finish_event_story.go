package event

import (
	"elichika/client/response"
	"elichika/generic"
	"elichika/log"
	"elichika/subsystem/user_event/marathon"
	"elichika/userdata"
)

func FinishEventStory(session *userdata.Session, storyEventMasterId int32, isAutoMode generic.Nullable[bool]) *response.UserModelResponse {
	if isAutoMode.HasValue {
		session.UserStatus.IsAutoMode = isAutoMode.Value
	}
	eventStory := session.Gamedata.EventStory[storyEventMasterId]
	userEvent := marathon.GetUserEventMarathon(session)
	log.Println(storyEventMasterId)
	log.Println(eventStory)
	log.Println(userEvent)
	if eventStory.EventMasterId != userEvent.EventMasterId {
		log.Panic("event changed")
	}
	if userEvent.ReadStoryNumber < eventStory.StoryNumber {
		userEvent.ReadStoryNumber = eventStory.StoryNumber
		session.UserModel.UserEventMarathonByEventMasterId.Set(userEvent.EventMasterId, userEvent)
		userEventStatus := GetUserEventStatus(session, eventStory.EventMasterId)
		userEventStatus.IsNew = true
		UpdateUserEventStatus(session, userEventStatus)
	}
	return &response.UserModelResponse{
		UserModel: &session.UserModel,
	}
}
