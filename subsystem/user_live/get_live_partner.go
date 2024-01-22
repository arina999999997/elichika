package user_live

import (
	"elichika/client"
	"elichika/subsystem/user_card"
	"elichika/userdata"
)

func GetLivePartner(session *userdata.Session, otherUserId int32) client.LivePartner {
	partner := client.LivePartner{}
	userdata.FetchDBProfile(otherUserId, &partner)
	partner.IsFriend = true
	for i := int32(1); i <= 7; i++ {
		partnerCard := user_card.GetOtherUserProfileLivePartnerCard(session, otherUserId, i)
		if partnerCard.PartnerCard.CardMasterId != 0 {
			partner.CardByCategory.Set(i, partnerCard.PartnerCard)
		}
	}
	return partner
}
