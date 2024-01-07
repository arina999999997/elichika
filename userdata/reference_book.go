package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) InsertReferenceBook(referenceBookId int) {
	userReferenceBook := model.UserReferenceBook{
		UserId:          session.UserStatus.UserId,
		ReferenceBookId: referenceBookId,
	}
	session.UserModel.UserReferenceBookById.PushBack(userReferenceBook)
}

func referenceBookFinalizer(session *Session) {
	// guaranteed to be unique
	for _, userReferenceBook := range session.UserModel.UserReferenceBookById.Objects {
		_, err := session.Db.Table("u_reference_book").Insert(userReferenceBook)
		utils.CheckErr(err)
	}
}

func init() {
	addFinalizer(referenceBookFinalizer)
	addGenericTableFieldPopulator("u_reference_book", "UserReferenceBookById")
}
