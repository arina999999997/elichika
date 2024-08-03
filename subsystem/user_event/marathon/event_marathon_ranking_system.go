package marathon

import (
	"elichika/generic/ranking"
	"elichika/utils"

	"xorm.io/xorm"
)

// TODO(threading): There is no lock here, because it's implicitly using the database lock
type RankingType = ranking.Ranking[int32, int32]

// this get invalidated everytime a new event start
// failure to setup the schedule properly will return in bad data
var eventMarathonRanking *RankingType = nil

func GetRanking(userdata_db *xorm.Session, eventId int32) *RankingType {
	if eventMarathonRanking != nil {
		return eventMarathonRanking
	}
	eventMarathonRanking = ranking.NewRanking[int32, int32]()
	type userIdEp struct {
		UserId     int32
		EventPoint int32
	}
	records := []userIdEp{}
	err := userdata_db.Table("u_event_marathon").Where("event_master_id = ? AND event_point > 0", eventId).Find(&records)
	utils.CheckErr(err)

	for _, record := range records {
		eventMarathonRanking.Update(record.UserId, record.EventPoint)
	}
	return eventMarathonRanking
}

func ResetRanking() {
	eventMarathonRanking = nil
}
