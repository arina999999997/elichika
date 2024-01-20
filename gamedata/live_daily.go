package gamedata

import (
	_ "elichika/clientdb"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveDaily struct {
	Id                   int32 `xorm:"pk 'id'"`
	LiveId               int32 `xorm:"'live_id'"`
	LimitCount           int32 `xorm:"'limit_count'"`
	MaxLimitCountRecover int32 `xorm:"'max_limit_count_recover'"`
	// AppealText string
	Weekday int32 `xorm:"'weekday'"`
	StartAt int64 `xorm:"'start_at'"` // maybe client only user int
	EndAt   int64 `xorm:"'end_at'"`
}

func loadLiveDaily(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveDaily")
	gamedata.LiveDaily = make(map[int32]*LiveDaily)
	err := masterdata_db.Table("m_live_daily").Find(&gamedata.LiveDaily)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadLiveDaily)
}
