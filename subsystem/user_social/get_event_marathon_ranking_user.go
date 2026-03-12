package user_social

import (
	"elichika/client"
	"elichika/subsystem/cache"
	"elichika/userdata"
	"elichika/utils"
)

var (
	getEventMarathonRankingUserCache = cache.UniquePointerMap[int32, cache.CachedObject[client.EventMarathonRankingUser]]{}
)

func GetEventMarathonRankingUser(session *userdata.Session, userId int32) client.EventMarathonRankingUser {
	cacher := getEventMarathonRankingUserCache.Get(userId)
	cacher.Acquire()
	defer cacher.Release()
	if cacher.ExpireAt <= session.Time.Unix() {
		cacher.ExpireAt = session.Time.Unix() + EventMarathonRankingUserCache
		cacher.Value = getEventMarathonRankingUserNoCache(session, userId)
	}
	return *cacher.Value
}

func getEventMarathonRankingUserNoCache(session *userdata.Session, userId int32) *client.EventMarathonRankingUser {
	user := client.EventMarathonRankingUser{}
	exist, err := session.Db.Table("u_status").Where("user_id = ?", userId).Cols(
		"user_id", "name", "rank", "recommend_card_master_id", "emblem_id").
		Get(&user.UserId, &user.UserName, &user.UserRank, &user.CardMasterId, &user.EmblemMasterId)
	utils.CheckErrMustExist(err, exist)
	exist, err = session.Db.Table("u_card").Where("user_id = ? AND card_master_id = ?", userId, user.CardMasterId).
		Cols("level", "is_awakening_image", "is_all_training_activated").
		Get(&user.Level, &user.IsAwakening, &user.IsAllTrainingActivated)
	utils.CheckErrMustExist(err, exist)
	return &user
}
