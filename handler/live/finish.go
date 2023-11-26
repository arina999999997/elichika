package live

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler"
	"elichika/klab"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type LiveFinishCard struct {
	CardMasterID        int `json:"-"`
	GotVoltage          int `json:"got_voltage"`
	SkillTriggeredCount int `json:"skill_triggered_count"`
	AppealCount         int `json:"appeal_count"`
}

func (obj *LiveFinishCard) SetID(id int64) {
	obj.CardMasterID = int(id)
}

type MemberLoveStatus struct {
	CardMasterID    int `json:"-"`
	RewardLovePoint int `json:"reward_love_point"`
}

func (obj *MemberLoveStatus) ID() int64 {
	return int64(obj.CardMasterID)
}
func (obj *MemberLoveStatus) SetID(id int64) {
	obj.CardMasterID = int(id)
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

type LiveFinishReq struct {
	LiveID           int64 `json:"live_id"`
	LiveFinishStatus int   `json:"live_finish_status"`
	LiveScore        struct {
		StartInfo                  any                                          `json:"start_info"`
		FinishInfo                 any                                          `json:"finish_info"`
		ResultDict                 []any                                        `json:"result_dict"`
		WaveStatDict               []any                                        `json:"wave_stat_dict"`
		TurnStatDict               []any                                        `json:"turn_stat_dict"`
		CardStatDict               generic.ObjectByObjectIDList[LiveFinishCard] `json:"card_stat_dict"`
		TargetScore                int                                          `json:"target_score"`
		CurrentScore               int                                          `json:"current_score"`
		ComboCount                 int                                          `json:"combo_count"`
		ChangeSquadCount           int                                          `json:"change_squad_count"`
		HighestComboCount          int                                          `json:"highest_combo_count"`
		RemainingStamina           int                                          `json:"remaining_stamina"`
		IsPerfectLive              bool                                         `json:"is_perfect_live"`
		IsPerfectFullCombo         bool                                         `json:"is_perfect_full_combo"`
		UseVoltageActiveSkillCount int                                          `json:"use_voltage_active_skill_count"`
		UseHealActiveSkillCount    int                                          `json:"use_heal_active_skill_count"`
		UseDebufActiveSkillCount   int                                          `json:"use_debuf_active_skill_count"`
		UseBufActiveSkillCount     int                                          `json:"use_buf_active_skill_count"`
		UseSpSkillCount            int                                          `json:"use_sp_skill_count"`
		CompleteAppealChanceCount  int                                          `json:"complete_appeal_chance_count"`
		TriggerCriticalCount       int                                          `json:"triggered_critical_count"`
		LivePower                  int                                          `json:"live_power"`
		SpSkillScoreList           []int                                        `json:"sp_skill_score_list"`
	} `json:"live_score"`
	ResumeFinishInfo any `json:"resume_finish_info"`
	RoomID           int `json:"room_id"`
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
	Voltage                       int  `json:"voltage"`
	LastBestVoltage               int  `json:"last_best_voltage"`
	BeforeUserExp                 int  `json:"before_user_exp"`
	GainUserExp                   int  `json:"gain_user_exp"`
	IsRewardAccessoryInPresentBox bool `json:"is_reward_accessory_in_present_box"`
	ActiveEventResult             *any `json:"active_event_result"`
	LiveResultTower               *any `json:"live_result_tower"`
	LiveResultMemberGuild         *any `json:"live_result_member_guild"`
	LiveFinishStatus              int  `json:"live_finish_status"`
}

type LiveDifficultyMission struct {
	Position    int
	TargetValue int
	Reward      model.Content `xorm:"extends"`
}

func LiveFinish(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := LiveFinishReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	exist, liveState := userdata.LoadLiveState(userID)
	utils.MustExist(exist)

	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	liveState.DeckID = session.UserStatus.LatestLiveDeckID
	liveState.LiveStage.LiveDifficultyID = session.UserStatus.LastLiveDifficultyID

	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	liveDifficulty := gamedata.LiveDifficulty[liveState.LiveStage.LiveDifficultyID]

	centerPositions := []int{}
	for _, memberMapping := range liveDifficulty.Live.LiveMemberMapping {
		if memberMapping.IsCenter {
			centerPositions = append(centerPositions, memberMapping.Position)
		}
	}

	rewardCenterLovePoint := klab.CenterBondGainBasedOnBondGain(liveDifficulty.RewardBaseLovePoint) / len(centerPositions)

	// record this live
	liveRecord := session.GetLiveDifficulty(session.UserStatus.LastLiveDifficultyID)
	liveRecord.IsNew = false
	lastPlayDeck := session.BuildLastPlayLiveDifficultyDeck(liveState.DeckID, liveState.LiveStage.LiveDifficultyID)
	lastPlayDeck.Voltage = req.LiveScore.CurrentScore

	liveResult := LiveFinishLiveResult{
		LiveDifficultyMasterID: session.UserStatus.LastLiveDifficultyID,
		LiveDeckID:             session.UserStatus.LatestLiveDeckID,
		StandardDrops:          []any{},
		AdditionalDrops:        []any{},
		GimmickDrops:           []any{},
		Voltage:                lastPlayDeck.Voltage,
		LastBestVoltage:        liveRecord.MaxScore,
		BeforeUserExp:          session.UserStatus.Exp,
		LiveFinishStatus:       req.LiveFinishStatus}

	liveRecord.PlayCount++
	lastPlayDeck.IsCleared = req.LiveFinishStatus == enum.LiveFinishStatusCleared
	liveResult.LiveResultAchievements.AppendNewWithID(1).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement1 != nil
	liveResult.LiveResultAchievements.AppendNewWithID(2).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement2 != nil
	liveResult.LiveResultAchievements.AppendNewWithID(3).IsAlreadyAchieved = liveRecord.ClearedDifficultyAchievement3 != nil
	if lastPlayDeck.IsCleared {
		// add story if it is a story mode
		if liveState.CellID != nil {
			session.InsertUserStoryMain(*liveState.CellID)
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

	bondCardPosition := make(map[int]int)
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
			isCenter := (i+1 == centerPositions[0])
			isCenter = isCenter || ((len(centerPositions) > 1) && (i+1 == centerPositions[1]))
			addedBond := liveDifficulty.RewardBaseLovePoint
			if isCenter {
				addedBond += rewardCenterLovePoint
			}

			userCard := session.GetUserCard(liveFinishCard.CardMasterID)
			userCard.LiveJoinCount++
			userCard.ActiveSkillPlayCount += liveFinishCard.SkillTriggeredCount
			session.UpdateUserCard(userCard)
			// update member love point
			memberMasterID := gamedata.Card[liveFinishCard.CardMasterID].Member.ID

			pos, exist := bondCardPosition[memberMasterID]
			// only use 1 card master id or an idol might be shown multiple times
			if !exist {
				memberLoveStatus := liveResult.MemberLoveStatuses.AppendNewWithID(int64(liveFinishCard.CardMasterID))
				memberLoveStatus.RewardLovePoint = addedBond
				bondCardPosition[memberMasterID] = liveResult.MemberLoveStatuses.Length - 1
			} else {
				liveResult.MemberLoveStatuses.Objects[pos].RewardLovePoint += addedBond
			}
		}
	}
	for memberMasterID, pos := range bondCardPosition {
		addedBond := session.AddLovePoint(memberMasterID, liveResult.MemberLoveStatuses.Objects[pos].RewardLovePoint)
		liveResult.MemberLoveStatuses.Objects[pos].RewardLovePoint = addedBond
	}

	liveResult.LiveResultAchievementStatus.ClearCount = liveRecord.ClearCount
	liveResult.LiveResultAchievementStatus.GotVoltage = req.LiveScore.CurrentScore
	liveResult.LiveResultAchievementStatus.RemainingStamina = req.LiveScore.RemainingStamina
	if liveState.PartnerUserID != 0 {
		liveResult.Partner = new(model.UserBasicInfo)
		*liveResult.Partner = session.GetOtherUserBasicProfile(liveState.PartnerUserID)
	}
	session.UpdateLiveDifficulty(liveRecord)
	session.SetLastPlayLiveDifficultyDeck(lastPlayDeck)
	liveFinishResp := session.Finalize("{}", "user_model_diff")
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result", liveResult)

	resp := handler.SignResp(ctx, liveFinishResp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
