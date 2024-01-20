package live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchLiveMusicSelect(ctx *gin.Context) {
	// ther is no request body

	userId := ctx.GetInt("user_id")
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

	session.Finalize("{}", "dummy")
	handler.JsonResponse(ctx, &resp)
}

func FetchLivePartners(ctx *gin.Context) {
	// a set of partners player (i.e. friends and others), then fetch the card for them
	// this set include the current user, so we can use our own cards.
	// currently only have current user
	// note that all card are available, but we need to use the filter functionality in the client to actually get them to show up.

	resp := response.FetchLiveParntersResponse{}

	// there is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	partnerUserIds := []int32{} // TODO(friend): Fill this with some users
	partnerUserIds = append(partnerUserIds, int32(userId))

	for _, partnerId := range partnerUserIds {

		partner := client.LivePartner{}
		userdata.FetchDBProfile(partnerId, &partner)

		partner.IsFriend = true
		partnerCards := userdata.FetchPartnerCards(int(partnerId)) // client.UserCard
		if len(partnerCards) == 0 {
			continue
		}
		for _, card := range partnerCards {
			for i := 1; i <= 7; i++ {
				if (card.LivePartnerCategories & (1 << i)) != 0 {
					partner.CardByCategory.Set(int32(i), session.GetOtherUserCard(partnerId, card.CardMasterId))
				}
			}
		}
		resp.PartnerSelectState.LivePartners.Append(partner)
	}
	resp.PartnerSelectState.FriendCount = int32(resp.PartnerSelectState.LivePartners.Size())
	handler.JsonResponse(ctx, &resp)
}

func LiveStart(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.StartLiveRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
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
		resp.Live.LivePartnerCard = generic.NewNullable(session.GetOtherUserCard(req.PartnerUserId, req.PartnerCardMasterId))
	}

	session.SaveUserLive(resp.Live)

	session.Finalize("{}", "dummy")
	handler.JsonResponse(ctx, &resp)
}
