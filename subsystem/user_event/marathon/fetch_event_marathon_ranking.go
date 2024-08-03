package marathon

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_social"
	"elichika/userdata"

	"sort"
)

func FetchEventMarathonRanking(session *userdata.Session, eventId int32) (*response.FetchEventMarathonRankingResponse, *response.RecoverableExceptionResponse) {
	event := session.Gamedata.EventActive.GetActiveEvent(session.Time)
	if (event == nil) || (event.EventId != eventId) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeEventMarathonOutOfDate,
		}
	}
	// constants like amount or border is from captured network data

	ranking := GetRanking(session.Db, eventId)
	resp := &response.FetchEventMarathonRankingResponse{}

	{
		records := ranking.GetRange(1, 100) // confirmed from network record
		for i, record := range records {
			if (i == 0) || (record.Score != records[i-1].Score) {
				resp.TopRankingCells.Append(client.EventMarathonRankingCell{
					Order:                    int32(i + 1),
					EventPoint:               record.Score,
					EventMarathonRankingUser: user_social.GetEventMarathonRankingUser(session, record.Id),
				})
			} else {
				resp.TopRankingCells.Append(client.EventMarathonRankingCell{
					Order:                    resp.TopRankingCells.Slice[i-1].Order,
					EventPoint:               record.Score,
					EventMarathonRankingUser: user_social.GetEventMarathonRankingUser(session, record.Id),
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
				resp.MyRankingCells.Append(client.EventMarathonRankingCell{
					Order:                    int32(i + low),
					EventPoint:               record.Score,
					EventMarathonRankingUser: user_social.GetEventMarathonRankingUser(session, record.Id),
				})
			} else {
				resp.MyRankingCells.Append(client.EventMarathonRankingCell{
					Order:                    resp.MyRankingCells.Slice[i-1].Order,
					EventPoint:               record.Score,
					EventMarathonRankingUser: user_social.GetEventMarathonRankingUser(session, record.Id),
				})
			}
		}
	}

	// friend ranking
	friendUserIds := user_social.GetFriendUserIds(session)
	for _, friendId := range friendUserIds {
		ep, exist := ranking.ScoreOf(friendId)
		if !exist {
			continue
		}
		resp.FriendRankingCells.Append(client.EventMarathonRankingCell{
			EventPoint:               ep,
			EventMarathonRankingUser: user_social.GetEventMarathonRankingUser(session, friendId),
		})
	}

	if hasRank {
		myEp, _ := ranking.ScoreOf(session.UserId)
		resp.FriendRankingCells.Append(client.EventMarathonRankingCell{
			EventPoint:               myEp,
			EventMarathonRankingUser: user_social.GetEventMarathonRankingUser(session, session.UserId),
		})
	}

	slice := &resp.FriendRankingCells.Slice
	sort.Slice(*slice, func(i, j int) bool {
		if (*slice)[i].EventPoint != (*slice)[j].EventPoint {
			return (*slice)[i].EventPoint > (*slice)[j].EventPoint
		}
		return (*slice)[i].EventMarathonRankingUser.UserId < (*slice)[j].EventMarathonRankingUser.UserId
	})
	for i := range *slice {
		if (i == 0) || ((*slice)[i].EventPoint != (*slice)[i-1].EventPoint) {
			(*slice)[i].Order = int32(i + 1)
		} else {
			(*slice)[i].Order = (*slice)[i-1].Order
		}
	}

	upperRanks := []int32{1, 1001, 3001, 10001, 30001, 50001, 70001, 90001, 100001, 0}
	for i, upperRank := range upperRanks {
		lowerRank := upperRanks[i+1] - 1
		if lowerRank != -1 {
			lastRecord := ranking.GetRange(int(lowerRank), int(lowerRank))
			rankingBorderPoint := int32(0)
			if len(lastRecord) > 0 {
				rankingBorderPoint = lastRecord[0].Score
			}
			resp.RankingBorderInfo.Append(client.EventMarathonRankingBorderInfo{
				RankingBorderPoint: rankingBorderPoint,
				RankingBorderMasterRow: client.EventMarathonRankingBorderMasterRow{
					RankingType:  enum.EventCommonRankingTypeAll,
					UpperRank:    upperRank,
					LowerRank:    generic.NewNullable(lowerRank),
					DisplayOrder: int32(i) + 1,
				},
			})
		} else {
			resp.RankingBorderInfo.Append(client.EventMarathonRankingBorderInfo{
				RankingBorderMasterRow: client.EventMarathonRankingBorderMasterRow{
					RankingType:  enum.EventCommonRankingTypeAll,
					UpperRank:    upperRank,
					DisplayOrder: int32(i) + 1,
				},
			})
		}

		if i == 0 {
			resp.RankingBorderInfo.Append(client.EventMarathonRankingBorderInfo{
				RankingBorderMasterRow: client.EventMarathonRankingBorderMasterRow{
					RankingType:  enum.EventCommonRankingTypeFriend,
					UpperRank:    1,
					DisplayOrder: 1,
				},
			})
		}

		if lowerRank == -1 {
			break
		}
	}
	return resp, nil
}
