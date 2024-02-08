package user_reference_book

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertUserReferenceBook(session *userdata.Session, referenceBookId int32) {
	session.UserModel.UserReferenceBookById.Set(referenceBookId, client.UserReferenceBook{
		ReferenceBookId: referenceBookId,
	})
}
