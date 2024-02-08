package user_story_side

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertStorySide(session *userdata.Session, storySideMasterId int32) {
	userStorySide := client.UserStorySide{
		StorySideMasterId: storySideMasterId,
		IsNew:             true,
		AcquiredAt:        session.Time.Unix(),
	}
	session.UserModel.UserStorySideById.Set(storySideMasterId, userStorySide)
}
