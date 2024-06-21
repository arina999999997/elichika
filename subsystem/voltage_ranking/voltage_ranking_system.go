package voltage_ranking

import (
	"elichika/generic/ranking"
	"elichika/userdata"
	"elichika/utils"
)

// TODO(threading): There is no lock here, because it's implicitly using the database lock.
type RankingType = ranking.Ranking[int32, int32]

var rankingByLiveDifficultyId = map[int32]*RankingType{}

func GetRankingByLiveDifficultyId(session *userdata.Session, liveDifficultyId int32) *RankingType {
	rank, exist := rankingByLiveDifficultyId[liveDifficultyId]
	if exist {
		return rank
	}
	rank = ranking.NewRanking[int32, int32]()
	type userIdScore struct {
		UserId       int32
		VoltagePoint int32
	}
	records := []userIdScore{}
	err := session.Db.Table("u_voltage_ranking").Where("live_difficulty_id = ?", liveDifficultyId).
		OrderBy("voltage_point DESC").Limit(VoltageRankingLimit).Find(&records)
	utils.CheckErr(err)

	for _, record := range records {
		rank.Update(record.UserId, record.VoltagePoint)
	}
	rankingByLiveDifficultyId[liveDifficultyId] = rank
	return rank
}
