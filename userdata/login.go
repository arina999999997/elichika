package userdata

import (
	// "fmt"
	"time"
)

func (session *Session) Login() {
	// perform a login, load the relevant data into user model common
	session.UserStatus.LastLoginAt = time.Now().Unix()
	for _, populator := range populators {
		populator(session)
	}
	// fmt.Println(session.UserModel)
}
