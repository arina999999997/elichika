package live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/klab"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_profile"
	"elichika/subsystem/user_status"
	"elichika/subsystem/voltage_ranking"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func handleLiveTypeManual(ctx *gin.Context, req request.FinishLiveRequest, session *userdata.Session, live client.Live, startReq request.StartLiveRequest) {
	gamedata := session.Gamedata
	liveDifficulty := gamedata.LiveDifficulty[session.UserStatus.LastLiveDifficultyId]

	resp := response.FinishLiveResponse{
		LiveResult: client.LiveResult{
			LiveDifficultyMasterId: session.UserStatus.LastLiveDifficultyId,
			LiveDeckId:             session.UserStatus.LatestLiveDeckId,
			Voltage:                req.LiveScore.CurrentScore,
			BeforeUserExp:          session.UserStatus.Exp,
			LiveFinishStatus:       req.LiveFinishStatus,
		},
		UserModelDiff: &session.UserModel,
	}

	isCenter := map[int32]bool{}

	for _, memberMapping := range liveDifficulty.Live.LiveMemberMapping {
		if memberMapping.IsCenter && (memberMapping.Position <= 9) {
			isCenter[int32(memberMapping.Position-1)] = true
		}
	}

	rewardCenterLovePoint := int32(0)
	if len(isCenter) != 0 {
		// liella songs have no center
		rewardCenterLovePoint = klab.CenterBondGainBasedOnBondGain(liveDifficulty.RewardBaseLovePoint) / int32(len(isCenter))
	}

	// record this live and build the last played deck
	userLiveDifficulty := session.GetUserLiveDifficulty(session.UserStatus.LastLiveDifficultyId)
	userLiveDifficulty.IsNew = false
	userLiveDifficulty.IsAutoplay = startReq.IsAutoPlay

	lastPlayDeck := client.LastPlayLiveDifficultyDeck{
		LiveDifficultyId: resp.LiveResult.LiveDifficultyMasterId,
		Voltage:          req.LiveScore.CurrentScore,
		IsCleared:        req.LiveFinishStatus == enum.LiveFinishStatusSucceeded,
		RecordedAt:       session.Time.Unix(),
	}

	userLiveDeck := session.GetUserLiveDeck(session.UserStatus.LatestLiveDeckId)
	for position := 1; position <= 9; position++ {
		cardMasterId := reflect.ValueOf(userLiveDeck).Field(1 + position).Interface().(generic.Nullable[int32]).Value
		suitMasterId := reflect.ValueOf(userLiveDeck).Field(1 + position + 9).Interface().(generic.Nullable[int32]).Value
		lastPlayDeck.CardWithSuitDict.Set(cardMasterId, suitMasterId)
	}
	liveParties := session.GetUserLivePartiesWithDeckId(session.UserStatus.LatestLiveDeckId)
	for _, liveParty := range liveParties {
		liveSquad := client.LiveSquad{}
		liveSquad.CardMasterIds.Append(liveParty.CardMasterId1.Value)
		liveSquad.CardMasterIds.Append(liveParty.CardMasterId2.Value)
		liveSquad.CardMasterIds.Append(liveParty.CardMasterId3.Value)
		liveSquad.UserAccessoryIds.Append(liveParty.UserAccessoryId1)
		liveSquad.UserAccessoryIds.Append(liveParty.UserAccessoryId2)
		liveSquad.UserAccessoryIds.Append(liveParty.UserAccessoryId3)
		lastPlayDeck.SquadDict.Set(liveParty.PartyId%10-1, liveSquad)
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

	if lastPlayDeck.IsCleared {
		// add story if it is a story mode
		if live.CellId.HasValue {
			session.InsertUserStoryMain(live.CellId.Value)
		}

		// update clear record
		userLiveDifficulty.ClearCount++
		if userLiveDifficulty.MaxScore < req.LiveScore.CurrentScore {
			userLiveDifficulty.MaxScore = req.LiveScore.CurrentScore
			// update voltage ranking
			userVoltageRanking := voltage_ranking.UserVoltageRanking{
				UserId:           session.UserId,
				LiveDifficultyId: userLiveDifficulty.LiveDifficultyId,
				VoltagePoint:     req.LiveScore.CurrentScore,
				DeckDetail: client.OtherUserDeckDetail{
					Deck: client.OtherUserDeck{
						Name: userLiveDeck.Name,
					},
				},
			}

			for _, liveParty := range liveParties {
				otherUserParty := client.OtherUserParty{
					Id: liveParty.PartyId,
				}
				otherUserParty.CardIds.Append(liveParty.CardMasterId1.Value)
				otherUserParty.CardIds.Append(liveParty.CardMasterId2.Value)
				otherUserParty.CardIds.Append(liveParty.CardMasterId3.Value)
				if liveParty.UserAccessoryId1.HasValue {
					otherUserParty.Accessories.Append(session.GetUserAccessory(liveParty.UserAccessoryId1.Value).ToOtherUserAccessory())
				}
				if liveParty.UserAccessoryId2.HasValue {
					otherUserParty.Accessories.Append(session.GetUserAccessory(liveParty.UserAccessoryId2.Value).ToOtherUserAccessory())
				}
				if liveParty.UserAccessoryId3.HasValue {
					otherUserParty.Accessories.Append(session.GetUserAccessory(liveParty.UserAccessoryId3.Value).ToOtherUserAccessory())
				}
				userVoltageRanking.DeckDetail.Deck.Parties.Append(otherUserParty)
			}
			for i := 1; i <= 9; i++ {
				cardMasterId := reflect.ValueOf(userLiveDeck).Field(1 + i).Interface().(generic.Nullable[int32]).Value
				suitMasterId := reflect.ValueOf(userLiveDeck).Field(1 + i + 9).Interface().(generic.Nullable[int32]).Value
				memberId := gamedata.Card[cardMasterId].Member.Id
				if !userVoltageRanking.DeckDetail.MemberLoveLevels.Has(memberId) {
					userVoltageRanking.DeckDetail.MemberLoveLevels.Set(memberId, session.GetMember(memberId).LoveLevel)
				}
				// no idea why the necessary stuff is like this, bad client code maybe
				otherUserCard := user_card.GetOtherUserCard(session, session.UserId, cardMasterId)
				for j := otherUserCard.AdditionalPassiveSkillIds.Size(); j < 9; j++ {
					otherUserCard.AdditionalPassiveSkillIds.Append(0) // this is the official server behaviour
				}
				otherUserCard.LoveLevel = *userVoltageRanking.DeckDetail.MemberLoveLevels.GetOnly(memberId)

				userVoltageRanking.DeckDetail.Deck.Cards.Append(otherUserCard)
				userVoltageRanking.DeckDetail.Deck.CardIds.Append(cardMasterId)
				userVoltageRanking.DeckDetail.Deck.SuitMasterIds.Append(suitMasterId)
			}

			voltage_ranking.SetVoltageRanking(session, userVoltageRanking)

		}
		if userLiveDifficulty.MaxCombo < req.LiveScore.HighestComboCount {
			userLiveDifficulty.MaxCombo = req.LiveScore.HighestComboCount
		}

		// award items
		for i, mission := range liveDifficulty.Missions {
			if (i == 0) || (int(req.LiveScore.CurrentScore) >= mission.TargetValue) {
				resp.LiveResult.LiveResultAchievements.Map[int32(i+1)].IsCurrentlyAchieved = true
				if !resp.LiveResult.LiveResultAchievements.Map[int32(i+1)].IsAlreadyAchieved { // new, add reward
					session.AddContent(mission.Reward)
					switch i {
					case 0:
						userLiveDifficulty.ClearedDifficultyAchievement1 = generic.NewNullable(int32(1))
					case 1:
						userLiveDifficulty.ClearedDifficultyAchievement2 = generic.NewNullable(int32(2))
					case 2:
						userLiveDifficulty.ClearedDifficultyAchievement3 = generic.NewNullable(int32(3))
					}
				}
			}
		}
		resp.LiveResult.GainUserExp = liveDifficulty.RewardUserExp
		user_status.AddUserExp(session, resp.LiveResult.GainUserExp)
	}

	resp.LiveResult.LastBestVoltage = userLiveDifficulty.MaxScore

	memberRepresentativeCard := make(map[int32]int32)
	memberLoveGained := make(map[int32]int32)
	for i := range req.LiveScore.CardStatDict.Map {
		liveFinishCard := req.LiveScore.CardStatDict.Map[i]

		// calculate mvp
		if liveFinishCard.GotVoltage > resp.LiveResult.Mvp.GetVoltage {
			resp.LiveResult.Mvp.GetVoltage = liveFinishCard.GotVoltage
			resp.LiveResult.Mvp.CardMasterId = liveFinishCard.CardMasterId
			resp.LiveResult.Mvp.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			resp.LiveResult.Mvp.AppealCount = liveFinishCard.AppealCount
		}

		// update card stat and member bond if cleared
		if lastPlayDeck.IsCleared {

			addedLove := liveDifficulty.RewardBaseLovePoint
			if isCenter[i] {
				addedLove += rewardCenterLovePoint
			}

			userCardPlayCountStat := session.GetUserCardPlayCountStat(liveFinishCard.CardMasterId)
			userCardPlayCountStat.LiveJoinCount++
			userCardPlayCountStat.ActiveSkillPlayCount += liveFinishCard.SkillTriggeredCount
			session.UpdateUserCardPlayCountStat(userCardPlayCountStat)
			// update member love point
			memberMasterId := gamedata.Card[liveFinishCard.CardMasterId].Member.Id

			_, exist := memberRepresentativeCard[memberMasterId]
			// only use 1 card master id or an idol might be shown multiple times
			if !exist {
				memberRepresentativeCard[memberMasterId] = liveFinishCard.CardMasterId
			}
			memberLoveGained[memberMasterId] += int32(addedLove)
		}
	}
	// it's normal to show +0 on the bond screen if the person is already maxed
	// this is checked against (video) recording
	for _, i := range req.LiveScore.CardStatDict.Order {
		liveFinishCard := req.LiveScore.CardStatDict.Map[i]
		memberMasterId := gamedata.Card[liveFinishCard.CardMasterId].Member.Id
		if memberRepresentativeCard[memberMasterId] != i {
			continue
		}
		addedLove := session.AddLovePoint(memberMasterId, memberLoveGained[memberMasterId])
		resp.LiveResult.MemberLoveStatuses.Set(liveFinishCard.CardMasterId, client.LiveResultMemberLoveStatus{
			RewardLovePoint: addedLove,
		})
	}

	resp.LiveResult.LiveResultAchievementStatus.ClearCount = userLiveDifficulty.ClearCount
	resp.LiveResult.LiveResultAchievementStatus.GotVoltage = req.LiveScore.CurrentScore
	resp.LiveResult.LiveResultAchievementStatus.RemainingStamina = req.LiveScore.RemainingStamina
	if live.LivePartnerCard.HasValue {
		resp.LiveResult.Partner = generic.NewNullable(user_profile.GetOtherUser(session, startReq.PartnerUserId))
	}
	session.UpdateLiveDifficulty(userLiveDifficulty)

	session.UpdateLastPlayLiveDifficultyDeck(lastPlayDeck)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func handleLiveTypeTower(ctx *gin.Context, req request.FinishLiveRequest, session *userdata.Session, live client.Live, startReq request.StartLiveRequest) {
	gamedata := session.Gamedata
	// liveDifficulty := gamedata.LiveDifficulty[session.UserStatus.LastLiveDifficultyId]

	// official server only record the Id, all other field are zero-valued
	userLiveDifficulty := session.GetUserLiveDifficulty(session.UserStatus.LastLiveDifficultyId)
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
		if liveFinishCard.GotVoltage > resp.LiveResult.Mvp.GetVoltage {
			resp.LiveResult.Mvp.GetVoltage = liveFinishCard.GotVoltage
			resp.LiveResult.Mvp.CardMasterId = liveFinishCard.CardMasterId
			resp.LiveResult.Mvp.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			resp.LiveResult.Mvp.AppealCount = liveFinishCard.AppealCount
		}
	}

	increasePlayCount := false
	awardFirstClearReward := false
	tower := gamedata.Tower[live.TowerLive.Value.TowerId]
	// manually quiting out shouldn't count as a clear
	if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded || req.LiveFinishStatus == enum.LiveFinishStatusFailure {
		userTower := session.GetUserTower(live.TowerLive.Value.TowerId)
		if tower.Floor[live.TowerLive.Value.FloorNo].TowerCellType == enum.TowerCellTypeBonusLive {
			// bonus live is only accepted when it's fully cleared
			if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded {
				// update the max score, while we can reuse user_live_difficulty, they seems to have zero values for the official server
				// so it's better to just use something else
				// that will also help with displaying the ranking
				currentScore := session.GetUserTowerVoltageRankingScore(live.TowerLive.Value.TowerId, live.TowerLive.Value.FloorNo)
				if (req.LiveScore.CurrentScore >= req.LiveScore.TargetScore) && (currentScore.Voltage < req.LiveScore.CurrentScore) {
					increasePlayCount = true
					awardFirstClearReward = currentScore.Voltage == 0
					currentScore.Voltage = req.LiveScore.CurrentScore
					session.UpdateUserTowerVoltageRankingScore(currentScore)
				}
			}
		} else if req.LiveScore.CurrentScore >= req.LiveScore.TargetScore { // first clear
			increasePlayCount = true
			awardFirstClearReward = true
			userTower.ClearedFloor = live.TowerLive.Value.FloorNo
			userTower.Voltage = 0
		} else { // not cleared
			increasePlayCount = true
			userTower.Voltage = int32(req.LiveScore.CurrentScore)
		}
		session.UpdateUserTower(userTower)
	}

	if increasePlayCount {
		// update card used stuff
		for i := range req.LiveScore.CardStatDict.Map {
			liveFinishCard := req.LiveScore.CardStatDict.Map[i]
			cardUsedCount := session.GetUserTowerCardUsed(live.TowerLive.Value.TowerId, liveFinishCard.CardMasterId)
			cardUsedCount.UsedCount++
			cardUsedCount.LastUsedAt = session.Time.Unix()
			session.UpdateUserTowerCardUsed(tower.TowerId, cardUsedCount)
			resp.LiveResult.LiveResultTower.Value.TowerCardUsedCounts.Append(cardUsedCount)
		}
	}
	if awardFirstClearReward {
		// TODO(present box): Reward are actually added to present box in official server, we just add them directly here
		if tower.Floor[live.TowerLive.Value.FloorNo].TowerClearRewardId != nil {
			session.AddTriggerBasic(
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopClearRewardReceived,
					ParamInt:        generic.NewNullable(live.TowerLive.Value.TowerId),
				})
			for _, reward := range tower.Floor[live.TowerLive.Value.FloorNo].TowerClearRewards {
				session.AddContent(reward)
			}
		}
		if tower.Floor[live.TowerLive.Value.FloorNo].TowerProgressRewardId != nil {
			session.AddTriggerBasic(
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopProgressRewardReceived,
					ParamInt:        generic.NewNullable(live.TowerLive.Value.TowerId),
				})
			for _, reward := range tower.Floor[live.TowerLive.Value.FloorNo].TowerProgressRewards {
				session.AddContent(reward)
			}
		}
	}

	session.UpdateLiveDifficulty(userLiveDifficulty)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func finish(ctx *gin.Context) {
	// this is pretty different for different type of live
	// for simplicity we just read the request and call different handlers, even though we might be able to save some extra work
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	exist, live, startReq := session.LoadUserLive()
	utils.MustExist(exist)
	session.ClearUserLive()
	// TODO(lp): Remove LP here if we want that

	switch live.LiveType {
	case enum.LiveTypeManual:
		handleLiveTypeManual(ctx, req, session, live, startReq)
	case enum.LiveTypeTower:
		handleLiveTypeTower(ctx, req, session, live, startReq)
	default:
		panic("not handled")
	}
}

func init() {
	router.AddHandler("/live/finish", finish)
	router.AddHandler("/live/finishTutorial", finish) // not necessary?
}
