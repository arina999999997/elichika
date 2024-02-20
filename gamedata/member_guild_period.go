package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGuildPeriod struct {
	Id                int
	StartAt           int64
	EndAt             *int64
	TransferStartSecs int64
	TransferEndSecs   int64
	RankingStartSecs  int64
	RankingEndSecs    int64
	OneCycleSecs      int64
}

func loadMemberGuildPeriod(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberGuildPeriod")
	exist, err := masterdata_db.Table("m_member_guild_period").Get(&gamedata.MemberGuildPeriod)
	utils.CheckErrMustExist(err, exist)
}

func init() {
	addLoadFunc(loadMemberGuildPeriod)
}
