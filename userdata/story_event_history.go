package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UnlockEventStory(eventStoryMasterID int) {
	userStoryEventHistory := model.UserStoryEventHistory{
		UserID:       session.UserStatus.UserID,
		StoryEventID: eventStoryMasterID,
	}
	session.UserModel.UserStoryEventHistoryByID.PushBack(userStoryEventHistory)
}

func eventStoryFinalizer(session *Session) {
	for _, userStoryEventHistory := range session.UserModel.UserStoryEventHistoryByID.Objects {
		_, err := session.Db.Table("u_story_event_history").Insert(userStoryEventHistory)
		utils.CheckErr(err)
	}
}

func init() {
	addFinalizer(eventStoryFinalizer)
	addGenericTableFieldPopulator("u_story_event_history", "UserStoryEventHistoryByID")
}
