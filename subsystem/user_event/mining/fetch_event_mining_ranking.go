package mining

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_social"
	"elichika/userdata"

	"sort"
)

func FetchEventMiningRanking(session *userdata.Session, eventId int32) (*response.FetchEventMiningRankingResponse, *response.RecoverableExceptionResponse) {
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if (event == nil) || (event.EventId != eventId) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeEventMiningOutOfDate,
		}
	}
	// constants like amount or border is from captured network data
	friendUserIds := user_social.GetFriendUserIds(session)

	resp := &response.FetchEventMiningRankingResponse{}
	{ // points
		ranking := GetPointRanking(session.Db, eventId)
		{
			records := ranking.GetRange(1, 100) // confirmed from network record
			for i, record := range records {
				if (i == 0) || (record.Score != records[i-1].Score) {
					resp.PointTopRankingCells.Append(client.EventMiningRankingCell{
						Order:                  int32(i + 1),
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				} else {
					resp.PointTopRankingCells.Append(client.EventMiningRankingCell{
						Order:                  resp.PointTopRankingCells.Slice[i-1].Order,
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				}
			}
		}
		// return 81 record starting from own position - 39
		// so 39 before and 41 after
		// it's likely that there were an off by 1 error somewhere and they meant 40 40
		// or maybe they don't have proper read write lock
		myRank, hasRank := ranking.RankOf(session.UserId)
		if hasRank {
			low := myRank - 39
			if low < 1 {
				low = 1
			}
			high := low + 81 - 1
			records := ranking.GetRange(low, high)
			for i, record := range records {
				if (i == 0) || (record.Score != records[i-1].Score) {
					resp.PointMyRankingCells.Append(client.EventMiningRankingCell{
						Order:                  int32(i + low),
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				} else {
					resp.PointMyRankingCells.Append(client.EventMiningRankingCell{
						Order:                  resp.PointMyRankingCells.Slice[i-1].Order,
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				}
			}
		}

		// friend ranking
		for _, friendId := range friendUserIds {
			ep, exist := ranking.ScoreOf(friendId)
			if !exist {
				continue
			}
			resp.PointFriendRankingCells.Append(client.EventMiningRankingCell{
				EventPoint:             ep,
				EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, friendId),
			})
		}

		if hasRank {
			myEp, _ := ranking.ScoreOf(session.UserId)
			resp.PointFriendRankingCells.Append(client.EventMiningRankingCell{
				EventPoint:             myEp,
				EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, session.UserId),
			})
		}

		slice := &resp.PointFriendRankingCells.Slice
		sort.Slice(*slice, func(i, j int) bool {
			if (*slice)[i].EventPoint != (*slice)[j].EventPoint {
				return (*slice)[i].EventPoint > (*slice)[j].EventPoint
			}
			return (*slice)[i].EventMiningRankingUser.UserId < (*slice)[j].EventMiningRankingUser.UserId
		})
		for i := range *slice {
			if (i == 0) || ((*slice)[i].EventPoint != (*slice)[i-1].EventPoint) {
				(*slice)[i].Order = int32(i + 1)
			} else {
				(*slice)[i].Order = (*slice)[i-1].Order
			}
		}

		upperRanks := []int32{1, 501, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001, 10001, 20001, 30001, 40001, 50001, 60001, 70001, 80001, 90001, 100001, 0}
		for i, upperRank := range upperRanks {
			lowerRank := upperRanks[i+1] - 1
			if lowerRank != -1 {
				lastRecord := ranking.GetRange(int(lowerRank), int(lowerRank))
				rankingBorderPoint := int32(0)
				if len(lastRecord) > 0 {
					rankingBorderPoint = lastRecord[0].Score
				}
				resp.PointRankingBorderInfo.Append(client.EventMiningRankingBorderInfo{
					RankingBorderPoint: rankingBorderPoint,
					RankingBorderMasterRow: client.EventMiningRankingBorderMasterRow{
						RankingCategory: enum.EventRankingCategoryPoint,
						RankingType:     enum.EventCommonRankingTypeAll,
						UpperRank:       upperRank,
						LowerRank:       generic.NewNullable(lowerRank),
						DisplayOrder:    int32(i) + 1,
					},
				})
			} else {
				resp.PointRankingBorderInfo.Append(client.EventMiningRankingBorderInfo{
					RankingBorderMasterRow: client.EventMiningRankingBorderMasterRow{
						RankingCategory: enum.EventRankingCategoryPoint,
						RankingType:     enum.EventCommonRankingTypeAll,
						UpperRank:       upperRank,
						DisplayOrder:    int32(i) + 1,
					},
				})
			}

			if i == 0 {
				resp.PointRankingBorderInfo.Append(client.EventMiningRankingBorderInfo{
					RankingBorderMasterRow: client.EventMiningRankingBorderMasterRow{
						RankingCategory: enum.EventRankingCategoryPoint,
						RankingType:     enum.EventCommonRankingTypeFriend,
						UpperRank:       1,
						DisplayOrder:    1,
					},
				})
			}

			if lowerRank == -1 {
				break
			}
		}
	}

	{ // voltages
		ranking := GetVoltageRanking(session.Db, event.EventId)
		{
			records := ranking.GetRange(1, 100) // confirmed from network record
			for i, record := range records {
				if (i == 0) || (record.Score != records[i-1].Score) {
					resp.VoltageTopRankingCells.Append(client.EventMiningRankingCell{
						Order:                  int32(i + 1),
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				} else {
					resp.VoltageTopRankingCells.Append(client.EventMiningRankingCell{
						Order:                  resp.VoltageTopRankingCells.Slice[i-1].Order,
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				}
			}
		}
		// return 81 record starting from own position - 39
		// so 39 before and 41 after
		// it's likely that there were an off by 1 error somewhere and they meant 40 40
		// or maybe they don't have proper read write lock
		myRank, hasRank := ranking.RankOf(session.UserId)
		if hasRank {
			low := myRank - 39
			if low < 1 {
				low = 1
			}
			high := low + 81 - 1
			records := ranking.GetRange(low, high)
			for i, record := range records {
				if (i == 0) || (record.Score != records[i-1].Score) {
					resp.VoltageMyRankingCells.Append(client.EventMiningRankingCell{
						Order:                  int32(i + low),
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				} else {
					resp.VoltageMyRankingCells.Append(client.EventMiningRankingCell{
						Order:                  resp.VoltageMyRankingCells.Slice[i-1].Order,
						EventPoint:             record.Score,
						EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, record.Id),
					})
				}
			}
		}

		// friend ranking
		for _, friendId := range friendUserIds {
			ep, exist := ranking.ScoreOf(friendId)
			if !exist {
				continue
			}
			resp.VoltageFriendRankingCells.Append(client.EventMiningRankingCell{
				EventPoint:             ep,
				EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, friendId),
			})
		}

		if hasRank {
			myEp, _ := ranking.ScoreOf(session.UserId)
			resp.VoltageFriendRankingCells.Append(client.EventMiningRankingCell{
				EventPoint:             myEp,
				EventMiningRankingUser: user_social.GetEventMiningRankingUser(session, session.UserId),
			})
		}

		slice := &resp.VoltageFriendRankingCells.Slice
		sort.Slice(*slice, func(i, j int) bool {
			if (*slice)[i].EventPoint != (*slice)[j].EventPoint {
				return (*slice)[i].EventPoint > (*slice)[j].EventPoint
			}
			return (*slice)[i].EventMiningRankingUser.UserId < (*slice)[j].EventMiningRankingUser.UserId
		})
		for i := range *slice {
			if (i == 0) || ((*slice)[i].EventPoint != (*slice)[i-1].EventPoint) {
				(*slice)[i].Order = int32(i + 1)
			} else {
				(*slice)[i].Order = (*slice)[i-1].Order
			}
		}

		upperRanks := []int32{1, 51, 101, 301, 501, 1001, 2001, 3001, 4001, 5001, 6001, 7001, 8001, 9001, 10001, 0}
		for i, upperRank := range upperRanks {
			lowerRank := upperRanks[i+1] - 1
			if lowerRank != -1 {
				lastRecord := ranking.GetRange(int(lowerRank), int(lowerRank))
				rankingBorderPoint := int32(0)
				if len(lastRecord) > 0 {
					rankingBorderPoint = lastRecord[0].Score
				}
				resp.VoltageRankingBorderInfo.Append(client.EventMiningRankingBorderInfo{
					RankingBorderPoint: rankingBorderPoint,
					RankingBorderMasterRow: client.EventMiningRankingBorderMasterRow{
						RankingCategory: enum.EventRankingCategoryVoltage,
						RankingType:     enum.EventCommonRankingTypeAll,
						UpperRank:       upperRank,
						LowerRank:       generic.NewNullable(lowerRank),
						DisplayOrder:    int32(i) + 1,
					},
				})
			} else {
				resp.VoltageRankingBorderInfo.Append(client.EventMiningRankingBorderInfo{
					RankingBorderMasterRow: client.EventMiningRankingBorderMasterRow{
						RankingCategory: enum.EventRankingCategoryVoltage,
						RankingType:     enum.EventCommonRankingTypeAll,
						UpperRank:       upperRank,
						DisplayOrder:    int32(i) + 1,
					},
				})
			}

			if i == 0 {
				resp.VoltageRankingBorderInfo.Append(client.EventMiningRankingBorderInfo{
					RankingBorderMasterRow: client.EventMiningRankingBorderMasterRow{
						RankingCategory: enum.EventRankingCategoryVoltage,
						RankingType:     enum.EventCommonRankingTypeFriend,
						UpperRank:       1,
						DisplayOrder:    1,
					},
				})
			}

			if lowerRank == -1 {
				break
			}
		}
	}

	return resp, nil
}
