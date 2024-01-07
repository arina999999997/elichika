package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/item"
	"elichika/protocol/request"
	"elichika/protocol/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	// "github.com/tidwall/sjson"
)

func FetchTowerSelect(ctx *gin.Context) {
	// there's no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// no need to return anything, the same use database for this
	respObj := response.FetchTowerSelectResponse{
		TowerIds:      []int{},
		UserModelDiff: &session.UserModel,
	}

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchTowerTop(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTowerTopRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	respObj := response.FetchTowerTopResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
		IsShowUnlockEffect:     false,
		// other fields are for DLP with voltage ranking
	}

	userTower := session.GetUserTower(req.TowerId)
	tower := gamedata.Tower[req.TowerId]
	if userTower.ClearedFloor == userTower.ReadFloor {
		tower := gamedata.Tower[req.TowerId]
		if userTower.ReadFloor < tower.FloorCount {
			userTower.ReadFloor += 1
			respObj.IsShowUnlockEffect = true
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
		//
		// EachBonusLiveVoltage should be filled with zero for everything, then fill in the score

		respObj.EachBonusLiveVoltage = make([]int, tower.FloorCount)
		respObj.Order = new(int)
		*respObj.Order = 1
		// fetch the score
		scores := session.GetUserTowerVoltageRankingScores(req.TowerId)
		for _, score := range scores {
			respObj.EachBonusLiveVoltage[score.FloorNo-1] = score.Voltage
		}
	}

	session.Finalize("", "dummy")

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ClearedTowerFloor(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ClearedTowerFloorRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	respObj := response.ClearedTowerFloorResponse{
		UserModelDiff:      &session.UserModel,
		IsShowUnlockEffect: false,
	}

	userTower := session.GetUserTower(req.TowerId)
	if userTower.ClearedFloor < req.FloorNo {
		userTower.ClearedFloor = req.FloorNo
		session.UpdateUserTower(userTower)
	}
	session.UserStatus.IsAutoMode = req.IsAutoMode
	session.Finalize("", "dummy")

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func RecoveryTowerCardUsed(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	tower := gamedata.Tower[req.TowerId]

	for _, cardMasterId := range req.CardMasterIds {
		cardUsedCount := session.GetUserTowerCardUsed(req.TowerId, cardMasterId)
		cardUsedCount.UsedCount--
		cardUsedCount.RecoveredCount++
		session.UpdateUserTowerCardUsed(cardUsedCount)
	}
	// remove the item
	has := session.GetUserResource(enum.ContentTypeRecoveryTowerCardUsedCount, 24001).Content.ContentAmount
	if has >= int32(len(req.CardMasterIds)) {
		session.RemoveResource(item.PerformanceDrink.Amount(int32(len(req.CardMasterIds))))
	} else {
		session.RemoveResource(item.PerformanceDrink.Amount(has))
		session.RemoveSnsCoin((int32(len(req.CardMasterIds)) - has) * int32(tower.RecoverCostBySnsCoin))

	}
	session.Finalize("", "dummy")
	respObj := response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
	}

	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func RecoveryTowerCardUsedAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoveryTowerCardUsedAllRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	respObj := response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
	}
	for i := range respObj.TowerCardUsedCountRows {
		respObj.TowerCardUsedCountRows[i].UsedCount = 0
		respObj.TowerCardUsedCountRows[i].RecoveredCount = 0
		session.UpdateUserTowerCardUsed(respObj.TowerCardUsedCountRows[i])
	}
	userTower := session.GetUserTower(req.TowerId)
	session.UpdateUserTower(userTower)

	session.Finalize("", "dummy")
	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func FetchTowerRanking(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTowerRankingRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// TODO(multiplayer ranking): return actual data for this
	respObj := response.FetchTowerRankingResponse{
		MyOrder: 1,
	}
	towerRankingCell := session.GetTowerRankingCell(req.TowerId)
	respObj.TopRankingCells = append(respObj.TopRankingCells, towerRankingCell)
	respObj.MyRankingCells = append(respObj.MyRankingCells, towerRankingCell)
	respObj.FriendRankingCells = append(respObj.FriendRankingCells, towerRankingCell)
	respObj.RankingBorderInfo = append(respObj.RankingBorderInfo,
		response.TowerRankingBorderInfo{
			RankingBorderVoltage: 0,
			RankingBorderMasterRow: response.TowerRankingBorderMasterRow{
				RankingType:  enum.EventCommonRankingTypeAll,
				UpperRank:    1,
				LowerRank:    1,
				DisplayOrder: 1,
			}})
	respObj.RankingBorderInfo = append(respObj.RankingBorderInfo,
		response.TowerRankingBorderInfo{
			RankingBorderVoltage: 0,
			RankingBorderMasterRow: response.TowerRankingBorderMasterRow{
				RankingType:  enum.EventCommonRankingTypeFriend,
				UpperRank:    1,
				LowerRank:    1,
				DisplayOrder: 1,
			}})
	respBytes, _ := json.Marshal(respObj)
	resp := SignResp(ctx, string(respBytes), config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
