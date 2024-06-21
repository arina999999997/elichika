package user_tower

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_social"
	"elichika/userdata"

	"sort"
)

// there's actually no error response for now, as tower are eternal and there's no reward payout
func FetchTowerRanking(session *userdata.Session, towerId int32) (*response.FetchTowerRankingResponse, *response.RecoverableExceptionResponse) {
	resp := response.FetchTowerRankingResponse{}
	ranking := GetRankingByTowerId(session, towerId)
	records := ranking.GetRange(1, 100)
	// top ranking has 100 people, not really sure what the official numbers are
	for i := range records {
		voltage := records[i].Score
		userId := records[i].Id
		order := int32(i + 1)
		if (i > 0) && (resp.TopRankingCells.Slice[i-1].SumVoltage == voltage) {
			order = resp.TopRankingCells.Slice[i-1].Order
		}
		resp.TopRankingCells.Append(client.TowerRankingCell{
			Order:            order,
			SumVoltage:       voltage,
			TowerRankingUser: GetTowerRankingUser(session, userId),
		})
	}

	// my ranking cell is from [myRank - 100, myRank + 100), but is rounded a bit
	myRank, hasRank := ranking.RankOf(session.UserId)
	if hasRank {
		resp.MyOrder = generic.NewNullable(int32(myRank))
		low := myRank - 100
		if low < 1 {
			low = 1
		}
		high := low + 200 - 1
		records = ranking.GetRange(low, high)
		for i := range records {
			voltage := records[i].Score
			userId := records[i].Id
			order := int32(i + low)
			if (i > 0) && (resp.TopRankingCells.Slice[i-1].SumVoltage == voltage) {
				order = resp.TopRankingCells.Slice[i-1].Order
			}
			resp.MyRankingCells.Append(client.TowerRankingCell{
				Order:            order,
				SumVoltage:       voltage,
				TowerRankingUser: GetTowerRankingUser(session, userId),
			})
		}
	}

	// friend ranking
	friendUserIds := user_social.GetFriendUserIds(session)
	for _, friendId := range friendUserIds {
		voltage, exist := ranking.ScoreOf(friendId)
		if !exist {
			continue
		}
		resp.FriendRankingCells.Append(client.TowerRankingCell{
			SumVoltage:       voltage,
			TowerRankingUser: GetTowerRankingUser(session, friendId),
		})
	}
	if hasRank {
		myVoltage, _ := ranking.ScoreOf(session.UserId)
		resp.FriendRankingCells.Append(client.TowerRankingCell{
			SumVoltage:       myVoltage,
			TowerRankingUser: GetTowerRankingUser(session, session.UserId),
		})
	}
	slice := &resp.FriendRankingCells.Slice
	sort.Slice(*slice, func(i, j int) bool {
		if (*slice)[i].SumVoltage != (*slice)[j].SumVoltage {
			return (*slice)[i].SumVoltage > (*slice)[j].SumVoltage
		}
		return (*slice)[i].TowerRankingUser.UserId < (*slice)[j].TowerRankingUser.UserId
	})
	for i := range *slice {
		if (i == 0) || ((*slice)[i].SumVoltage != (*slice)[i-1].SumVoltage) {
			(*slice)[i].Order = int32(i + 1)
		} else {
			(*slice)[i].Order = (*slice)[i-1].Order
		}
	}

	// TODO(ranking): the borders seems to be the reward tier
	// so even 1-1, 2-2, 3-3 and so show up
	// we'll just not bother with it for now
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
	return &resp, nil
}
