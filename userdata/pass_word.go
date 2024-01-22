package userdata

import (
	"elichika/userdata/database"
	"elichika/utils"
)

func (session *Session) SetPassWord(passWord string) {
	affected, err := session.Db.Table("u_pass_word").Where("user_id = ?", session.UserId).AllCols().Update(
		database.UserPassWord{
			PassWord: passWord,
		})
	utils.CheckErr(err)
	if affected == 0 {
		genericDatabaseInsert(session, "u_pass_word", database.UserPassWord{
			PassWord: passWord,
		})
	}
}

func (session *Session) CheckPassWord(passWord string) bool {
	userPassWord := database.UserPassWord{}
	exist, err := session.Db.Table("u_pass_word").Where("user_id = ?", session.UserId).Get(&userPassWord)
	utils.CheckErr(err)
	if !exist {
		session.SetPassWord(passWord)
		return true
	}
	return userPassWord.PassWord == passWord
}
