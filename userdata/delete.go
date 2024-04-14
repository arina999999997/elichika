package userdata

import (
	"elichika/userdata/database"
)

// delete userdata for the GAME, the account password and authentication are still kept
// this is so user can start from "fresh" or import old data while keeping the user id
func (session *Session) DeleteUserGameData() {
	for table := range database.UserDataTableNameToInterface {
		if table == "u_pass_word" || table == "u_authentication" {
			continue
		}
		session.Db.Table(table).Where("user_id = ?", session.UserId).Delete()
	}
}
