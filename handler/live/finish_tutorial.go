package live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
	"elichika/subsystem/user_live_difficulty"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func finishTutorial(ctx *gin.Context) {
	// this is pretty different for different type of live
	// for simplicity we just read the request and call different handlers, even though we might be able to save some extra work
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	exist, _, startReq := user_live.LoadUserLive(session)
	utils.MustExist(exist)
	user_live.ClearUserLive(session)

	userLiveDifficulty := user_live_difficulty.GetUserLiveDifficulty(session, session.UserStatus.LastLiveDifficultyId)
	userLiveDifficulty.IsNew = false
	userLiveDifficulty.IsAutoplay = startReq.IsAutoPlay

	resp := response.FinishLiveResponse{
		LiveResult: client.LiveResult{
			LiveDifficultyMasterId: session.UserStatus.LastLiveDifficultyId,
			LiveDeckId:             session.UserStatus.LatestLiveDeckId,
			Voltage:                req.LiveScore.CurrentScore,
			BeforeUserExp:          session.UserStatus.Exp,
			LiveFinishStatus:       req.LiveFinishStatus,
			LastBestVoltage:        userLiveDifficulty.MaxScore,
		},
		UserModelDiff: &session.UserModel,
	}

	userLiveDifficulty.PlayCount++
	resp.LiveResult.LiveResultAchievements.Set(1, client.LiveResultAchievement{
		Position:          1,
		IsAlreadyAchieved: userLiveDifficulty.ClearedDifficultyAchievement1.HasValue,
	})
	resp.LiveResult.LiveResultAchievements.Set(2, client.LiveResultAchievement{
		Position:          2,
		IsAlreadyAchieved: userLiveDifficulty.ClearedDifficultyAchievement2.HasValue,
	})
	resp.LiveResult.LiveResultAchievements.Set(3, client.LiveResultAchievement{
		Position:          3,
		IsAlreadyAchieved: userLiveDifficulty.ClearedDifficultyAchievement3.HasValue,
	})

	resp.LiveResult.LiveResultAchievementStatus.ClearCount = userLiveDifficulty.ClearCount
	resp.LiveResult.LiveResultAchievementStatus.GotVoltage = req.LiveScore.CurrentScore
	resp.LiveResult.LiveResultAchievementStatus.RemainingStamina = req.LiveScore.RemainingStamina

	user_live_difficulty.UpdateLiveDifficulty(session, userLiveDifficulty)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/finishTutorial", finishTutorial) // this is correct
}
