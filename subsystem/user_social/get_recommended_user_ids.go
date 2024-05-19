package user_social

import (
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"

	"math/rand"
)

func GetRecommendedUserIds(session *userdata.Session) []int32 {
	// get at most random 20 user id who has the matching condition, then add them
	// we will use order_by last login to select, even though it's not explicitly mentioned
	// official server likely kept a list of recent player that it pick from, but we won't have that many user for that to be viable
	ids := []int32{}
	err := session.Db.Table("u_status").Where("rank >= ? AND last_login_at >= ?",
		session.Gamedata.ConstantInt[enum.ConstantIntFriendRecommendedPlayerConditionUserRank],
		session.Time.Unix()-int64(enum.DaySecondCount*session.Gamedata.ConstantInt[enum.ConstantIntFriendRecommendedPlayerConditionLastLoginDays]),
	).OrderBy("last_login_at DESC").
		Limit(int(session.Gamedata.ConstantInt[enum.ConstantIntFriendRecommendedPlayerConditionUserRank] * 2)).
		Cols("user_id").Find(&ids)
	utils.CheckErr(err)

	indices := rand.Perm(len(ids))
	res := []int32{}
	count := session.Gamedata.ConstantInt[enum.ConstantIntFriendRecommendedPlayerConditionUserRank]
	for _, i := range indices {
		if IsFriend(session, ids[i]) {
			continue
		}
		res = append(res, ids[i])
		count--
		if count <= 0 {
			break
		}
	}
	return res
}
