package user_login_bonus

import (
	"elichika/client"
	"elichika/enum"
	"elichika/userdata"
)

func GetBootstrapLoginBonus(session *userdata.Session) client.BootstrapLoginBonus {
	res := client.BootstrapLoginBonus{
		NextLoginBonsReceiveAt: nextLoginBonusTime(session.Time).Unix(),
	}

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTutorialEnd {
		// users in tutorial mode shouldn't get login bonus
		for _, loginBonus := range session.Gamedata.LoginBonus {
			handler[loginBonus.LoginBonusHandler](loginBonus.LoginBonusHandlerConfig, session, loginBonus, &res)
		}
	}

	return res
}
