package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGroup struct {
	// from m_MemberGroup_group
	MemberGroup int32  `xorm:"pk 'member_group'" enum:"MemberGroup"` // muse aqour niji
	GroupName   string `xorm:"'group_name'"`
}

func (mg *MemberGroup) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	mg.GroupName = dictionary.Resolve(mg.GroupName)
}

func loadMemberGroup(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberGroup")
	gamedata.MemberGroup = make(map[int32]*MemberGroup)
	err := masterdata_db.Table("m_member_group").Find(&gamedata.MemberGroup)
	utils.CheckErr(err)
	for _, mg := range gamedata.MemberGroup {
		mg.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadMemberGroup)
}
