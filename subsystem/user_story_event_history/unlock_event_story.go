package user_story_event_history

import (
	"elichika/client"
	"elichika/userdata"
)

func UnlockEventStory(session *userdata.Session, eventStoryMasterId int32) {
	userStoryEventHistory := client.UserStoryEventHistory{
		StoryEventId: eventStoryMasterId,
	}
	if userdata.GenericDatabaseExist(session, "u_story_event_history", userStoryEventHistory) {
		return
	}
	session.UserModel.UserStoryEventHistoryById.Set(eventStoryMasterId, userStoryEventHistory)
}
