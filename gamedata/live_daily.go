package gamedata

import (
	_ "elichika/clientdb"

	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveDaily struct {
	// from m_live_daily
	Id                   int32 `xorm:"pk 'id'"`
	LiveId               int32 `xorm:"'live_id'"`
	LimitCount           int32 `xorm:"'limit_count'"`
	MaxLimitCountRecover int32 `xorm:"'max_limit_count_recover'"`
	// AppealText string
	Weekday int32 `xorm:"'weekday'"`
	StartAt int64 `xorm:"'start_at'"` // maybe client only uses int
	EndAt   int64 `xorm:"'end_at'"`

	// from m_live_daily_difficulty
	DropContentGroupId     int32          `xorm:"-"`
	DropContentGroup       *LiveDropGroup `xorm:"-"`
	RareDropContentGroupId int32          `xorm:"-"`
	RareDropContentGroup   *LiveDropGroup `xorm:"-"`
}

func (ld *LiveDaily) populate(gamedata *Gamedata) {
	var exist bool
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		exist, err = session.Table("m_live_daily_difficulty").Where("live_daily_id = ?", ld.Id).Cols("drop_content_group_id", "rare_drop_content_group_id").Get(&ld.DropContentGroupId, &ld.RareDropContentGroupId)
	})
	utils.CheckErrMustExist(err, exist)
	ld.DropContentGroup = gamedata.LiveDropGroup[ld.DropContentGroupId]
	ld.DropContentGroup.Check()
	ld.RareDropContentGroup = gamedata.LiveDropGroup[ld.RareDropContentGroupId]
	ld.RareDropContentGroup.Check()
}

func loadLiveDaily(gamedata *Gamedata) {
	fmt.Println("Loading LiveDaily")
	gamedata.LiveDaily = make(map[int32]*LiveDaily)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_live_daily").Find(&gamedata.LiveDaily)
	})
	utils.CheckErr(err)
	for _, liveDaily := range gamedata.LiveDaily {
		liveDaily.populate(gamedata)
		gamedata.Live[liveDaily.LiveId].LiveDailies = append(gamedata.Live[liveDaily.LiveId].LiveDailies, liveDaily)
	}

}

func init() {
	addLoadFunc(loadLiveDaily)
	addPrequisite(loadLiveDaily, loadLive)
	addPrequisite(loadLiveDaily, loadLiveDropGroup)
}
