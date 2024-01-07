package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberLoginBonusBirthday struct {
	Id int `xorm:"pk 'id'"`
	// StartAt int64
	// EndAt int64
	SuitMasterId int `xorm:"'suit_master_id'"`
}

func loadMemberLoginBonusBirthday(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberLoginBonusBirthday")
	bonuses := []MemberLoginBonusBirthday{}
	err := masterdata_db.Table("m_login_bonus_birthday").OrderBy("id DESC").Find(&bonuses)
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
