package login_bonus

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"

	"time"
)

// the latest login bonus that can be claimed
func latestLoginBonusTime(timePoint time.Time) time.Time {
	year, month, day := timePoint.Date()
	res := time.Date(year, month, day, 0, 0, *config.Conf.LoginBonusSecond, 0, timePoint.Location())
	if res.After(timePoint) {
		res = res.AddDate(0, 0, -1)
	}
	return res
}
func NextLoginBonusTime(timePoint time.Time) time.Time {
	return latestLoginBonusTime(timePoint).AddDate(0, 0, 1)
}

func normalLoginBonusHandler(_ string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *model.BootstrapLoginBonus) {
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
	if userLoginBonus.LastReceivedReward == len(loginBonus.LoginBonusRewards) {
		userLoginBonus.LastReceivedReward = 0
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