package userdata

import (
	"elichika/protocol/response"
	"elichika/utils"

	"fmt"
	"time"
)

func (session *Session) GetLoginData() response.Login {
	login := response.Login{}
	exists, err := session.Db.Table("u_login").Where("user_id = ?", session.UserId).Get(&login)
	utils.CheckErr(err)
	if !exists {
		login = response.Login{
			IsPlatformServiceLinked: true,
			LastTimestamp:           time.Now().UnixMilli(),
			Cautions:                []int{},
			CheckMaintenance:        true,
			FromEea:                 false,
		}
		login.ReproInfo.GroupNo = 1
		genericDatabaseInsert(session, "u_login", login)
	}
	return login
}

func (session *Session) UpdateLoginData(login response.Login) {
	affected, err := session.Db.Table("u_login").Where("user_id = ?", session.UserId).AllCols().Update(&login)
	utils.CheckErr(err)
	if affected == 0 {
		genericDatabaseInsert(session, "u_login", login)
	}
}

func (session *Session) Login() response.Login {
	// perform a login, load the relevant data into UserModel and load login data into login
	login := session.GetLoginData()
	login.UserModel = &session.UserModel
	session.UserStatus.LastLoginAt = time.Now().Unix()
	session.SessionType = SessionTypeLogin

	fmt.Println("before")
	for _, populator := range populators {
		populator(session)
	}
	fmt.Println("after")
	// only this part is necessary
	login.MemberLovePanels = session.MemberLovePanels
	return login
}
