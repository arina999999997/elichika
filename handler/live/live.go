package live

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/handler"
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

func FetchLiveMusicSelect(ctx *gin.Context) {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	// does this work if it's EOY?
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	liveDailyList := []model.LiveDaily{}
	for _, liveDaily := range gamedata.LiveDaily {
		if liveDaily.Weekday != weekday {
			continue
		}
		liveDailyList = append(liveDailyList,
			model.LiveDaily{
				LiveDailyMasterId: liveDaily.Id,
				LiveMasterId:      liveDaily.LiveId,
			})
	}
	for k := range liveDailyList {
		liveDailyList[k].EndAt = int(tomorrow)
		liveDailyList[k].RemainingPlayCount = 5
		liveDailyList[k].RemainingRecoveryCount = 10
	}

	signBody := handler.GetData("fetchLiveMusicSelect.json")
	signBody, _ = sjson.Set(signBody, "weekday_state.weekday", weekday)
	signBody, _ = sjson.Set(signBody, "weekday_state.next_weekday_at", tomorrow)
	signBody, _ = sjson.Set(signBody, "live_daily_list", liveDailyList)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	signBody = session.Finalize(signBody, "user_model_diff")
	resp := handler.SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchLivePartners(ctx *gin.Context) {
	// a set of partners player (i.e. friends and others), then fetch the card for them
	// this set include the current user, so we can use our own cards.
	// currently only have current user
	// note that all card are available, but we need to use the filter functionality to actually get them to show up.
	partnerIds := []int{}
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	partnerIds = append(partnerIds, userId)
	livePartners := []model.LiveStartLivePartner{}
	for _, partnerId := range partnerIds {
		partner := model.LiveStartLivePartner{}
		partner.IsFriend = true
		userdata.FetchDBProfile(partnerId, &partner)
		partnerCards := userdata.FetchPartnerCards(partnerId) // client.UserCard
		if len(partnerCards) == 0 {
			continue
		}
		for _, card := range partnerCards {
			for i := 1; i <= 7; i++ {
				if (card.LivePartnerCategories & (1 << i)) != 0 {
					partnerCardInfo := session.GetPartnerCardFromUserCard(card)
					partner.CardByCategory = append(partner.CardByCategory, i)
					partner.CardByCategory = append(partner.CardByCategory, partnerCardInfo)
				}
			}
		}
		livePartners = append(livePartners, partner)
	}

	signBody := "{}"
	signBody, _ = sjson.Set(signBody, "partner_select_state.live_partners", livePartners)
	signBody, _ = sjson.Set(signBody, "partner_select_state.friend_count", len(livePartners))
	resp := handler.SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LiveStart(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LiveStartRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	masterLiveDifficulty := gamedata.LiveDifficulty[req.LiveDifficultyId]
	masterLiveDifficulty.ConstructLiveStage(gamedata)
	session.UserStatus.LastLiveDifficultyId = int32(req.LiveDifficultyId)
	session.UserStatus.LatestLiveDeckId = int32(req.DeckId)

	// 保存请求包因为 /live/finish 接口的响应包里有部分字段不在该接口的请求包里
	// live is stored in db
	live := model.UserLive{
		PartnerUserId:   req.PartnerUserId,
		LiveId:          time.Now().UnixNano(),
		LiveType:        enum.LiveTypeManual,
		IsPartnerFriend: true,
		DeckId:          req.DeckId,
		CellId:          req.CellId,
		IsAutoplay:      req.IsAutoPlay,
	}
	live.LiveStage = masterLiveDifficulty.LiveStage.Copy()

	if req.LiveTowerStatus != nil {
		// is tower live, fetch this tower
		// TODO: fetch from database instead
		userTower := session.GetUserTower(req.LiveTowerStatus.TowerId)

		if userTower.ReadFloor != req.LiveTowerStatus.FloorNo {
			userTower.ReadFloor = req.LiveTowerStatus.FloorNo
			session.UpdateUserTower(userTower)
		}
		dummy := int32(gamedata.Tower[req.LiveTowerStatus.TowerId].Floor[req.LiveTowerStatus.FloorNo].TargetVoltage)
		live.TowerLive = model.TowerLive{
			TowerId:       &req.LiveTowerStatus.TowerId,
			FloorNo:       &req.LiveTowerStatus.FloorNo,
			TargetVoltage: &dummy,
			StartVoltage:  &userTower.Voltage,
		}
		live.LiveType = enum.LiveTypeTower
	}

	if req.IsAutoPlay {
		for k := range live.LiveStage.LiveNotes {
			live.LiveStage.LiveNotes[k].AutoJudgeType = *config.Conf.AutoJudgeType
		}
	}

	if req.PartnerUserId != 0 {
		live.LivePartnerCard = session.GetPartnerCardFromUserCard(
			userdata.GetOtherUserCard(req.PartnerUserId, req.PartnerCardMasterId))
	}

	liveStartResp := session.Finalize("{}", "user_model_diff")
	liveStartResp, _ = sjson.Set(liveStartResp, "live", live)
	session.SaveUserLive(live)
	resp := handler.SignResp(ctx, liveStartResp, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
