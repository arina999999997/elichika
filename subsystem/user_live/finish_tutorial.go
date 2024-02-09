package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/subsystem/user_live_difficulty"
	"elichika/userdata"
	"elichika/utils"
)

func FinishTutorial(session *userdata.Session, req request.FinishLiveRequest) response.FinishLiveResponse {
	exist, _, startReq := LoadUserLive(session)
	utils.MustExist(exist)
	ClearUserLive(session)

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

	return resp
}
