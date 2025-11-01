package mining

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

func GetEventMiningUserRanking(session *userdata.Session) client.EventMiningUserRanking {
	pointRanking, pointRanked := GetPointRanking(session.Db, -1).TiedRankOf(session.UserId)
	voltageRanking, voltageRanked := GetVoltageRanking(session.Db, -1).TiedRankOf(session.UserId)
	eventMiningUserRanking := client.EventMiningUserRanking{}
	if pointRanked {
		eventMiningUserRanking.PointRankingOrder = int32(pointRanking)
		point, _ := GetPointRanking(session.Db, -1).ScoreOf(session.UserId)
		eventMiningUserRanking.PointRankingTotalPoint = generic.NewNullable(point)
	}
	if voltageRanked {
		eventMiningUserRanking.VoltageRankingOrder = int32(voltageRanking)
		voltage, _ := GetVoltageRanking(session.Db, -1).ScoreOf(session.UserId)
		eventMiningUserRanking.VoltageRankingPoint = generic.NewNullable(voltage)
	}
	return eventMiningUserRanking
}
