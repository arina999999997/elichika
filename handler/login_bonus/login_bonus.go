package login_bonus

import (
	"elichika/client"
	// "elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/login_bonus"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func ReadLoginBonus(ctx *gin.Context) {
	// this doesn't need to do anything
	// reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	// req := request.ReadLoginBonusRequest{}
	// err := json.Unmarshal([]byte(reqBody), &req)
	// utils.CheckErr(err)
	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func GetBootstrapLoginBonus(ctx *gin.Context, session *userdata.Session) client.BootstrapLoginBonus {
	res := client.BootstrapLoginBonus{
		NextLoginBonsReceiveAt: login_bonus.NextLoginBonusTime(session.Time).Unix(),
	}

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTutorialEnd {
		// users in tutorial mode shouldn't get login bonus
		for _, loginBonus := range session.Gamedata.LoginBonus {
			handler := login_bonus.Handler[loginBonus.LoginBonusHandler]
			handler(loginBonus.LoginBonusHandlerConfig, session, loginBonus, &res)
		}
	}

	return res
}

func init() {
	router.AddHandler("/loginBonus/readLoginBonus", ReadLoginBonus)
}
