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
		genericDatabaseInsert(session, "u_story_event_history", *userStoryEventHistory)
	}
}

func init() {
	addFinalizer(eventStoryFinalizer)
	addGenericTableFieldPopulator("u_story_event_history", "UserStoryEventHistoryById")
}
