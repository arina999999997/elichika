package mining

import (
	"elichika/gamedata"
	"elichika/generic/ranking"
	"elichika/utils"

	"xorm.io/xorm"
)

// TODO(threading): There is no lock here, because it's implicitly using the database lock
type PointRankingType = ranking.Ranking[int32, int32]

// this get invalidated everytime a new event start
// failure to setup the schedule properly will return in bad data
var eventMiningPointRanking *PointRankingType = nil

// note that eventID is necessary as sometime, we need to get the data after the event has ended, but before it was clean up
// to make it convinience to use, if eventId is passed as a negative value, get from the active event
func GetPointRanking(userdata_db *xorm.Session, eventId int32) *PointRankingType {
	if eventId < 0 {
		eventId = gamedata.Instance.EventActive.GetEventMining().EventId
	}
	if eventMiningPointRanking != nil {
		return eventMiningPointRanking
	}
	eventMiningPointRanking = ranking.NewRanking[int32, int32]()
	type UserIdEventPoint struct {
		UserId     int32 `xorm:"'user_id'"`
		EventPoint int32 `xorm:"'event_point'"`
	}
	records := []UserIdEventPoint{}

	err := userdata_db.Table("u_event_mining").Where("event_master_id = ? AND event_point > 0", eventId).Find(&records)
	utils.CheckErr(err)

	for _, record := range records {
		eventMiningPointRanking.Update(record.UserId, record.EventPoint)
	}
	return eventMiningPointRanking
}

func ResetPointRanking() {
	eventMiningPointRanking = nil
}
