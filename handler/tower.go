package handler

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchTowerSelect(ctx *gin.Context) {
	// there's no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// no need to return anything, the client uses database for this
	// probably used to add DLP without having to add anything to database
	JsonResponse(ctx, &response.FetchTowerSelectResponse{
		UserModelDiff: &session.UserModel,
	})
}

func FetchTowerTop(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTowerTopRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
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

	session.Finalize("", "dummy")
	JsonResponse(ctx, &resp)
}

func ClearedTowerFloor(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ClearedTowerFloorRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userTower := session.GetUserTower(req.TowerId)
	if userTower.ClearedFloor < req.FloorNo {
		userTower.ClearedFloor = req.FloorNo
		session.UpdateUserTower(userTower)
	}
	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}

	session.Finalize("", "dummy")
	JsonResponse(ctx, &response.ClearedTowerFloorResponse{
		UserModelDiff: &session.UserModel,
	})
}

func RecoveryTowerCardUsed(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, cardMasterId := range req.CardMasterIds.Slice {
		cardUsedCount := session.GetUserTowerCardUsed(req.TowerId, cardMasterId)
		cardUsedCount.UsedCount--
		cardUsedCount.RecoveredCount++
		session.UpdateUserTowerCardUsed(req.TowerId, cardUsedCount)
	}
	// remove the item
	has := session.GetUserResourceByContent(item.PerformanceDrink).Content.ContentAmount
	cardCount := int32(req.CardMasterIds.Size())
	if has >= cardCount {
		session.RemoveResource(item.PerformanceDrink.Amount(cardCount))
	} else {
		session.RemoveResource(item.PerformanceDrink.Amount(has))
		session.RemoveSnsCoin((cardCount - has) * int32(session.Gamedata.Tower[req.TowerId].RecoverCostBySnsCoin))
	}

	session.Finalize("", "dummy")
	JsonResponse(ctx, &response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
	})
}

func RecoveryTowerCardUsedAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
	}
	for i := range resp.TowerCardUsedCountRows.Slice {
		resp.TowerCardUsedCountRows.Slice[i].UsedCount = 0
		resp.TowerCardUsedCountRows.Slice[i].RecoveredCount = 0
		session.UpdateUserTowerCardUsed(req.TowerId, resp.TowerCardUsedCountRows.Slice[i])
	}

	session.Finalize("", "dummy")
	JsonResponse(ctx, &resp)
}

func FetchTowerRanking(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTowerRankingRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// // TODO(multiplayer ranking): return actual data for this
	resp := response.FetchTowerRankingResponse{}
	resp.TopRankingCells.Append(session.GetTowerRankingCell(req.TowerId))
	resp.MyRankingCells.Append(session.GetTowerRankingCell(req.TowerId))
	resp.FriendRankingCells.Append(session.GetTowerRankingCell(req.TowerId))
	resp.RankingBorderInfo.Append(client.TowerRankingBorderInfo{
		RankingBorderVoltage: 0,
		RankingBorderMasterRow: client.TowerRankingBorderMasterRow{
			RankingType:  enum.EventCommonRankingTypeAll,
			UpperRank:    1,
			DisplayOrder: 1,
		}})
	resp.RankingBorderInfo.Append(client.TowerRankingBorderInfo{
		RankingBorderVoltage: 0,
		RankingBorderMasterRow: client.TowerRankingBorderMasterRow{
			RankingType:  enum.EventCommonRankingTypeFriend,
			UpperRank:    1,
			DisplayOrder: 1,
		}})
	resp.MyOrder = generic.NewNullable(int32(1))

	JsonResponse(ctx, &resp)
}
