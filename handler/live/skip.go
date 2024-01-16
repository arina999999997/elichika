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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DropWithTypes struct {
	StandardDrops   [0]model.RewardDrop `json:"standard_drops"`
	AdditionalDrops [0]model.RewardDrop `json:"additional_drops"`
	GimmickDrops    [0]model.RewardDrop `json:"gimmick_drops"`
}

type SkipLiveResult struct {
	LiveDifficultyMasterId        int                                            `json:"live_difficulty_master_id"`
	LiveDeckId                    int                                            `json:"live_deck_id"`
	Drops                         []DropWithTypes                                `json:"drops"`
	MemberLoveStatuses            generic.ObjectByObjectIdList[MemberLoveStatus] `json:"member_love_statuses"`
	GainUserExp                   int                                            `json:"gain_user_exp"`
	IsRewardAccessoryInPresentBox bool                                           `json:"is_reward_accessory_in_present_box"`
	ActiveEventResult             *int                                           `json:"active_event_result"`
	LiveResultMemberGuild         *int                                           `json:"live_result_member_guild"`
}

// TODO(refactor): Change to use request and response types
func LiveSkip(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type LiveSkipReq struct {
		LiveDifficultyMasterId int   `json:"live_difficulty_master_id"`
		DeckId                 int32 `json:"deck_id"`
		TicketUseCount         int   `json:"ticket_use_count"`
	}
	req := LiveSkipReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	session.UserStatus.LastLiveDifficultyId = int32(req.LiveDifficultyMasterId)
	liveDifficulty := gamedata.LiveDifficulty[req.LiveDifficultyMasterId]
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

	skipLiveResult := SkipLiveResult{
		LiveDifficultyMasterId: req.LiveDifficultyMasterId,
		LiveDeckId:             int(req.DeckId),
		GainUserExp:            liveDifficulty.RewardUserExp * req.TicketUseCount}

	for i := 1; i <= req.TicketUseCount; i++ {
		skipLiveResult.Drops = append(skipLiveResult.Drops, DropWithTypes{})
	}
	session.UserStatus.Exp += int32(skipLiveResult.GainUserExp)
	deck := session.GetUserLiveDeck(req.DeckId)
	deckJsonByte, _ := json.Marshal(deck)
	cardMasterIds := []int32{}
	gjson.Parse(string(deckJsonByte)).ForEach(func(key, value gjson.Result) bool {
		if strings.Contains(key.String(), "card_master_id") {
			cardMasterIds = append(cardMasterIds, int32(value.Int()))
		}
		return true
	})

	bondCardPosition := make(map[int32]int)
	for i, cardMasterId := range cardMasterIds {
		userCard := session.GetUserCard(cardMasterId)
		userCard.LiveJoinCount += req.TicketUseCount // count skip clear in pfp
		session.UpdateUserCard(userCard)
		// update member love point
		addedLove := liveDifficulty.RewardBaseLovePoint
		if isCenter[i] {
			addedLove += rewardCenterLovePoint
		}
		memberMasterId := gamedata.Card[cardMasterId].Member.Id

		pos, exist := bondCardPosition[memberMasterId]
		// only use 1 card master id or an idol might be shown multiple times
		if !exist {
			memberLoveStatus := skipLiveResult.MemberLoveStatuses.AppendNewWithId(int64(cardMasterId))
			memberLoveStatus.RewardLovePoint = int32(addedLove)
			bondCardPosition[memberMasterId] = skipLiveResult.MemberLoveStatuses.Length - 1
		} else {
			skipLiveResult.MemberLoveStatuses.Objects[pos].RewardLovePoint += int32(addedLove)
		}
	}
	for memberMasterId, pos := range bondCardPosition {
		addedLove := session.AddLovePoint(memberMasterId,
			int32(req.TicketUseCount)*skipLiveResult.MemberLoveStatuses.Objects[pos].RewardLovePoint)
		skipLiveResult.MemberLoveStatuses.Objects[pos].RewardLovePoint = addedLove
	}

	if liveDifficulty.IsCountTarget { // counted toward target and profiles
		liveStats := session.GetUserLiveStats()
		idx := enum.LiveDifficultyIndex[liveDifficulty.LiveDifficultyType]
		liveStats.LivePlayCount[idx] += req.TicketUseCount
		session.UpdateUserLiveStats(liveStats)
	}

	signBody := session.Finalize("{}", "user_model_diff")
	signBody, _ = sjson.Set(signBody, "skip_live_result", skipLiveResult)

	resp := handler.SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
