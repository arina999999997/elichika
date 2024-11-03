package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberLoginBonusBirthday struct {
	Id int32 `xorm:"pk 'id'"`
	// StartAt int64
	// EndAt int64
	SuitMasterId int32 `xorm:"'suit_master_id'"`
}

func loadMemberLoginBonusBirthday(gamedata *Gamedata) {
	fmt.Println("Loading MemberLoginBonusBirthday")
	bonuses := []MemberLoginBonusBirthday{}
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_login_bonus_birthday").OrderBy("id DESC").Find(&bonuses)
	})
	utils.CheckErr(err)
	for _, memberLoginBonusBirthday := range bonuses {
		gamedata.Member[gamedata.Suit[memberLoginBonusBirthday.SuitMasterId].Member.Id].MemberLoginBonusBirthdays = append(
			gamedata.Member[gamedata.Suit[memberLoginBonusBirthday.SuitMasterId].Member.Id].MemberLoginBonusBirthdays,
			memberLoginBonusBirthday)
	}
}

func init() {
	addLoadFunc(loadMemberLoginBonusBirthday)
	addPrequisite(loadMemberLoginBonusBirthday, loadSuit)
	addPrequisite(loadMemberLoginBonusBirthday, loadMember)
}
