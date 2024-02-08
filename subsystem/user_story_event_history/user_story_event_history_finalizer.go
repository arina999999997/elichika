package user_story_event_history

import (
	"elichika/userdata"
)

func userStoryEventHistoryFinalizer(session *userdata.Session) {
	// this can only happen if the state is wrong, so no need to check for existing items
	for _, userStoryEventHistory := range session.UserModel.UserStoryEventHistoryById.Map {
		userdata.GenericDatabaseInsert(session, "u_story_event_history", *userStoryEventHistory)
	}
}

func init() {
	userdata.AddFinalizer(userStoryEventHistoryFinalizer)
}
