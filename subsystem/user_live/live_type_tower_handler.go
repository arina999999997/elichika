package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_live_difficulty"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_tower"
	"elichika/subsystem/user_mission"
	"elichika/userdata"

	"fmt"
)

func liveTypeTowerHandler(session *userdata.Session, req request.FinishLiveRequest, live client.Live, startReq request.StartLiveRequest) response.FinishLiveResponse {
	gamedata := session.Gamedata
	// liveDifficulty := gamedata.LiveDifficulty[session.UserStatus.LastLiveDifficultyId]

	// official server only record the Id, all other field are zero-valued
	userLiveDifficulty := user_live_difficulty.GetUserLiveDifficulty(session, session.UserStatus.LastLiveDifficultyId)
	userLiveDifficulty.IsNew = false

	resp := response.FinishLiveResponse{
		LiveResult: client.LiveResult{
			LiveDifficultyMasterId: session.UserStatus.LastLiveDifficultyId,
			LiveDeckId:             session.UserStatus.LatestLiveDeckId,
			Voltage:                req.LiveScore.CurrentScore,
			LastBestVoltage:        userLiveDifficulty.MaxScore,
			BeforeUserExp:          session.UserStatus.Exp,
			LiveFinishStatus:       req.LiveFinishStatus,
			LiveResultTower: generic.NewNullable(client.LiveResultTower{
				TowerId:       live.TowerLive.Value.TowerId,
				FloorNo:       live.TowerLive.Value.FloorNo,
				TotalVoltage:  req.LiveScore.CurrentScore,
				GettedVoltage: req.LiveScore.CurrentScore - live.TowerLive.Value.StartVoltage,
			})},
		UserModelDiff: &session.UserModel,
	}

	for _, liveFinishCard := range req.LiveScore.CardStatDict.Map {
		// calculate mvp
		if liveFinishCard.GotVoltage > resp.LiveResult.Mvp.Value.GetVoltage {

			resp.LiveResult.Mvp = generic.NewNullable(client.LiveResultMvp{
				CardMasterId:        liveFinishCard.CardMasterId,
				GetVoltage:          liveFinishCard.GotVoltage,
				SkillTriggeredCount: liveFinishCard.SkillTriggeredCount,
				AppealCount:         liveFinishCard.AppealCount,
			})
		}
	}

	increasePlayCount := false
	awardFirstClearReward := false
	tower := gamedata.Tower[live.TowerLive.Value.TowerId]
	// manually quiting out shouldn't count as a clear
	if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded || req.LiveFinishStatus == enum.LiveFinishStatusFailure {
		userTower := user_tower.GetUserTower(session, live.TowerLive.Value.TowerId)
		if tower.Floor[live.TowerLive.Value.FloorNo].TowerCellType == enum.TowerCellTypeBonusLive {
			// bonus live is only accepted when it's fully cleared
			if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded {
				// update the max score, while we can reuse user_live_difficulty, they seems to have zero values for the official server
				// so it's better to just use something else
				// that will also help with displaying the ranking
				currentScore := user_tower.GetUserTowerVoltageRankingScore(session, live.TowerLive.Value.TowerId, live.TowerLive.Value.FloorNo)
				if (req.LiveScore.CurrentScore >= req.LiveScore.TargetScore) && (currentScore.Voltage < req.LiveScore.CurrentScore) {
					increasePlayCount = true
					awardFirstClearReward = currentScore.Voltage == 0
					currentScore.Voltage = req.LiveScore.CurrentScore
					user_tower.UpdateUserTowerVoltageRankingScore(session, currentScore)
				}
				// mission tracking
				user_mission.UpdateProgress(session, enum.MissionClearConditionTypeTowerClearLiveStage, nil, nil,
					user_mission.AddProgressHandler, int32(1))
			}
		} else if req.LiveScore.CurrentScore >= req.LiveScore.TargetScore { // first clear
			increasePlayCount = true
			awardFirstClearReward = true
			userTower.ClearedFloor = live.TowerLive.Value.FloorNo
			userTower.Voltage = 0
			// mission tracking
			user_mission.UpdateProgress(session, enum.MissionClearConditionTypeTowerClearLiveStage, nil, nil,
				user_mission.AddProgressHandler, int32(1))
		} else { // not cleared
			increasePlayCount = true
			userTower.Voltage = int32(req.LiveScore.CurrentScore)
		}
		user_tower.UpdateUserTower(session, userTower)
	}

	if increasePlayCount {
		// update card used stuff
		for i := range req.LiveScore.CardStatDict.Map {
			liveFinishCard := req.LiveScore.CardStatDict.Map[i]
			cardUsedCount := user_tower.GetUserTowerCardUsed(session, live.TowerLive.Value.TowerId, liveFinishCard.CardMasterId)
			cardUsedCount.UsedCount++
			cardUsedCount.LastUsedAt = session.Time.Unix()
			user_tower.UpdateUserTowerCardUsed(session, tower.TowerId, cardUsedCount)
			resp.LiveResult.LiveResultTower.Value.TowerCardUsedCounts.Append(cardUsedCount)
		}
	}
	if awardFirstClearReward {
		if tower.Floor[live.TowerLive.Value.FloorNo].TowerClearRewardId != nil {
			user_info_trigger.AddTriggerBasic(session,
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopClearRewardReceived,
					ParamInt:        generic.NewNullable(live.TowerLive.Value.TowerId),
				})
			for _, reward := range tower.Floor[live.TowerLive.Value.FloorNo].TowerClearRewards {
				user_present.AddPresentWithDuration(session, client.PresentItem{
					Content:          reward,
					PresentRouteType: enum.PresentRouteTypeTowerClearReward,
					PresentRouteId:   generic.NewNullable(tower.TowerId),
					ParamServer:      generic.NewNullable(tower.Title),
					ParamClient:      generic.NewNullable(fmt.Sprint(live.TowerLive.Value.FloorNo)),
				}, user_present.Duration30Days)
			}
		}
		if tower.Floor[live.TowerLive.Value.FloorNo].TowerProgressRewardId != nil {
			user_info_trigger.AddTriggerBasic(session,
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopProgressRewardReceived,
					ParamInt:        generic.NewNullable(live.TowerLive.Value.TowerId),
				})
			for _, reward := range tower.Floor[live.TowerLive.Value.FloorNo].TowerProgressRewards {
				user_present.AddPresentWithDuration(session, client.PresentItem{
					Content:          reward,
					PresentRouteType: enum.PresentRouteTypeTowerProgressReward,
					PresentRouteId:   generic.NewNullable(tower.TowerId),
					ParamServer:      generic.NewNullable(tower.Title),
					ParamClient:      generic.NewNullable(fmt.Sprint(live.TowerLive.Value.FloorNo)),
				}, user_present.Duration30Days)
			}
		}
	}

	user_live_difficulty.UpdateLiveDifficulty(session, userLiveDifficulty)
	return resp
}
