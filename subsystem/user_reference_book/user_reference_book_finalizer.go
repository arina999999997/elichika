package user_reference_book

import (
	"elichika/userdata"
)

func userReferenceBookFinalizer(session *userdata.Session) {
	// guaranteed to be unique
	for _, userReferenceBook := range session.UserModel.UserReferenceBookById.Map {
		userdata.GenericDatabaseInsert(session, "u_reference_book", *userReferenceBook)
	}
}

func init() {
	userdata.AddFinalizer(userReferenceBookFinalizer)
}
