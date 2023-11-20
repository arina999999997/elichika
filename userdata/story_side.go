package userdata

import (
	"elichika/model"
	"elichika/utils"

	"fmt"
	"time"
)

func (session *Session) InsertStorySide(storySideMasterID int) {
	userStorySide := model.UserStorySide{
		UserID:            session.UserStatus.UserID,
		StorySideMasterID: storySideMasterID,
		IsNew:             true,
		AcquiredAt:        time.Now().Unix(),
	}

	_, err := session.Db.Table("u_story_side").Insert(userStorySide)
	session.UserModel.UserStorySideByID.PushBack(userStorySide)
	utils.CheckErr(err)
}

func (session *Session) FinishStorySide(storySideMasterID int) {
	userStorySide := model.UserStorySide{}
	exists, err := session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
		session.UserStatus.UserID, storySideMasterID).Get(&userStorySide)
	utils.CheckErr(err)
	if !exists {
		panic("side story must exist to be read")
	}
	if !userStorySide.IsNew { // already read
		return
	}
	userStorySide.IsNew = false
	_, err = session.Db.Table("u_story_side").Where("user_id = ? AND story_side_master_id = ?",
		userStorySide.UserID, userStorySide.StorySideMasterID).AllCols().Update(userStorySide)
	utils.CheckErr(err)
	fmt.Println(userStorySide)
	session.UserModel.UserStorySideByID.PushBack(userStorySide)
}

func init() {
	addGenericTableFieldPopulator("u_story_side", "UserStorySideByID")
}
