package user_authentication

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"

	"golang.org/x/crypto/bcrypt"
)

func SetPassWord(session *userdata.Session, passWord string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passWord), bcryptCost)
	utils.CheckErr(err)
	affected, err := session.Db.Table("u_pass_word").Where("user_id = ?", session.UserId).AllCols().Update(
		database.UserPassWord{
			Hash: hash,
		})
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_pass_word", database.UserPassWord{
			Hash: hash,
		})
	}
}
