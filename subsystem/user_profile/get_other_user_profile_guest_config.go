package user_profile

import (
	"elichika/client"
	"elichika/subsystem/user_card"
	"elichika/userdata"
)

func GetOtherUserProfileGuestConfig(session *userdata.Session, otherUserId int32) client.ProfileGuestConfig {
	res := client.ProfileGuestConfig{}
	for i := int32(1); i <= 7; i++ {
		res.LivePartnerCards.Append(user_card.GetOtherUserProfileLivePartnerCard(session, otherUserId, i))
	}
	return res
}
