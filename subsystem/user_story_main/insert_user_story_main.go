package user_story_main

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertUserStoryMain(session *userdata.Session, storyMainMasterId int32) bool {
	userStoryMain := client.UserStoryMain{
		StoryMainMasterId: storyMainMasterId,
	}
	if userdata.GenericDatabaseExist(session, "u_story_main", userStoryMain) {
		return false
	}
	session.UserModel.UserStoryMainByStoryMainId.Set(storyMainMasterId, userStoryMain)
	// main story also used to unlock scenes (feature), but they are unlocked by default now
	return true
}
