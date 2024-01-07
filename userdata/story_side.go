package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) InsertStorySide(storySideMasterId int) {
	userStorySide := model.UserStorySide{
		UserId:            session.UserStatus.UserId,
		StorySideMasterId: storySideMasterId,
		IsNew:             true,
		AcquiredAt:        session.Time.Unix(),
	}
	session.UserModel.UserStorySideById.PushBack(userStorySide)
}

func (session *Session) FinishStorySide(storySideMasterId int) {
	userStorySide := model.UserStorySide{}
	exist, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
		session.UserStatus.UserId, storySideMasterId).Get(&userStorySide)
	utils.CheckErr(err)
	if !exist {
		panic("side story must exist to be read")
	}
	if !userStorySide.IsNew { // already read
		return
	}
	userStorySide.IsNew = false
	session.UserModel.UserStorySideById.PushBack(userStorySide)
}

func storySideFinalizer(session *Session) {
	for _, userStorySide := range session.UserModel.UserStorySideById.Objects {
		affected, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
			userStorySide.UserId, userStorySide.StorySideMasterId).AllCols().Update(userStorySide)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			_, err = session.Db.Table("u_story_side").Insert(userStorySide)
			utils.CheckErr(err)
		}

	}
}

func init() {
	addGenericTableFieldPopulator("u_story_side", "UserStorySideById")
	addFinalizer(storySideFinalizer)
}
