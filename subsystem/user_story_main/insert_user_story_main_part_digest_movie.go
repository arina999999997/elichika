package user_story_main

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertUserStoryMainPartDigestMovie(session *userdata.Session, partId int32) {
	session.UserModel.UserStoryMainPartDigestMovieById.Set(partId, client.UserStoryMainPartDigestMovie{
		StoryMainPartMasterId: partId,
	})
}
