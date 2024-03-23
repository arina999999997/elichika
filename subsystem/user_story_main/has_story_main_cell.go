package user_story_main

import (
	"elichika/client"
	"elichika/userdata"
)

func hasStoryMainCell(session *userdata.Session, storyMainMasterId int32) bool {
	userStoryMain := client.UserStoryMain{
		StoryMainMasterId: storyMainMasterId,
	}
	return userdata.GenericDatabaseExist(session, "u_story_main", userStoryMain)
}
