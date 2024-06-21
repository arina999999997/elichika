package database

type UserTowerVoltageRankingScore struct {
	UserId  int32 `xorm:"pk 'user_id'"`
	TowerId int32 `xorm:"pk 'tower_id'"`
	FloorNo int32 `xorm:"pk 'floor_no'"`
	Voltage int32 `xorm:"'voltage'"`
}

func init() {
	AddTable("u_tower_voltage_ranking_score", UserTowerVoltageRankingScore{})
}
