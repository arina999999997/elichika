package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_tower"
	"elichika/userdata"

	"time"
)

func StartLive(session *userdata.Session, req request.StartLiveRequest) response.StartLiveResponse {
	gamedata := session.Gamedata

	masterLiveDifficulty := gamedata.LiveDifficulty[req.LiveDifficultyId]
	masterLiveDifficulty.ConstructLiveStage(gamedata)
	session.UserStatus.LastLiveDifficultyId = req.LiveDifficultyId
	session.UserStatus.LatestLiveDeckId = req.DeckId

	// 保存请求包因为 /live/finish 接口的响应包里有部分字段不在该接口的请求包里
	// live is stored in db for /live/finish because some necessary info are not sent there by the client
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
		userTower := user_tower.GetUserTower(session, req.LiveTowerStatus.Value.TowerId)
		reqTower := &req.LiveTowerStatus.Value
		if userTower.ReadFloor != reqTower.FloorNo {
			userTower.ReadFloor = reqTower.FloorNo
			user_tower.UpdateUserTower(session, userTower)
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

	SaveUserLive(session, resp.Live, req)

	return resp
}
