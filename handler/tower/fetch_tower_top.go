package tower

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_tower"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchTowerTop(ctx *gin.Context) {
	req := request.FetchTowerTopRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.FetchTowerTopResponse{
		TowerCardUsedCountRows: user_tower.GetUserTowerCardUsedList(session, req.TowerId),
		UserModelDiff:          &session.UserModel,
		// other fields are for DLP with voltage ranking
	}

	userTower := user_tower.GetUserTower(session, req.TowerId)
	tower := session.Gamedata.Tower[req.TowerId]
	if userTower.ClearedFloor == userTower.ReadFloor {
		tower := session.Gamedata.Tower[req.TowerId]
		if userTower.ReadFloor < tower.FloorCount {
			userTower.ReadFloor += 1
			resp.IsShowUnlockEffect = true
			// unlock all the bonus live at once
			for ; userTower.ReadFloor < tower.FloorCount; userTower.ReadFloor++ {
				if tower.Floor[userTower.ReadFloor].TowerCellType != enum.TowerCellTypeBonusLive {
					break
				}
			}
		}
	}
	user_tower.UpdateUserTower(session, userTower)

	// if tower with voltage ranking, then we have to prepare that
	if tower.IsVoltageRanked {
		// EachBonusLiveVoltage should be filled with zero for everything, then fill in the score

		resp.EachBonusLiveVoltage.Slice = make([]int32, tower.FloorCount)
		resp.Order = generic.NewNullable(int32(1))
		// fetch the score
		scores := user_tower.GetUserTowerVoltageRankingScores(session, req.TowerId)
		for _, score := range scores {
			resp.EachBonusLiveVoltage.Slice[score.FloorNo-1] = score.Voltage
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/tower/fetchTowerTop", fetchTowerTop)
}
