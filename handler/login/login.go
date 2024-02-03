package login

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_account"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func login(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LoginRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session == nil {
		user_account.CreateNewAccount(ctx, userId, "")
		session = userdata.GetSession(ctx, userId)
		defer session.Close()
	}

	fmt.Println("User logins: ", userId)

	resp := session.Login()
	resp.SessionKey = LoginSessionKey(req.Mask)
	{
		exist, _, startLiveRequest := session.LoadUserLive()
		if exist {
			liveDifficulty := session.Gamedata.LiveDifficulty[startLiveRequest.LiveDifficultyId]
			if (liveDifficulty.UnlockPattern != enum.LiveUnlockPatternCoopOnly) &&
				(liveDifficulty.UnlockPattern != enum.LiveUnlockPatternTowerOnly) {
				resp.LiveResume = generic.NewNullable(client.LiveResume{
					LiveDifficultyId: startLiveRequest.LiveDifficultyId,
					DeckId:           startLiveRequest.DeckId,
					ConsumedLp:       liveDifficulty.ConsumedLP, // this thing is only to show how much lp is spent
				})
			} else { // just cancel this as it's not a relevant live (event and such)
				session.ClearUserLive()
			}
		}
	}
	session.Finalize()
	common.JsonResponse(ctx, resp)

	{
		backupText, err := json.Marshal(resp)
		utils.CheckErr(err)
		utils.WriteAllText(fmt.Sprint(config.UserDataBackupPath, "login_", userId, ".json"), string(backupText))
	}
}

func init() {
	router.AddHandler("/login/login", login)
}
