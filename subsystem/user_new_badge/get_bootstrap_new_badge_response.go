package user_new_badge

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

// TODO(new_badge): This is fetched from database, but it's not actually updated
func GetBootstrapNewBadgeResponse(session *userdata.Session) client.BootstrapNewBadge {
	newBadge := client.BootstrapNewBadge{}
	exist, err := session.Db.Table("u_new_badge").Where("user_id = ?", session.UserId).Get(&newBadge)
	utils.CheckErr(err)
	if !exist {
		newBadge = client.BootstrapNewBadge{
			IsNewMainStory:                     false,
			UnreceivedPresentBox:               0,
			IsUnreceivedPresentBoxSubscription: false,
			IsUpdateFriend:                     false,
			UnreceivedMission:                  0,
			UnreceivedChallengeBeginner:        0,
		}
	}
	return newBadge
}
