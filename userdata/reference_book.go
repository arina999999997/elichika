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
		GenericDatabaseInsert(session, "u_reference_book", *userReferenceBook)
	}
}

func init() {
	AddFinalizer(referenceBookFinalizer)
}
