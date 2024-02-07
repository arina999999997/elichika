package database

import (
	"elichika/generic"
)

type UserPassWord struct {
	Hash []byte
}

func init() {
	AddTable("u_pass_word", generic.UserIdWrapper[UserPassWord]{})
}
