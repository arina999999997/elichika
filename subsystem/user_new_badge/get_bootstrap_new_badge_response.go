package user_new_badge

import (
	"elichika/client"
	"elichika/subsystem/user_mission"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_social"
	"elichika/userdata"
)

func GetBootstrapNewBadgeResponse(session *userdata.Session) client.BootstrapNewBadge {
	return client.BootstrapNewBadge{
		IsNewMainStory:                     false,
		UnreceivedPresentBox:               user_present.FetchPresentCount(session),
		IsUnreceivedPresentBoxSubscription: false, // TODO(present box, subscription)
		IsUpdateFriend:                     user_social.IsUpdateFriend(session),
		UnreceivedMission:                  user_mission.CountUnreceivedMission(session),
		UnreceivedChallengeBeginner:        0, // TODO(beginner guide)
	}
}
