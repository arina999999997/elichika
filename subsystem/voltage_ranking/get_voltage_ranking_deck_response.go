package voltage_ranking

import (
	"elichika/client/response"
	"elichika/subsystem/cache"
	"elichika/subsystem/user_profile"
	"elichika/userdata"
	"elichika/utils"
)

var (
	getVoltageRankingDeckResponseCache = cache.UniquePointerMap[int64, cache.CachedObject[response.GetVoltageRankingDeckResponse]]{}
)

func GetVoltageRankingDeckResponse(session *userdata.Session, liveDifficultyId int32, userId int32) *response.GetVoltageRankingDeckResponse {
	key := (int64(liveDifficultyId) << 32) ^ (int64(userId))
	cacher := getVoltageRankingDeckResponseCache.Get(key)
	cacher.Acquire()
	defer cacher.Release()
	if cacher.ExpireAt > session.Time.Unix() {
		return cacher.Value
	}

	cacher.ExpireAt = session.Time.Unix() + VoltageRankingDeckCache
	cacher.Value = getVoltageRankingDeckResponseNoCache(session, liveDifficultyId, userId)
	return cacher.Value
}

func getVoltageRankingDeckResponseNoCache(session *userdata.Session, liveDifficultyId int32, userId int32) *response.GetVoltageRankingDeckResponse {
	resp := response.GetVoltageRankingDeckResponse{}

	exist, err := session.Db.Table("u_voltage_ranking").
		Where("live_difficulty_id = ? AND user_id = ?", liveDifficultyId, userId).
		Cols("deck_detail").Get(&resp)
	utils.CheckErrMustExist(err, exist)
	resp.User = user_profile.GetOtherUser(session, userId)
	return &resp
}
