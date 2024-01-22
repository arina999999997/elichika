package user_card

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserProfileLivePartnerCard(session *userdata.Session, otherUserId, livePartnerCategoryMasterId int32) client.ProfileLivePartnerCard {
	res := client.ProfileLivePartnerCard{
		LivePartnerCategoryMasterId: livePartnerCategoryMasterId,
	}

	var cardMasterId int32
	exist, err := session.Db.Table("u_live_partner").
		Where("user_id = ? AND live_partner_category_id = ?", otherUserId, livePartnerCategoryMasterId).
		Cols("card_master_id").Get(&cardMasterId)
	utils.CheckErr(err)
	if !exist {
		return res
	}
	res.PartnerCard = GetOtherUserCard(session, otherUserId, cardMasterId)
	return res
}
