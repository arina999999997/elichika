package user_social

import (
	"elichika/client"
	"elichika/subsystem/cache"
	"elichika/userdata"
	"elichika/utils"
)

var (
	getRankingUserCache = cache.UniquePointerMap[int64, cache.CachedObject[client.RankingUser]]{}
)

func GetRankingUser(session *userdata.Session, rankingUserId int32) client.RankingUser {
	key := (int64(rankingUserId) << 32)
	cacher := getRankingUserCache.Get(key)
	cacher.Acquire()
	defer cacher.Release()
	if cacher.ExpireAt <= session.Time.Unix() {
		cacher.ExpireAt = session.Time.Unix() + RankingUserCache
		cacher.Value = getRankingUserNoCache(session, rankingUserId)
	}
	return *cacher.Value
}

func getRankingUserNoCache(session *userdata.Session, rankingUserId int32) *client.RankingUser {
	rankingUser := client.RankingUser{}

	exist, err := session.Db.Table("u_status").Where("user_id = ?", rankingUserId).Cols(
		"user_id", "name", "rank", "recommend_card_master_id", "emblem_id").Get(
		&rankingUser.UserId, &rankingUser.Name, &rankingUser.Rank, &rankingUser.FavoriteCardMasterId, &rankingUser.EmblemId)
	utils.CheckErrMustExist(err, exist)

	exist, err = session.Db.Table("u_card").Where("user_id = ? AND card_master_id = ?", rankingUserId, rankingUser.FavoriteCardMasterId).Cols(
		"level", "is_awakening_image", "is_all_training_activated").Get(
		&rankingUser.FavoriteCardLevel, &rankingUser.IsAwakeningImage, &rankingUser.IsAllTrainingActivated)
	utils.CheckErrMustExist(err, exist)

	return &rankingUser
}
