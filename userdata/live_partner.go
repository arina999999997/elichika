package userdata

import (
	"elichika/userdata/database"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func (session *Session) GetUserLivePartner(livePartnerCategoryId int32) database.UserLivePartner {
	userLivePartner := database.UserLivePartner{}
	exist, err := session.Db.Table("u_live_partner").
		Where("user_id = ? AND live_partner_category_id = ?", session.UserId, livePartnerCategoryId).
		Get(&userLivePartner)
	utils.CheckErr(err)
	if !exist {
		userLivePartner = database.UserLivePartner{
			LivePartnerCategoryId: livePartnerCategoryId,
			CardMasterId:          0,
		}
	}
	return userLivePartner
}

func (session *Session) UpdateUserLivePartner(userLivePartner database.UserLivePartner) {
	if userLivePartner.CardMasterId == 0 {
		return
	}
	affected, err := session.Db.Table("u_live_partner").
		Where("user_id = ? AND live_partner_category_id = ?", session.UserId, userLivePartner.LivePartnerCategoryId).
		AllCols().Update(&userLivePartner)
	utils.CheckErr(err)
	if affected == 0 {
		GenericDatabaseInsert(session, "u_live_partner", userLivePartner)
	}
}
