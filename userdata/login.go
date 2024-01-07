package userdata

import (
	"elichika/protocol/response"
	"elichika/utils"

	"time"
)

func (session *Session) Login() response.Login {
	// perform a login, load the relevant data into UserModel and load login data into login
	login := response.Login{}
	exists, err := session.Db.Table("u_login").Where("user_id = ?", session.UserId).Get(&login)
	utils.CheckErr(err)
	if !exists {
		login = response.Login{
			IsPlatformServiceLinked: true,
			LastTimestamp:           time.Now().UnixMilli(),
			Cautions:                []int{},
			CheckMaintenance:        true,
			FromEea:                 true,
		}
		login.ReproInfo.GroupNo = 1
		genericDatabaseInsert(session, "u_login", login)
	}

	login.UserModel = &session.UserModel
	session.UserStatus.LastLoginAt = time.Now().Unix()
	session.SessionType = SessionTypeLogin

	for _, populator := range populators {
		populator(session)
	}
	// only this part is necessary
	login.MemberLovePanels = session.UserMemberLovePanels
	return login
}
