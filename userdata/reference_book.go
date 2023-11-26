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
	session.UserModel.UserReferenceBookByID.PushBack(userReferenceBook)
}

func referenceBookFinalizer(session *Session) {
	// guaranteed to be unique
	for _, userReferenceBook := range session.UserModel.UserReferenceBookByID.Objects {
		_, err := session.Db.Table("u_reference_book").Insert(userReferenceBook)
		utils.CheckErr(err)
	}
}

func init() {
	addFinalizer(referenceBookFinalizer)
	addGenericTableFieldPopulator("u_reference_book", "UserReferenceBookByID")
}
