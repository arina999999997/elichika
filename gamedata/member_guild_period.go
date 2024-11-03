package gamedata

import (
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

func loadMemberGuildPeriod(gamedata *Gamedata) {
	fmt.Println("Loading MemberGuildPeriod")
	var exist bool
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		exist, err = session.Table("m_member_guild_period").Get(&gamedata.MemberGuildPeriod)
	})
	utils.CheckErrMustExist(err, exist)
}

func init() {
	addLoadFunc(loadMemberGuildPeriod)
}
