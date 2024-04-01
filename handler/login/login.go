package login

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	req := request.LoginRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	ctx.Set("sign_key", session.AuthorizationKey())
	if session.AuthenticationData.AuthorizationCount+1 != req.AuthCount { // wrong authcount
		common.JsonResponseWithRespnoseType(ctx, response.InvalidAuthCountResponse{
			AuthorizationCount: session.AuthenticationData.AuthorizationCount,
		}, 1)
		return
	} else {
		session.AuthenticationData.AuthorizationCount++
	}

	// fmt.Println("User logins: ", userId)

	resp := session.Login()
	resp.SessionKey = session.EncodedSessionKey(req.Mask)
	{
		exist, _, startLiveRequest := user_live.LoadUserLive(session)
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
				user_live.ClearUserLive(session)
			}
		}
	}
	common.JsonResponse(ctx, &resp)

	{
		backupText, err := json.Marshal(resp)
		utils.CheckErr(err)
		utils.WriteAllText(fmt.Sprint(config.UserDataBackupPath, "login_", session.UserId, ".json"), string(backupText))
	}
}

func init() {
	router.AddHandler("/login/login", login)
}
