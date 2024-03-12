package database

import (
	"elichika/generic"
)

type UserPassWord struct {
	Hash []byte
}

type UserAuthentication struct {
	AuthorizationKey   []byte // generated at account creation and never change
	AuthorizationCount int32  // number of login
	SessionKey         []byte // generated at login for a session
	CommandId          int32  // current command id
}

func init() {
	AddTable("u_pass_word", generic.UserIdWrapper[UserPassWord]{})
	AddTable("u_authentication", generic.UserIdWrapper[UserAuthentication]{})
}
