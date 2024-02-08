package user_story_main

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserStoryMainSelected(session *userdata.Session, storyMainCellId, selectedId int32) {
	userStoryMainSelected := client.UserStoryMainSelected{
		StoryMainCellId: storyMainCellId,
		SelectedId:      selectedId,
	}
	session.UserModel.UserStoryMainSelectedByStoryMainCellId.Set(storyMainCellId, userStoryMainSelected)
}