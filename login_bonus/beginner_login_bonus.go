package login_bonus

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/userdata"
)

func beginnerLoginBonusHandler(_ string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *client.BootstrapLoginBonus) {
	if loginBonus.LoginBonusType != enum.LoginBonusTypeNormal {
		panic("wrong handler used")
	}
	userLoginBonus := session.GetUserLoginBonus(loginBonus.LoginBonusId)
	lastUnlocked := latestLoginBonusTime(session.Time)
	if userLoginBonus.LastReceivedAt >= lastUnlocked.Unix() { // already got it
		return
	}

	userLoginBonus.LastReceivedAt = session.Time.Unix()
	userLoginBonus.LastReceivedReward++
	if userLoginBonus.LastReceivedReward >= len(loginBonus.LoginBonusRewards) { // already received everything
		return
	}
	naviLoginBonus := loginBonus.NaviLoginBonus()
	for i := range naviLoginBonus.LoginBonusRewards {
		if i < userLoginBonus.LastReceivedReward {
			naviLoginBonus.LoginBonusRewards[i].Status = enum.LoginBonusReceiveStatusReceived
		} else if i > userLoginBonus.LastReceivedReward {
			naviLoginBonus.LoginBonusRewards[i].Status = enum.LoginBonusReceiveStatusUnreceived
		} else {
			naviLoginBonus.LoginBonusRewards[i].Status = enum.LoginBonusReceiveStatusReceiving
		}
	}
	target.LoginBonuses = append(target.LoginBonuses, naviLoginBonus)
	for _, content := range loginBonus.LoginBonusRewards[userLoginBonus.LastReceivedReward].LoginBonusContents {
		// TODO(present_box): This correctly has to go to the present box, but we just do it here
		session.AddResource(content)
	}
	session.UpdateUserLoginBonus(userLoginBonus)
}
