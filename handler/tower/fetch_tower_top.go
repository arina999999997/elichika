package tower

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func fetchTowerTop(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTowerTopRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchTowerTopResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
		// other fields are for DLP with voltage ranking
	}

	userTower := session.GetUserTower(req.TowerId)
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
	session.UpdateUserTower(userTower)

	// if tower with voltage ranking, then we have to prepare that
	if tower.IsVoltageRanked {
		// EachBonusLiveVoltage should be filled with zero for everything, then fill in the score

		resp.EachBonusLiveVoltage.Slice = make([]int32, tower.FloorCount)
		resp.Order = generic.NewNullable(int32(1))
		// fetch the score
		scores := session.GetUserTowerVoltageRankingScores(req.TowerId)
		for _, score := range scores {
			resp.EachBonusLiveVoltage.Slice[score.FloorNo-1] = score.Voltage
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/tower/fetchTowerTop", fetchTowerTop)
}
