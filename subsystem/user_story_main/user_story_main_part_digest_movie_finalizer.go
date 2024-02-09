package user_story_main

import (
	"elichika/userdata"
)

func userStoryMainPartDigestMovieFinalizer(session *userdata.Session) {
	for _, userStoryMainPartDigestMovie := range session.UserModel.UserStoryMainPartDigestMovieById.Map {
		if !userdata.GenericDatabaseExist(session, "u_story_main_part_digest_movie", *userStoryMainPartDigestMovie) {
			userdata.GenericDatabaseInsert(session, "u_story_main_part_digest_movie", *userStoryMainPartDigestMovie)
		}
	}
}

func init() {
	userdata.AddFinalizer(userStoryMainPartDigestMovieFinalizer)
}
