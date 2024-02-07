package user_login_bonus

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/subsystem/user_present"
	"elichika/userdata"

	"fmt"
)

func beginnerLoginBonusHandler(_ string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *client.BootstrapLoginBonus) {
	if loginBonus.LoginBonusType != enum.LoginBonusTypeNormal {
		panic("wrong handler used")
	}
	userLoginBonus := getUserLoginBonus(session, loginBonus.LoginBonusId)
	lastUnlocked := latestLoginBonusTime(session.Time)
	if userLoginBonus.LastReceivedAt >= lastUnlocked.Unix() { // already got it
		return
	}

	userLoginBonus.LastReceivedAt = session.Time.Unix()
	userLoginBonus.LastReceivedReward++
	if userLoginBonus.LastReceivedReward >= loginBonus.LoginBonusRewards.Size() { // already received everything
		return
	}
	naviLoginBonus := loginBonus.NaviLoginBonus()
	for i := range naviLoginBonus.LoginBonusRewards.Slice {
		if i < userLoginBonus.LastReceivedReward {
			naviLoginBonus.LoginBonusRewards.Slice[i].Status = enum.LoginBonusReceiveStatusReceived
		} else if i > userLoginBonus.LastReceivedReward {
			naviLoginBonus.LoginBonusRewards.Slice[i].Status = enum.LoginBonusReceiveStatusUnreceived
		} else {
			naviLoginBonus.LoginBonusRewards.Slice[i].Status = enum.LoginBonusReceiveStatusReceiving
		}
	}
	target.LoginBonuses.Append(naviLoginBonus)
	for _, content := range loginBonus.LoginBonusRewards.Slice[userLoginBonus.LastReceivedReward].LoginBonusContents.Slice {
		user_present.AddPresent(session, client.PresentItem{
			Content:          content,
			PresentRouteType: enum.PresentRouteTypeLoginBonus,
			PresentRouteId:   generic.NewNullable(int32(1000001)),
			// TODO(localization): This is not localized to the correct language, also we don't even know if this is the correct string
			ParamServer: generic.NewNullable(client.LocalizedText{
				DotUnderText: "Beginner Login Bonus",
			}),
			ParamClient: generic.NewNullable(fmt.Sprint(userLoginBonus.LastReceivedReward + 1)),
		})
	}
	updateUserLoginBonus(session, userLoginBonus)
}
