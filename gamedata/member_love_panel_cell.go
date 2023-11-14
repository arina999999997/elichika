package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberLovePanelCell struct {
	// from m_member_love_panel_cell
	ID int `xorm:"pk 'id'"`
	PanelIndex int `xorm:"'panel_index'"`
	MemberLovePanelMasterID *int `xorm:"'member_love_panel_master_id'"`
	MemberLovePanel *MemberLovePanel `xorm:"-"`
	// BonusType
	// BonusValue

	// from m_member_love_panel_cell_source_content
	Resources []model.Content `xorm:"-"`
}


func (cell *MemberLovePanelCell) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	cell.MemberLovePanel = gamedata.MemberLovePanel[*cell.MemberLovePanelMasterID]
	cell.MemberLovePanelMasterID = &cell.MemberLovePanel.ID
	err := masterdata_db.Table("m_member_love_panel_cell_source_content").Where("member_love_panel_cell_master_id = ?", cell.ID).Find(&cell.Resources)
	utils.CheckErr(err)
}


func loadMemberLovePanelCell(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading MemberLovePanelCell")
	gamedata.MemberLovePanelCell = make(map[int]*MemberLovePanelCell)
	err := masterdata_db.Table("m_member_love_panel_cell").Find(&gamedata.MemberLovePanelCell)
	utils.CheckErr(err)
	for _, cell := range gamedata.MemberLovePanelCell {
		cell.populate(gamedata, masterdata_db,serverdata_db,dictionary)
	}
}

func init() {
	addLoadFunc(loadMemberLovePanelCell)
	addPrequisite(loadMemberLovePanelCell, loadMemberLovePanel)
}