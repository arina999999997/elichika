package userdata

import (
	"elichika/client"
)

func (session *Session) UnlockEventStory(eventStoryMasterId int32) {
	userStoryEventHistory := client.UserStoryEventHistory{
		StoryEventId: eventStoryMasterId,
	}
	session.UserModel.UserStoryEventHistoryById.Set(eventStoryMasterId, userStoryEventHistory)
}

func eventStoryFinalizer(session *Session) {
	for _, userStoryEventHistory := range session.UserModel.UserStoryEventHistoryById.Map {
		GenericDatabaseInsert(session, "u_story_event_history", *userStoryEventHistory)
	}
}

func init() {
	AddContentFinalizer(eventStoryFinalizer)
}
