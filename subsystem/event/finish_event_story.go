package event

import (
	"elichika/client/response"
	"elichika/generic"
	"elichika/subsystem/user_event/marathon"
	"elichika/userdata"

	"fmt"
)

func FinishEventStory(session *userdata.Session, storyEventMasterId int32, isAutoMode generic.Nullable[bool]) *response.UserModelResponse {
	if isAutoMode.HasValue {
		session.UserStatus.IsAutoMode = isAutoMode.Value
	}
	eventStory := session.Gamedata.EventStory[storyEventMasterId]
	userEvent := marathon.GetUserEventMarathon(session)
	fmt.Println(storyEventMasterId)
	fmt.Println(eventStory)
	fmt.Println(userEvent)
	if eventStory.EventMasterId != userEvent.EventMasterId {
		panic("event changed")
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
