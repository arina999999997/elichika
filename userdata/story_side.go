package userdata

import (
	"elichika/client"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func (session *Session) InsertStorySide(storySideMasterId int32) {
	userStorySide := client.UserStorySide{
		StorySideMasterId: storySideMasterId,
		IsNew:             true,
		AcquiredAt:        session.Time.Unix(),
	}
	session.UserModel.UserStorySideById.Set(storySideMasterId, userStorySide)
}

func (session *Session) FinishStorySide(storySideMasterId int32) {
	userStorySide := client.UserStorySide{}
	exist, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
		session.UserId, storySideMasterId).Get(&userStorySide)
	utils.CheckErr(err)
	if !exist {
		panic("side story must exist to be read")
	}
	if !userStorySide.IsNew { // already read
		return
	}
	userStorySide.IsNew = false
	session.UserModel.UserStorySideById.Set(storySideMasterId, userStorySide)
}

func storySideFinalizer(session *Session) {
	for _, userStorySide := range session.UserModel.UserStorySideById.Map {
		affected, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
			session.UserId, userStorySide.StorySideMasterId).AllCols().Update(userStorySide)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			GenericDatabaseInsert(session, "u_story_side", *userStorySide)
		}
	}
}

func init() {
	AddFinalizer(storySideFinalizer)
}
