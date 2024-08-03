package marathon

import (
	"elichika/userdata"
)

// return the ranking, or 0 if not present
func GetUserEventMarathonRanking(session *userdata.Session, eventId int32) int32 {
	rank, _ := GetRanking(session.Db, eventId).TiedRankOf(session.UserId)
	return int32(rank)
}
