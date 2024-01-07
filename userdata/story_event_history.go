package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) UnlockEventStory(eventStoryMasterId int) {
	userStoryEventHistory := model.UserStoryEventHistory{
		UserId:       session.UserStatus.UserId,
		StoryEventId: eventStoryMasterId,
	}
	session.UserModel.UserStoryEventHistoryById.PushBack(userStoryEventHistory)
}

func eventStoryFinalizer(session *Session) {
	for _, userStoryEventHistory := range session.UserModel.UserStoryEventHistoryById.Objects {
		_, err := session.Db.Table("u_story_event_history").Insert(userStoryEventHistory)
		utils.CheckErr(err)
	}
}

func init() {
	addFinalizer(eventStoryFinalizer)
	addGenericTableFieldPopulator("u_story_event_history", "UserStoryEventHistoryById")
}
