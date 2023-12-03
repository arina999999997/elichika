package gamedata

import (
	_ "elichika/clientdb"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveDaily struct {
	ID                   int `xorm:"pk 'id'"`
	LiveID               int `xorm:"'live_id'"`
	LimitCount           int `xorm:"'limit_count'"`
	MaxLimitCountRecover int `xorm:"'max_limit_count_recover'"`
	// AppealText string
	Weekday int   `xorm:"'weekday'"`
	StartAt int64 `xorm:"'start_at'"` // maybe client only user int
	EndAt   int64 `xorm:"'end_at'"`
}

func loadLiveDaily(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveDaily")
	gamedata.LiveDaily = make(map[int]*LiveDaily)
	err := masterdata_db.Table("m_live_daily").Find(&gamedata.LiveDaily)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadLiveDaily)
}
