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

	_, err := session.Db.Table("u_story_event_history").Insert(userStoryEventHistory)
	session.UserModel.UserStoryEventHistoryByID.PushBack(userStoryEventHistory)
	utils.CheckErr(err)
}

func init() {
	addGenericTableFieldPopulator("u_story_event_history", "UserStoryEventHistoryByID")
}
