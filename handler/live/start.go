package live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
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

func start(ctx *gin.Context) {
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
	resp := response.StartLiveResponse{
		Live: client.Live{
			LiveId:          time.Now().UnixNano(),
			LiveType:        enum.LiveTypeManual,
			IsPartnerFriend: true,
			DeckId:          req.DeckId,
			CellId:          req.CellId,
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

	session.SaveUserLive(resp.Live, req)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/start", start)
}
