package login_bonus

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/subsystem/user_present"
	"elichika/userdata"

	"fmt"
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
func nextLoginBonusTime(timePoint time.Time) time.Time {
	return latestLoginBonusTime(timePoint).AddDate(0, 0, 1)
}

func normalLoginBonusHandler(_ string, session *userdata.Session, loginBonus *gamedata.LoginBonus, target *client.BootstrapLoginBonus) {
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
	if userLoginBonus.LastReceivedReward == loginBonus.LoginBonusRewards.Size() {
		userLoginBonus.LastReceivedReward = 0
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
			PresentRouteId:   generic.NewNullable(int32(1000002)), // this doesn't really matter much even though it's sent
			// TODO(localization): This is not localized to the correct language
			ParamServer: generic.NewNullable(client.LocalizedText{
				DotUnderText: "Daily Login Bonus",
			}),
			ParamClient: generic.NewNullable(fmt.Sprint(userLoginBonus.LastReceivedReward + 1)),
		})
	}
	session.UpdateUserLoginBonus(userLoginBonus)
}
