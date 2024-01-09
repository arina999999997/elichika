package userdata

import (
	"elichika/client"
)

func (session *Session) InsertReferenceBook(referenceBookId int32) {
	userReferenceBook := client.UserReferenceBook{
		ReferenceBookId: referenceBookId,
	}
	session.UserModel.UserReferenceBookById.PushBack(userReferenceBook)
}

func referenceBookFinalizer(session *Session) {
	// guaranteed to be unique
	for _, userReferenceBook := range session.UserModel.UserReferenceBookById.Objects {
		genericDatabaseInsert(session, "u_reference_book", userReferenceBook)
	}
}

func init() {
	addFinalizer(referenceBookFinalizer)
	addGenericTableFieldPopulator("u_reference_book", "UserReferenceBookById")
}
