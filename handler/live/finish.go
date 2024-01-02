package live

import (
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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MemberLoveStatus struct {
	CardMasterID    int `json:"-"`
	RewardLovePoint int `json:"reward_love_point"`
}

func (mls *MemberLoveStatus) ID() int64 {
	return int64(mls.CardMasterID)
}
func (mls *MemberLoveStatus) SetID(id int64) {
	mls.CardMasterID = int(id)
}

type LiveResultAchievement struct {
	Position            int  `json:"position"`
	IsAlreadyAchieved   bool `json:"is_already_achieved"`
	IsCurrentlyAchieved bool `json:"is_currently_achieved"`
}

func (obj *LiveResultAchievement) ID() int64 {
	return int64(obj.Position)
}
func (obj *LiveResultAchievement) SetID(id int64) {
	obj.Position = int(id)
}

type LiveResultTower struct {
	TowerID             int                            `json:"tower_id"`
	FloorNo             int                            `json:"floor_no"`
	TotalVoltage        int                            `json:"total_voltage"`
	GettedVoltage       int                            `json:"getted_voltage"` // nice engrish
	TowerCardUsedCounts []model.UserTowerCardUsedCount `json:"tower_card_used_counts"`
}

type LiveFinishLiveResult struct {
	LiveDifficultyMasterID int                                            `json:"live_difficulty_master_id"`
	LiveDeckID             int                                            `json:"live_deck_id"`
	StandardDrops          []any                                          `json:"standard_drops"`
	AdditionalDrops        []any                                          `json:"additional_drops"`
	GimmickDrops           []any                                          `json:"gimmick_drops"`
	MemberLoveStatuses     generic.ObjectByObjectIDList[MemberLoveStatus] `json:"member_love_statuses"`
	MVP                    struct {
		CardMasterID        int `json:"card_master_id"`
		GetVoltage          int `json:"get_voltage"`
		SkillTriggeredCount int `json:"skill_triggered_count"`
		AppealCount         int `json:"appeal_count"`
	} `json:"mvp"`
	Partner                     *model.UserBasicInfo                                `json:"partner"`
	LiveResultAchievements      generic.ObjectByObjectIDList[LiveResultAchievement] `json:"live_result_achievements"`
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
	Reward      model.Content `xorm:"extends"`
}

func handleLiveTypeManual(ctx *gin.Context, req request.LiveFinishRequest, session *userdata.Session, live model.UserLive) {
	liveDifficultyID := session.UserStatus.LastLiveDifficultyID
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	liveDifficulty := gamedata.LiveDifficulty[liveDifficultyID]
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
	liveRecord := session.GetLiveDifficulty(session.UserStatus.LastLiveDifficultyID)
	liveRecord.IsNew = false
	lastPlayDeck := session.BuildLastPlayLiveDifficultyDeck(live.DeckID, liveDifficultyID)
	lastPlayDeck.Voltage = req.LiveScore.CurrentScore

	liveResult := LiveFinishLiveResult{
		LiveDifficultyMasterID: session.UserStatus.LastLiveDifficultyID,
		LiveDeckID:             session.UserStatus.LatestLiveDeckID,
		StandardDrops:          []any{},
		AdditionalDrops:        []any{},
		GimmickDrops:           []any{},
		Voltage:                req.LiveScore.CurrentScore,
		LastBestVoltage:        liveRecord.MaxScore,
		BeforeUserExp:          session.UserStatus.Exp,
		LiveFinishStatus:       req.LiveFinishStatus}

	liveRecord.PlayCount++
	lastPlayDeck.IsCleared = req.LiveFinishStatus == enum.LiveFinishStatusSucceeded
	liveResult.LiveResultAchievements.AppendNewWithID(1).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement1 != nil
	liveResult.LiveResultAchievements.AppendNewWithID(2).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement2 != nil
	liveResult.LiveResultAchievements.AppendNewWithID(3).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement3 != nil
	if lastPlayDeck.IsCleared {
		// add story if it is a story mode
		if live.CellID != nil {
			session.InsertUserStoryMain(*live.CellID)
		}

		// update clear record
		liveRecord.ClearCount++
		if liveRecord.MaxScore < req.LiveScore.CurrentScore { // if new high score
			liveRecord.MaxScore = req.LiveScore.CurrentScore
		}
		if liveRecord.MaxCombo < req.LiveScore.HighestComboCount {
			liveRecord.MaxCombo = req.LiveScore.HighestComboCount
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
						liveRecord.ClearedDifficultyAchievement1 = new(int)
						*liveRecord.ClearedDifficultyAchievement1 = 1
					case 1:
						liveRecord.ClearedDifficultyAchievement2 = new(int)
						*liveRecord.ClearedDifficultyAchievement2 = 2
					case 2:
						liveRecord.ClearedDifficultyAchievement3 = new(int)
						*liveRecord.ClearedDifficultyAchievement3 = 3
					}
				}
			}
		}
		liveResult.GainUserExp = liveDifficulty.RewardUserExp
	}

	memberPos := make(map[int]int)
	loveAmount := [9]int{}
	for i := range req.LiveScore.CardStatDict.Objects {
		liveFinishCard := req.LiveScore.CardStatDict.Objects[i]

		// calculate mvp
		if liveFinishCard.GotVoltage > liveResult.MVP.GetVoltage {
			liveResult.MVP.GetVoltage = liveFinishCard.GotVoltage
			liveResult.MVP.CardMasterID = liveFinishCard.CardMasterID
			liveResult.MVP.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			liveResult.MVP.AppealCount = liveFinishCard.AppealCount
		}

		// update card stat and member bond if cleared
		if lastPlayDeck.IsCleared {

			addedLove := liveDifficulty.RewardBaseLovePoint
			if isCenter[i] {
				addedLove += rewardCenterLovePoint
			}

			userCard := session.GetUserCard(liveFinishCard.CardMasterID)
			userCard.LiveJoinCount++
			userCard.ActiveSkillPlayCount += liveFinishCard.SkillTriggeredCount
			session.UpdateUserCard(userCard)
			// update member love point
			memberMasterID := gamedata.Card[liveFinishCard.CardMasterID].Member.ID

			_, exist := memberPos[memberMasterID]
			// only use 1 card master id or an idol might be shown multiple times
			if !exist {
				memberPos[memberMasterID] = i
			}
			loveAmount[memberPos[memberMasterID]] += addedLove
		}
	}
	// it's normal to show +0 on the bond screen if the person is already maxed
	// this is checked against (video) recording
	for i := range loveAmount {
		liveFinishCard := req.LiveScore.CardStatDict.Objects[i]
		memberMasterID := gamedata.Card[liveFinishCard.CardMasterID].Member.ID
		if memberPos[memberMasterID] != i {
			continue
		}
		addedLove := session.AddLovePoint(memberMasterID, loveAmount[i])
		liveResult.MemberLoveStatuses.PushBack(
			MemberLoveStatus{
				CardMasterID:    liveFinishCard.CardMasterID,
				RewardLovePoint: addedLove,
			})
	}

	liveResult.LiveResultAchievementStatus.ClearCount = liveRecord.ClearCount
	liveResult.LiveResultAchievementStatus.GotVoltage = req.LiveScore.CurrentScore
	liveResult.LiveResultAchievementStatus.RemainingStamina = req.LiveScore.RemainingStamina
	if live.PartnerUserID != 0 {
		liveResult.Partner = new(model.UserBasicInfo)
		*liveResult.Partner = session.GetOtherUserBasicProfile(live.PartnerUserID)
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
	// liveDifficulty := gamedata.LiveDifficulty[session.UserStatus.LastLiveDifficultyID]

	// official server only record the ID, all other field are zero-valued
	liveRecord := session.GetLiveDifficulty(session.UserStatus.LastLiveDifficultyID)
	liveRecord.IsNew = false
	liveRecord.IsAutoplay = live.IsAutoplay
	timeStamp := time.Now().Unix()
	liveResult := LiveFinishLiveResult{
		LiveDifficultyMasterID: session.UserStatus.LastLiveDifficultyID,
		LiveDeckID:             session.UserStatus.LatestLiveDeckID,
		StandardDrops:          []any{},
		AdditionalDrops:        []any{},
		GimmickDrops:           []any{},
		Voltage:                req.LiveScore.CurrentScore,
		LastBestVoltage:        liveRecord.MaxScore,
		BeforeUserExp:          session.UserStatus.Exp,
		LiveFinishStatus:       req.LiveFinishStatus,
		LiveResultTower: &LiveResultTower{
			TowerID:             *live.TowerLive.TowerID,
			FloorNo:             *live.TowerLive.FloorNo,
			TotalVoltage:        req.LiveScore.CurrentScore,
			GettedVoltage:       req.LiveScore.CurrentScore - *live.TowerLive.StartVoltage,
			TowerCardUsedCounts: []model.UserTowerCardUsedCount{},
		}}

	for i := range req.LiveScore.CardStatDict.Objects {
		liveFinishCard := req.LiveScore.CardStatDict.Objects[i]
		// calculate mvp
		if liveFinishCard.GotVoltage > liveResult.MVP.GetVoltage {
			liveResult.MVP.GetVoltage = liveFinishCard.GotVoltage
			liveResult.MVP.CardMasterID = liveFinishCard.CardMasterID
			liveResult.MVP.SkillTriggeredCount = liveFinishCard.SkillTriggeredCount
			liveResult.MVP.AppealCount = liveFinishCard.AppealCount
		}
	}

	increasePlayCount := false
	awardFirstClearReward := false
	tower := gamedata.Tower[*live.TowerLive.TowerID]
	// manually quiting out shouldn't count as a clear
	if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded || req.LiveFinishStatus == enum.LiveFinishStatusFailure {
		userTower := session.GetUserTower(*live.TowerLive.TowerID)
		if tower.Floor[*live.TowerLive.FloorNo].TowerCellType == enum.TowerCellTypeBonusLive {
			// bonus live is only accepted when it's fully cleared
			if req.LiveFinishStatus == enum.LiveFinishStatusSucceeded {
				// update the max score, while we can reuse user_live_difficulty, they seems to have zero values for the official server
				// so it's better to just use something else
				// that will also help with displaying the ranking
				currentScore := session.GetUserTowerVoltageRankingScore(*live.TowerLive.TowerID, *live.TowerLive.FloorNo)
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
			userTower.ClearedFloor = *live.TowerLive.FloorNo
			userTower.Voltage = 0
		} else { // not cleared
			increasePlayCount = true
			userTower.Voltage = req.LiveScore.CurrentScore
		}
		session.UpdateUserTower(userTower)
	}

	if increasePlayCount {
		// update card used stuff
		for i := range req.LiveScore.CardStatDict.Objects {
			liveFinishCard := req.LiveScore.CardStatDict.Objects[i]
			cardUsedCount := session.GetUserTowerCardUsed(*live.TowerLive.TowerID, liveFinishCard.CardMasterID)
			cardUsedCount.UsedCount++
			cardUsedCount.LastUsedAt = timeStamp
			session.UpdateUserTowerCardUsed(cardUsedCount)
			liveResult.LiveResultTower.TowerCardUsedCounts = append(liveResult.LiveResultTower.TowerCardUsedCounts, cardUsedCount)
		}
	}
	if awardFirstClearReward {
		// TODO(present box): Reward are actually added to present box in official server, we just add them directly here
		if tower.Floor[*live.TowerLive.FloorNo].TowerClearRewardID != nil {
			session.AddTriggerBasic(
				model.TriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopClearRewardReceived,
					ParamInt:        *live.TowerLive.TowerID,
				})
			for _, reward := range tower.Floor[*live.TowerLive.FloorNo].TowerClearRewards {
				session.AddResource(reward)
			}
		}
		if tower.Floor[*live.TowerLive.FloorNo].TowerProgressRewardID != nil {
			session.AddTriggerBasic(
				model.TriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeTowerTopProgressRewardReceived,
					ParamInt:        *live.TowerLive.TowerID,
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
	userID := ctx.GetInt("user_id")
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LiveFinishRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	session := userdata.GetSession(ctx, userID)
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
