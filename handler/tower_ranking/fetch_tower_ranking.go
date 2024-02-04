package tower_ranking

import (
	"elichika/client"
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

func fetchTowerRanking(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTowerRankingRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// TODO(multiplayer ranking): return actual data for this
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

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/towerRanking/fetchTowerRanking", fetchTowerRanking)
}
