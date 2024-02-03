package user_new_badge

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func SetBootstrapNewBadgeResponse(session *userdata.Session, newBadge client.BootstrapNewBadge) {
	affected, err := session.Db.Table("u_new_badge").Where("user_id = ?", session.UserId).Update(&newBadge)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_new_badge", newBadge)
	}
}
