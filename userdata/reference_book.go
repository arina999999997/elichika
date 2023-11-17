package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) InsertReferenceBook(referenceBookID int) {
	userReferenceBook := model.UserReferenceBook{
		UserID:          session.UserStatus.UserID,
		ReferenceBookID: referenceBookID,
	}
	// guaranteed to be unique
	_, err := session.Db.Table("u_reference_book").Insert(userReferenceBook)
	utils.CheckErr(err)
	session.UserModel.UserReferenceBookByID.PushBack(userReferenceBook)
}

func init() {
	addGenericTableFieldPopulator("u_reference_book", "UserReferenceBookByID")
}
