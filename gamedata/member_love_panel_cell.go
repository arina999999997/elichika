package gamedata

import (
	"elichika/client"

	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberLovePanelCell struct {
	// from m_member_love_panel_cell
	Id                      int32            `xorm:"pk 'id'"`
	PanelIndex              int32            `xorm:"'panel_index'"`
	MemberLovePanelMasterId *int32           `xorm:"'member_love_panel_master_id'"`
	MemberLovePanel         *MemberLovePanel `xorm:"-"`
	BonusType               int32            `xorm:"bonus_type" enum:"MemberLovePanelEffectType"`
	BonusValue              int32            `xorm:"bonus_value"`

	// from m_member_love_panel_cell_source_content
	Resources []client.Content `xorm:"-"`
}

func (cell *MemberLovePanelCell) populate(gamedata *Gamedata) {
	cell.MemberLovePanel = gamedata.MemberLovePanel[*cell.MemberLovePanelMasterId]
	cell.MemberLovePanelMasterId = &cell.MemberLovePanel.Id
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_love_panel_cell_source_content").Where("member_love_panel_cell_master_id = ?", cell.Id).Find(&cell.Resources)
	})
	utils.CheckErr(err)
}

func loadMemberLovePanelCell(gamedata *Gamedata) {
	fmt.Println("Loading MemberLovePanelCell")
	gamedata.MemberLovePanelCell = make(map[int32]*MemberLovePanelCell)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_love_panel_cell").Find(&gamedata.MemberLovePanelCell)
	})
	utils.CheckErr(err)
	for _, cell := range gamedata.MemberLovePanelCell {
		cell.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadMemberLovePanelCell)
	addPrequisite(loadMemberLovePanelCell, loadMemberLovePanel)
}
