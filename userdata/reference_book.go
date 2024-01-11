package userdata

import (
	"elichika/client"
)

func (session *Session) InsertReferenceBook(referenceBookId int32) {
	session.UserModel.UserReferenceBookById.Set(referenceBookId, client.UserReferenceBook{
		ReferenceBookId: referenceBookId,
	})
}

func referenceBookFinalizer(session *Session) {
	// guaranteed to be unique
	for _, userReferenceBook := range session.UserModel.UserReferenceBookById.Map {
		genericDatabaseInsert(session, "u_reference_book", userReferenceBook)
	}
}

func init() {
	addFinalizer(referenceBookFinalizer)
	addGenericTableFieldPopulator("u_reference_book", "UserReferenceBookById")
}
