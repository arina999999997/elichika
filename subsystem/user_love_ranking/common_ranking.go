package user_love_ranking

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_social"
	"elichika/userdata"
	"elichika/utils"

	"fmt"
)

type basicLoveRankingData struct {
	UserId    int32 `xorm:"pk 'user_id'"`
	LovePoint int32 `xorm:"'love_point'"`
	Order     int32 `xorm:"-"`
}

func commonRanking(session *userdata.Session, condition int32, rankingOrder generic.Nullable[int32]) response.FetchLoveRankingResponse {
	resp := response.FetchLoveRankingResponse{}
	if !rankingOrder.HasValue {
		rankingOrder.Value = findCommonRankingOrder(session, condition)
	}
	resp.MyRankingOrder.Value = findCommonRankingOrder(session, condition)
	if resp.MyRankingOrder.Value != 0 {
		resp.MyRankingOrder.HasValue = true
	}
	rankingOrderFirst := rankingOrder.Value - 124
	if rankingOrderFirst < 1 {
		rankingOrderFirst = 1
	}
	if (condition >= enum.LoveRankingConditionTypeMember1) && (condition <= enum.LoveRankingConditionTypeMember212) {
		// single condition
		memberId := getMemberIdFromCondition(condition)
		results := []basicLoveRankingData{}
		err := session.Db.Table("u_member").Where("member_master_id = ? AND love_point > 0", memberId).OrderBy("love_point DESC").
			Limit(250, int(rankingOrderFirst)).Find(&results)
		utils.CheckErr(err)
		for i := range results {
			if i == 0 {
				results[i].Order = findCommonRankingOrderByPoint(session, condition, results[i].LovePoint)
			} else if results[i].LovePoint == results[i-1].LovePoint {
				results[i].Order = results[i-1].Order
			} else {
				results[i].Order = results[0].Order + int32(i)
			}
			resp.LoveRankingData.Append(client.LoveRankingData{
				RankingUser: user_social.GetRankingUser(session, results[i].UserId),
				LovePoint:   results[i].LovePoint,
				Order:       results[i].Order,
			})
		}
	} else {
		userIds := []int32{}
		lovePoints := []int32{}
		orders := []int32{}
		// xorm is stupid, these can't be combined unless we create an interface that has the relevant key
		err := session.Db.Table("u_total_love_point_summary").
			Where(fmt.Sprintf("condition_%d > 0", condition)).
			OrderBy(fmt.Sprintf("condition_%d DESC, user_id", condition)).
			Limit(250, int(rankingOrderFirst)).Cols("user_id").Find(&userIds)
		utils.CheckErr(err)
		err = session.Db.Table("u_total_love_point_summary").
			Where(fmt.Sprintf("condition_%d > 0", condition)).
			OrderBy(fmt.Sprintf("condition_%d DESC, user_id", condition)).
			Limit(250, int(rankingOrderFirst)).Cols(fmt.Sprint("condition_", condition)).Find(&lovePoints)
		utils.CheckErr(err)
		for i := range userIds {
			orders = append(orders, 0)
			if i == 0 {
				orders[i] = findCommonRankingOrderByPoint(session, condition, lovePoints[i])
			} else if lovePoints[i] == lovePoints[i-1] {
				orders[i] = orders[i-1]
			} else {
				orders[i] = orders[0] + int32(i)
			}
			resp.LoveRankingData.Append(client.LoveRankingData{
				RankingUser: user_social.GetRankingUser(session, userIds[i]),
				LovePoint:   lovePoints[i],
				Order:       orders[i],
			})
		}
	}
	return resp
}

func findCommonRankingOrder(session *userdata.Session, condition int32) int32 {
	currentLovePoint := GetUserTotalLovePoint(session, session.UserId, condition)
	if currentLovePoint == 0 {
		return 0
	}
	return findCommonRankingOrderByPoint(session, condition, currentLovePoint)
}

func findCommonRankingOrderByPoint(session *userdata.Session, condition, totalPoint int32) int32 {
	if (condition >= enum.LoveRankingConditionTypeMember1) && (condition <= enum.LoveRankingConditionTypeMember212) {
		// single condition
		memberId := getMemberIdFromCondition(condition)
		count, err := session.Db.Table("u_member").Where("member_master_id = ? AND love_point > ?", memberId, totalPoint).Count()
		utils.CheckErr(err)
		return int32(count) + 1
	} else {
		generateTotalLovePointTable(session)
		count, err := session.Db.Table("u_total_love_point_summary").Where(fmt.Sprint("condition_", condition, " > ?"), totalPoint).Count()
		utils.CheckErr(err)
		return int32(count) + 1
	}
}
