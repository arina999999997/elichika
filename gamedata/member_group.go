package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberGroup struct {
	// from m_MemberGroup_group
	MemberGroup int32  `xorm:"pk 'member_group'" enum:"MemberGroup"` // muse aqour niji
	GroupName   string `xorm:"'group_name'"`
}

func (mg *MemberGroup) populate(gamedata *Gamedata) {
	mg.GroupName = gamedata.Dictionary.Resolve(mg.GroupName)
}

func loadMemberGroup(gamedata *Gamedata) {
	fmt.Println("Loading MemberGroup")
	gamedata.MemberGroup = make(map[int32]*MemberGroup)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_group").Find(&gamedata.MemberGroup)
	})
	utils.CheckErr(err)
	for _, mg := range gamedata.MemberGroup {
		mg.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadMemberGroup)
}
