package user_social

import (
	"elichika/client"
	"elichika/userdata"
)

func GetOtherUserProfileGuestConfig(session *userdata.Session, otherUserId int32) client.ProfileGuestConfig {
	res := client.ProfileGuestConfig{}
	for i := int32(1); i <= 7; i++ {
		res.LivePartnerCards.Append(GetOtherUserProfileLivePartnerCard(session, otherUserId, i))
	}
	return res
}
