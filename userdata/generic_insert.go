package userdata

import (
	"elichika/utils"
)

func genericDatabaseInsert[T any](session *Session, table string, item T) {
	type UserIdWrapper struct {
		UserId int `xorm:"pk 'user_id'"`
		Item   *T  `xorm:"extends"`
	}
	_, err := session.Db.Table(table).Insert(UserIdWrapper{
		UserId: session.UserId,
		Item:   &item,
	})
	utils.CheckErr(err)
}