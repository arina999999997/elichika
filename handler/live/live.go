package live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchLiveMusicSelect(ctx *gin.Context) {
	// ther is no request body

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()

	weekday := int32(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	resp := response.FetchLiveMusicSelectResponse{
		WeekdayState: client.WeekdayState{
			Weekday:       weekday,
			NextWeekdayAt: tomorrow,
		},
		UserModelDiff: &session.UserModel,
	}
	for _, liveDaily := range gamedata.LiveDaily {
		if liveDaily.Weekday != weekday {
			continue
		}
		resp.LiveDailyList.Append(client.LiveDaily{
			LiveDailyMasterId:      liveDaily.Id,
			LiveMasterId:           liveDaily.LiveId,
			EndAt:                  tomorrow,
			RemainingPlayCount:     5, // this is not kept track of
			RemainingRecoveryCount: generic.NewNullable(int32(10)),
		})
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func LiveStart(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.StartLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := session.Gamedata

	masterLiveDifficulty := gamedata.LiveDifficulty[req.LiveDifficultyId]
	masterLiveDifficulty.ConstructLiveStage(gamedata)
	session.UserStatus.LastLiveDifficultyId = req.LiveDifficultyId
	session.UserStatus.LatestLiveDeckId = req.DeckId

	// 保存请求包因为 /live/finish 接口的响应包里有部分字段不在该接口的请求包里
	// live is stored in db
	// TODO(refactor): Let's also store the request too
	resp := response.StartLiveResponse{
		Live: client.Live{
			// PartnerUserId:   req.PartnerUserId,
			LiveId:          time.Now().UnixNano(),
			LiveType:        enum.LiveTypeManual,
			IsPartnerFriend: true,
			DeckId:          req.DeckId,
			CellId:          req.CellId,
			// IsAutoplay:      req.IsAutoPlay,
		},
		UserModelDiff: &session.UserModel,
	}
	resp.Live.LiveStage = masterLiveDifficulty.LiveStage.Copy()
	// TODO(drop): fill in note here

	if req.LiveTowerStatus.HasValue {
		// is tower live, fetch this tower
		// TODO: fetch from database instead
		userTower := session.GetUserTower(req.LiveTowerStatus.Value.TowerId)
		reqTower := &req.LiveTowerStatus.Value
		if userTower.ReadFloor != reqTower.FloorNo {
			userTower.ReadFloor = reqTower.FloorNo
			session.UpdateUserTower(userTower)
		}
		resp.Live.TowerLive = generic.NewNullable(client.TowerLive{
			TowerId:       reqTower.TowerId,
			FloorNo:       reqTower.FloorNo,
			TargetVoltage: int32(gamedata.Tower[reqTower.TowerId].Floor[reqTower.FloorNo].TargetVoltage),
			StartVoltage:  userTower.Voltage,
		})
		resp.Live.LiveType = enum.LiveTypeTower
	}

	if req.IsAutoPlay {
		for k := range resp.Live.LiveStage.LiveNotes.Slice {
			resp.Live.LiveStage.LiveNotes.Slice[k].AutoJudgeType = *config.Conf.AutoJudgeType
		}
	}

	if req.PartnerUserId != 0 {
		resp.Live.LivePartnerCard = generic.NewNullable(user_card.GetOtherUserCard(session, req.PartnerUserId, req.PartnerCardMasterId))
	}

	session.SaveUserLive(resp.Live)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/fetchLiveMusicSelect", FetchLiveMusicSelect)
	router.AddHandler("/live/start", LiveStart)
	router.AddHandler("/live/finish", LiveFinish)
	router.AddHandler("/live/skip", LiveSkip)
	router.AddHandler("/live/updatePlayList", LiveUpdatePlayList)
	router.AddHandler("/live/finishTutorial", LiveFinish) // this works
}
