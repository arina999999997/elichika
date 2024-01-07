package userdata

import (
	"elichika/model"
)

func (session *Session) UnlockEventStory(eventStoryMasterId int) {
	userStoryEventHistory := model.UserStoryEventHistory{
		StoryEventId: eventStoryMasterId,
	}
	session.UserModel.UserStoryEventHistoryById.PushBack(userStoryEventHistory)
}

func eventStoryFinalizer(session *Session) {
	for _, userStoryEventHistory := range session.UserModel.UserStoryEventHistoryById.Objects {
		genericDatabaseInsert(session, "u_story_event_history", userStoryEventHistory)
	}
}

func init() {
	addFinalizer(eventStoryFinalizer)
	addGenericTableFieldPopulator("u_story_event_history", "UserStoryEventHistoryById")
}
