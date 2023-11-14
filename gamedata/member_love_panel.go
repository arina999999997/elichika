package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberLovePanel struct {
	// from m_member_love_panel
	ID int `xorm:"pk 'id'"`
	LoveLevelMasterLoveLevel int `xorm:"'love_level_master_love_level'"`
	MemberMasterID *int `xorm:"member_master_id"`
	Member *Member `xorm:"-"`
	NextPanel *MemberLovePanel `xorm:"-"`
}

func (panel *MemberLovePanel) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	panel.Member = gamedata.Member[*panel.MemberMasterID]
	panel.MemberMasterID = &panel.Member.ID
}


func loadMemberLovePanel(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberLovePanel")
	gamedata.MemberLovePanel = make(map[int]*MemberLovePanel)
	err := masterdata_db.Table("m_member_love_panel").Find(&gamedata.MemberLovePanel)
	utils.CheckErr(err)
	for _, panel := range gamedata.MemberLovePanel {
		panel.populate(gamedata, masterdata_db,serverdata_db,dictionary)
	}
	memberLovePanels := []MemberLovePanel{}
	err = masterdata_db.Table("m_member_love_panel").OrderBy("member_master_id, love_level_master_love_level").Find(&memberLovePanels)
	utils.CheckErr(err)
	for i := len(memberLovePanels) - 2; i >= 0; i-- {
		id := memberLovePanels[i].ID
		nID := memberLovePanels[i + 1].ID
		if *gamedata.MemberLovePanel[id].MemberMasterID == *gamedata.MemberLovePanel[nID].MemberMasterID {
			gamedata.MemberLovePanel[id].NextPanel = gamedata.MemberLovePanel[nID]
		}
	}
}


func init() {
	addLoadFunc(loadMemberLovePanel)
	addPrequisite(loadMemberLovePanel, loadMember)
	addPrequisite(loadMemberLovePanel, loadMemberLoveLevel)
}