package user_accessory

import (
	"elichika/userdata"
	"elichika/utils"
)

func ClearIsNewFlags(session* userdata.Session) {
	_, err := session.Db.Exec("UPDATE u_accessory SET is_new=0 WHERE user_id=? AND is_new=1", session.UserId)
	utils.CheckErr(err)
}