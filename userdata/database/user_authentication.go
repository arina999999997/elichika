package database

import (
	"elichika/generic"
)

// TODO(password): use bcrypt or something
type UserPassWord struct {
	PassWord string
}

func init() {
	AddTable("u_pass_word", generic.UserIdWrapper[UserPassWord]{})
}
