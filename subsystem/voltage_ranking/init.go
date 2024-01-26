package voltage_ranking

import (
	"elichika/client"
	"elichika/userdata/database"
)

// we have to save the whole deck because we should return the details when the record was created, not as it is currently
type UserVoltageRanking struct {
	UserId           int32 `xorm:"pk"`
	LiveDifficultyId int32 `xorm:"pk"`
	VoltagePoint     int32
	DeckDetail       client.OtherUserDeckDetail `xorm:"json"`
}

func init() {
	database.AddTable("u_voltage_ranking", UserVoltageRanking{})
}
