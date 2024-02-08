package user_story_main

import (
	"elichika/userdata"
)

func userStoryMainFinalizer(session *userdata.Session) {
	for _, userStoryMain := range session.UserModel.UserStoryMainByStoryMainId.Map {
		if !userdata.GenericDatabaseExist(session, "u_story_main", *userStoryMain) {
			userdata.GenericDatabaseInsert(session, "u_story_main", *userStoryMain)
		}
	}
}
func init() {
	userdata.AddFinalizer(userStoryMainFinalizer)
}
