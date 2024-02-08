package user_story_linkage

import (
	"elichika/userdata"
)

func userStoryLinkageFinalizer(session *userdata.Session) {
	for _, userStoryLinkage := range session.UserModel.UserStoryLinkageById.Map {
		if !userdata.GenericDatabaseExist(session, "u_story_linkage", *userStoryLinkage) {
			userdata.GenericDatabaseInsert(session, "u_story_linkage", *userStoryLinkage)
		}
	}
}

func init() {
	userdata.AddFinalizer(userStoryLinkageFinalizer)
}
