package user_tower

import (
	"elichika/client"
	"elichika/userdata"
)

func GetTowerRankingCell(session *userdata.Session, towerId int32) client.TowerRankingCell {
	scores := GetUserTowerVoltageRankingScores(session, towerId)
	cell := client.TowerRankingCell{
		Order:            1,
		SumVoltage:       0,
		TowerRankingUser: GetTowerRankingUser(session, session.UserId),
	}
	for _, score := range scores {
		cell.SumVoltage += score.Voltage
	}
	return cell
}
