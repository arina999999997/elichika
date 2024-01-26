package voltage_ranking

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/subsystem/cache"
	"elichika/subsystem/user_profile"
	"elichika/userdata"
	"elichika/utils"
)

var (
	getVoltageRankingResponseCache = cache.UniquePointerMap[int32, cache.CachedObject[response.GetVoltageRankingResponse]]{}
)

func GetVoltageRankingResponse(session *userdata.Session, liveDifficultyId int32) *response.GetVoltageRankingResponse {
	cacher := getVoltageRankingResponseCache.Get(liveDifficultyId)
	cacher.Acquire()
	defer cacher.Release()
	if cacher.ExpireAt > session.Time.Unix() {
		return cacher.Value
	}
	cacher.ExpireAt = session.Time.Unix() + VoltageRankingCache
	cacher.Value = getVoltageRankingResponseNoCache(session, liveDifficultyId)
	return cacher.Value
}

func getVoltageRankingResponseNoCache(session *userdata.Session, liveDifficultyId int32) *response.GetVoltageRankingResponse {
	resp := response.GetVoltageRankingResponse{}

	type userIdScore struct {
		UserId       int32
		VoltagePoint int32
	}
	records := []userIdScore{}
	err := session.Db.Table("u_voltage_ranking").Where("live_difficulty_id = ?", liveDifficultyId).
		OrderBy("voltage_point DESC").Limit(VoltageRankingLimit).Find(&records)
	utils.CheckErr(err)

	for i, record := range records {
		resp.VoltageRankingCells.Append(client.VoltageRankingCell{
			Order:              int32(i + 1),
			VoltagePoint:       record.VoltagePoint,
			VoltageRankingUser: user_profile.GetVoltageRankingUser(session, record.UserId),
		})
	}
	return &resp
}
