package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/login_bonus"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(refactor): Change to use request and response types
func ReadLoginBonus(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ReadLoginBonusRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	resp := SignResp(ctx, "{}", config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func GetBootstrapLoginBonus(ctx *gin.Context, session *userdata.Session) client.BootstrapLoginBonus {
	res := client.BootstrapLoginBonus{
		Event2DLoginBonuses:    []client.IllustLoginBonus{},
		LoginBonuses:           []client.NaviLoginBonus{},
		Event3DLoginBonus:      []client.NaviLoginBonus{},
		BeginnerLoginBonuses:   []client.NaviLoginBonus{},
		ComebackLoginBonuses:   []client.IllustLoginBonus{},
		BirthdayLoginBonuses:   []client.NaviLoginBonus{},
		BirthdayMember:         []client.LoginBonusBirthDayMember{},
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
