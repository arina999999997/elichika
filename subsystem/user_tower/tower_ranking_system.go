package user_tower

import (
	"elichika/generic/ranking"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

// the ranking system work as follow:
// - use elichika/generic/ranking to keep the ranking object in memory
// - the ranking is built from database if not available, otherwise it's updated in both memory and database

// TODO(threading): There is no lock here, because it's implicitly using the database lock.
// If we change how the database lock work, we need to add lock too
type RankingType = ranking.Ranking[int32, int32]

var rankingByTowerId = map[int32]*RankingType{}

func GetRankingByTowerId(session *userdata.Session, towerId int32) *RankingType {
	rank, exist := rankingByTowerId[towerId]
	if exist {
		return rank
	}

	// fetch from database and construct it
	records := []database.UserTowerVoltageRankingScore{}
	err := session.Db.Table("u_tower_voltage_ranking_score").Where("tower_id = ?", towerId).Find(&records)
	utils.CheckErr(err)
	totalVoltage := map[int32]int32{}
	for _, record := range records {
		totalVoltage[record.UserId] += record.Voltage
	}
	rank = ranking.NewRanking[int32, int32]()
	for userId, score := range totalVoltage {
		rank.Update(userId, score)
	}
	rankingByTowerId[towerId] = rank
	return rank
}
