package userdata

import (
	"elichika/model"
	"elichika/utils"

	"time"
)

func (session *Session) InsertStorySide(storySideMasterID int) {
	userStorySide := model.UserStorySide{
		UserID:            session.UserStatus.UserID,
		StorySideMasterID: storySideMasterID,
		IsNew:             true,
		AcquiredAt:        time.Now().Unix(),
	}
	session.UserModel.UserStorySideByID.PushBack(userStorySide)
}

func (session *Session) FinishStorySide(storySideMasterID int) {
	userStorySide := model.UserStorySide{}
	exist, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
		session.UserStatus.UserID, storySideMasterID).Get(&userStorySide)
	utils.CheckErr(err)
	if !exist {
		panic("side story must exist to be read")
	}
	if !userStorySide.IsNew { // already read
		return
	}
	userStorySide.IsNew = false
	session.UserModel.UserStorySideByID.PushBack(userStorySide)
}

func storySideFinalizer(session *Session) {
	for _, userStorySide := range session.UserModel.UserStorySideByID.Objects {
		affected, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
			userStorySide.UserID, userStorySide.StorySideMasterID).AllCols().Update(userStorySide)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			_, err = session.Db.Table("u_story_side").Insert(userStorySide)
			utils.CheckErr(err)
		}

	}
}

func init() {
	addGenericTableFieldPopulator("u_story_side", "UserStorySideByID")
	addFinalizer(storySideFinalizer)
}
