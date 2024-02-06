package user_live_partner

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func SetLivePartnerCard(session *userdata.Session, livePartnerCategoryId, cardMasterId int32) {
	userLivePartnerCard := database.UserLivePartnerCard{
		LivePartnerCategoryId: livePartnerCategoryId,
		CardMasterId:          cardMasterId,
	}
	affected, err := session.Db.Table("u_live_partner_card").
		Where("user_id = ? AND live_partner_category_id = ?", session.UserId, livePartnerCategoryId).
		AllCols().Update(&userLivePartnerCard)
	utils.CheckErr(err)

	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_live_partner_card", userLivePartnerCard)
	}
}
