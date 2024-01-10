package live

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler"
	"elichika/klab"
	"elichika/model"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MemberLoveStatus struct {
	// TODO(refactor): remove this field
	CardMasterId    int32 `json:"-"`
	RewardLovePoint int32 `json:"reward_love_point"`
}

func (mls *MemberLoveStatus) Id() int64 {
	return int64(mls.CardMasterId)
}
func (mls *MemberLoveStatus) SetId(id int64) {
	mls.CardMasterId = int32(id)
}

type LiveResultAchievement struct {
	Position            int  `json:"position"`
	IsAlreadyAchieved   bool `json:"is_already_achieved"`
	IsCurrentlyAchieved bool `json:"is_currently_achieved"`
}

func (obj *LiveResultAchievement) Id() int64 {
	return int64(obj.Position)
}
func (obj *LiveResultAchievement) SetId(id int64) {
	obj.Position = int(id)
}

type LiveResultTower struct {
	TowerId             int32                          `json:"tower_id"`
	FloorNo             int32                          `json:"floor_no"`
	TotalVoltage        int32                          `json:"total_voltage"`
	GettedVoltage       int32                          `json:"getted_voltage"` // nice engrish
	TowerCardUsedCounts []model.UserTowerCardUsedCount `json:"tower_card_used_counts"`
}

type LiveFinishLiveResult struct {
	LiveDifficultyMasterId int                                            `json:"live_difficulty_master_id"`
	LiveDeckId             int                                            `json:"live_deck_id"`
	StandardDrops          []any                                          `json:"standard_drops"`
	AdditionalDrops        []any                                          `json:"additional_drops"`
	GimmickDrops           []any                                          `json:"gimmick_drops"`
	MemberLoveStatuses     generic.ObjectByObjectIdList[MemberLoveStatus] `json:"member_love_statuses"`
	MVP                    struct {
		CardMasterId        int32 `json:"card_master_id"`
		GetVoltage          int   `json:"get_voltage"`
		SkillTriggeredCount int   `json:"skill_triggered_count"`
		AppealCount         int   `json:"appeal_count"`
	} `json:"mvp"`
	Partner                     *model.UserBasicInfo                                `json:"partner"`
	LiveResultAchievements      generic.ObjectByObjectIdList[LiveResultAchievement] `json:"live_result_achievements"`
	LiveResultAchievementStatus struct {
		ClearCount       int `json:"clear_count"`
		GotVoltage       int `json:"got_voltage"`
		RemainingStamina int `json:"remaining_stamina"`
	} `json:"live_result_achievement_status"`
	Voltage                       int              `json:"voltage"`
	LastBestVoltage               int              `json:"last_best_voltage"`
	BeforeUserExp                 int              `json:"before_user_exp"`
	GainUserExp                   int              `json:"gain_user_exp"`
	IsRewardAccessoryInPresentBox bool             `json:"is_reward_accessory_in_present_box"`
	ActiveEventResult             *any             `json:"active_event_result"`
	LiveResultTower               *LiveResultTower `json:"live_result_tower"`
	LiveResultMemberGuild         *any             `json:"live_result_member_guild"`
	LiveFinishStatus              int              `json:"live_finish_status"`
}

type LiveDifficultyMission struct {
	Position    int
	TargetValue int
	Reward      client.Content `xorm:"extends"`
}

func handleLiveTypeManual(ctx *gin.Context, req request.LiveFinishRequest, session *userdata.Session, live model.UserLive) {
	liveDifficultyId := session.UserStatus.LastLiveDifficultyId
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	liveDifficulty := gamedata.LiveDifficulty[int(liveDifficultyId)]
	isCenter := map[int]bool{}
	for _, memberMapping := range liveDifficulty.Live.LiveMemberMapping {
		if memberMapping.IsCenter && (memberMapping.Position <= 9) {
			isCenter[memberMapping.Position-1] = true
		}
	}
	rewardCenterLovePoint := 0
	if len(isCenter) != 0 {
		// liella songs have no center
		rewardCenterLovePoint = klab.CenterBondGainBasedOnBondGain(liveDifficulty.RewardBaseLovePoint) / len(isCenter)
	}

	// record this live
	liveRecord := session.GetLiveDifficulty(session.UserStatus.LastLiveDifficultyId)
	liveRecord.IsNew = false
	lastPlayDeck := session.BuildLastPlayLiveDifficultyDeck(live.DeckId, int(liveDifficultyId))
	lastPlayDeck.Voltage = req.LiveScore.CurrentScore

	liveResult := LiveFinishLiveResult{
		LiveDifficultyMasterId: int(session.UserStatus.LastLiveDifficultyId),
		LiveDeckId:             int(session.UserStatus.LatestLiveDeckId),
		StandardDrops:          []any{},
		AdditionalDrops:        []any{},
		GimmickDrops:           []any{},
		Voltage:                req.LiveScore.CurrentScore,
		LastBestVoltage:        int(liveRecord.MaxScore),
		BeforeUserExp:          int(session.UserStatus.Exp),
		LiveFinishStatus:       req.LiveFinishStatus}

	liveRecord.PlayCount++
	lastPlayDeck.IsCleared = req.LiveFinishStatus == enum.LiveFinishStatusSucceeded
	liveResult.LiveResultAchievements.AppendNewWithId(1).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement1.HasValue
	liveResult.LiveResultAchievements.AppendNewWithId(2).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement2.HasValue
	liveResult.LiveResultAchievements.AppendNewWithId(3).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement3.HasValue
	if lastPlayDeck.IsCleared {
		// add story if it is a story mode
		if live.CellId != nil {
			session.InsertUserStoryMain(*live.CellId)
		}

		// update clear record
		liveRecord.ClearCount++
		if liveRecord.MaxScore < int32(req.LiveScore.CurrentScore) { // if new high score
			liveRecord.MaxScore = int32(req.LiveScore.CurrentScore)
		}
		if liveRecord.MaxCombo < int32(req.LiveScore.HighestComboCount) {
			liveRecord.MaxCombo = int32(req.LiveScore.HighestComboCount)
		}
		if liveDifficulty.IsCountTarget { // counted toward target and profiles
			liveStats := session.GetUserLiveStats()
			// TODO: just use the map instead of this
			idx := enum.LiveDifficultyIndex[liveDifficulty.LiveDifficultyType]
			liveStats.LivePlayCount[idx]++
			if liveRecord.ClearCount == 1 { // 1st clear
				liveStats.LiveClearCount[idx]++
			}
			session.UpdateUserLiveStats(liveStats)
		}

		// and award items
		for i, mission := range liveDifficulty.Missions {
			// TODO: the award condition is not checked totally correctly
			if (i == 0) || (req.LiveScore.CurrentScore >= mission.TargetValue) {
				liveResult.LiveResultAchievements.Objects[i].IsCurrentlyAchieved = true
				if !liveResult.LiveResultAchievements.Objects[i].IsAlreadyAchieved { // new, add reward
					session.AddResource(mission.Reward)
					switch i {
					case 0:
						liveRecord.ClearedDifficultyAchievement1 = generic.NewNullable(int32(1))
					case 1:
						liveRecord.ClearedDifficultyAchievement2 = generic.NewNullable(int32(2))
					case 2:
						liveRecord.ClearedDifficultyAchievement3 = generic.NewNullable(int32(3))
					}
				}
			}
		}
		liveResult.GainUserExp = liveDifficulty.RewardUserExp
	}

	memberPos := make(map[int32]int)
	loveAmount := [9]int32{}
	for i := range req.LiveScore.CardStatDict.Objects {
		liveFinishCard := req.LiveScore.CardStatDict.Objects[i]

		// calculate mvp
		if liveFinishCard.GotVoltage > liveResult.MVP.GetVoltage {
			liveResult.MVP.GetVoltage = liveFinishCard.GotVoltage
			liveResult.MVP.CardMasterId = liveFinishCard.CardMasterId
			liveResult.MVP.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			liveResult.MVP.AppealCount = liveFinishCard.AppealCount
		}

		// update card stat and member bond if cleared
		if lastPlayDeck.IsCleared {

			addedLove := liveDifficulty.RewardBaseLovePoint
			if isCenter[i] {
				addedLove += rewardCenterLovePoint
			}

			userCard := session.GetUserCard(liveFinishCard.CardMasterId)
			userCard.LiveJoinCount++
			userCard.ActiveSkillPlayCount += liveFinishCard.SkillTriggeredCount
			session.UpdateUserCard(userCard)
			// update member love point
			memberMasterId := gamedata.Card[liveFinishCard.CardMasterId].Member.Id

			_, exist := memberPos[memberMasterId]
			// only use 1 card master id or an idol might be shown multiple times
			if !exist {
				memberPos[memberMasterId] = i
			}
			loveAmount[memberPos[memberMasterId]] += int32(addedLove)
		}
	}
	// it's normal to show +0 on the bond screen if the person is already maxed
	// this is checked against (video) recording
	for i := range loveAmount {
		liveFinishCard := req.LiveScore.CardStatDict.Objects[i]
		memberMasterId := gamedata.Card[liveFinishCard.CardMasterId].Member.Id
		if memberPos[memberMasterId] != i {
			continue
		}
		addedLove := session.AddLovePoint(memberMasterId, loveAmount[i])
		liveResult.MemberLoveStatuses.PushBack(
			MemberLoveStatus{
				CardMasterId:    liveFinishCard.CardMasterId,
				RewardLovePoint: addedLove,
			})
	}

	liveResult.LiveResultAchievementStatus.ClearCount = int(liveRecord.ClearCount)
	liveResult.LiveResultAchievementStatus.GotVoltage = req.LiveScore.CurrentScore
	liveResult.LiveResultAchievementStatus.RemainingStamina = req.LiveScore.RemainingStamina
	if live.PartnerUserId != 0 {
		liveResult.Partner = new(model.UserBasicInfo)
		*liveResult.Partner = session.GetOtherUserBasicProfile(live.PartnerUserId)
	}
	session.UpdateLiveDifficulty(liveRecord)
	session.SetLastPlayLiveDifficultyDeck(lastPlayDeck)
	liveFinishResp := session.Finalize("{}", "user_model_diff")
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result", liveResult)

	resp := handler.SignResp(ctx, liveFinishResp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func handleLiveTypeTower(ctx *gin.Context, req request.LiveFinishRequest, session *userdata.Session, live model.UserLive) {
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	// liveDifficulty := gamedata.LiveDifficulty[session.UserStatus.LastLiveDifficultyId]

	// official server only record the Id, all other field are zero-valued
	liveRecord := session.GetLiveDifficulty(session.UserStatus.LastLiveDifficultyId)
	liveRecord.IsNew = false
	liveRecord.IsAutoplay = live.IsAutoplay
	liveResult := LiveFinishLiveResult{
		LiveDifficultyMasterId: int(session.UserStatus.LastLiveDifficultyId),
		LiveDeckId:             int(session.UserStatus.LatestLiveDeckId),
		StandardDrops:          []any{},
		AdditionalDrops:        []any{},
		GimmickDrops:           []any{},
		Voltage:                req.LiveScore.CurrentScore,
		LastBestVoltage:        int(liveRecord.MaxScore),
		BeforeUserExp:          int(session.UserStatus.Exp),
		LiveFinishStatus:       req.LiveFinishStatus,
		LiveResultTower: &LiveResultTower{
			TowerId:             *live.TowerLive.TowerId,
			FloorNo:             *live.TowerLive.FloorNo,
			TotalVoltage:        int32(req.LiveScore.CurrentScore),
			GettedVoltage:       int32(req.LiveScore.CurrentScore) - *live.TowerLive.StartVoltage,
			TowerCardUsedCounts: []model.UserTowerCardUsedCount{},
		}}

	for i := range req.LiveScore.CardStatDict.Objects {
		liveFinishCard := req.LiveScore.CardStatDict.Objects[i]
		// calculate mvp
		if liveFinishCard.GotVoltage > liveResult.MVP.GetVoltage {
			liveResult.MVP.GetVoltage = liveFinishCard.GotVoltage
			liveResult.MVP.CardMasterId = liveFinishCard.CardMasterId
			liveResult.MVP.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			liveResult.MVP.AppealCount = liveFinishCard.AppealCount
		}
	}

	increasePlayCount := false
	awardFirstClearReward := false
	tower := gamedata.Tower[*live.TowerLive.TowerId]
	// manually quiting out shouldn't count as a clear
	if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded || req.LiveFinishStatus == enum.LiveFinishStatusFailure {
		userTower := session.GetUserTower(*live.TowerLive.TowerId)
		if tower.Floor[*live.TowerLive.FloorNo].TowerCellType == enum.TowerCellTypeBonusLive {
			// bonus live is only accepted when it's fully cleared
			if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded {
				// update the max score, while we can reuse user_live_difficulty, they seems to have zero values for the official server
				// so it's better to just use something else
				// that will also help with displaying the ranking
				currentScore := session.GetUserTowerVoltageRankingScore(*live.TowerLive.TowerId, *live.TowerLive.FloorNo)
				if (req.LiveScore.CurrentScore >= req.LiveScore.TargetScore) && (int(currentScore.Voltage) < req.LiveScore.CurrentScore) {
					increasePlayCount = true
					awardFirstClearReward = currentScore.Voltage == 0
					currentScore.Voltage = int32(req.LiveScore.CurrentScore)
					session.UpdateUserTowerVoltageRankingScore(currentScore)
				}
			}
		} else if req.LiveScore.CurrentScore >= req.LiveScore.TargetScore { // first clear
			increasePlayCount = true
			awardFirstClearReward = true
			userTower.ClearedFloor = *live.TowerLive.FloorNo
			userTower.Voltage = 0
		} else { // not cleared
			increasePlayCount = true
			userTower.Voltage = int32(req.LiveScore.CurrentScore)
		}
		session.UpdateUserTower(userTower)
	}

	if increasePlayCount {
		// update card used stuff
		for i := range req.LiveScore.CardStatDict.Objects {
			liveFinishCard := req.LiveScore.CardStatDict.Objects[i]
			cardUsedCount := session.GetUserTowerCardUsed(*live.TowerLive.TowerId, int32(liveFinishCard.CardMasterId))
			cardUsedCount.UsedCount++
			cardUsedCount.LastUsedAt = session.Time.Unix()
			session.UpdateUserTowerCardUsed(cardUsedCount)
			liveResult.LiveResultTower.TowerCardUsedCounts = append(liveResult.LiveResultTower.TowerCardUsedCounts, cardUsedCount)
		}
	}
	if awardFirstClearReward {
		// TODO(present box): Reward are actually added to present box in official server, we just add them directly here
		if tower.Floor[*live.TowerLive.FloorNo].TowerClearRewardId != nil {
			session.AddTriggerBasic(
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopClearRewardReceived,
					ParamInt:        generic.NewNullable(*live.TowerLive.TowerId),
				})
			for _, reward := range tower.Floor[*live.TowerLive.FloorNo].TowerClearRewards {
				session.AddResource(reward)
			}
		}
		if tower.Floor[*live.TowerLive.FloorNo].TowerProgressRewardId != nil {
			session.AddTriggerBasic(
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopProgressRewardReceived,
					ParamInt:        generic.NewNullable(*live.TowerLive.TowerId),
				})
			for _, reward := range tower.Floor[*live.TowerLive.FloorNo].TowerProgressRewards {
				session.AddResource(reward)
			}
		}
	}

	session.UpdateLiveDifficulty(liveRecord)

	liveFinishResp := session.Finalize("{}", "user_model_diff")
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result", liveResult)
	resp := handler.SignResp(ctx, liveFinishResp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveFinish(ctx *gin.Context) {
	// this is pretty different for different type of live
	// for simplicity we just read the request and call different handlers, even though we might be able to save some extra work
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LiveFinishRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	exist, live := session.LoadUserLive()
	utils.MustExist(exist)
	switch live.LiveType {
	case enum.LiveTypeManual:
		handleLiveTypeManual(ctx, req, session, live)
	case enum.LiveTypeTower:
		handleLiveTypeTower(ctx, req, session, live)
	default:
		panic("not handled")
	}
}
