package user_live_partner

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserLivePartner(session *userdata.Session, otherUserId int32) client.LivePartner {
	partner := client.LivePartner{}

	exist, err := session.Db.Table("u_status").Where("user_id = ?", otherUserId).Get(&partner)
	utils.CheckErrMustExist(err, exist)

	partner.IsFriend = true
	for i := int32(1); i <= 7; i++ {
		partnerCard := GetOtherUserProfileLivePartnerCard(session, otherUserId, i)
		if partnerCard.PartnerCard.CardMasterId != 0 {
			partner.CardByCategory.Set(i, partnerCard.PartnerCard)
		}
	}
	return partner
}
