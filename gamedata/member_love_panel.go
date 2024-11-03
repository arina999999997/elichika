package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type MemberLovePanelBonus struct {
	BonusType  int32 `xorm:"bonus_type" enum:"MemberLovePanelEffectType"`
	BonusValue int32 `xorm:"bonus_value"`
}
type MemberLovePanel struct {
	// from m_member_love_panel
	Id                       int32            `xorm:"pk 'id'"`
	LoveLevelMasterLoveLevel int32            `xorm:"'love_level_master_love_level'"`
	MemberMasterId           *int32           `xorm:"member_master_id"`
	Member                   *Member          `xorm:"-"`
	NextPanel                *MemberLovePanel `xorm:"-"`

	// from m_member_love_panel_bonus
	Bonuses []MemberLovePanelBonus `xorm:"-"`
}

func (panel *MemberLovePanel) populate(gamedata *Gamedata) {
	panel.Member = gamedata.Member[*panel.MemberMasterId]
	panel.MemberMasterId = &panel.Member.Id
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_love_panel_bonus").Where("member_love_panel_master_id = ?", panel.Id).Find(&panel.Bonuses)
	})
	utils.CheckErr(err)
}

func loadMemberLovePanel(gamedata *Gamedata) {
	fmt.Println("Loading MemberLovePanel")
	gamedata.MemberLovePanel = make(map[int32]*MemberLovePanel)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_love_panel").Find(&gamedata.MemberLovePanel)
	})
	utils.CheckErr(err)
	for _, panel := range gamedata.MemberLovePanel {
		panel.populate(gamedata)
	}
	memberLovePanels := []MemberLovePanel{}
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_member_love_panel").OrderBy("member_master_id, love_level_master_love_level").Find(&memberLovePanels)
	})
	utils.CheckErr(err)
	for i := len(memberLovePanels) - 2; i >= 0; i-- {
		id := memberLovePanels[i].Id
		nId := memberLovePanels[i+1].Id
		if *gamedata.MemberLovePanel[id].MemberMasterId == *gamedata.MemberLovePanel[nId].MemberMasterId {
			gamedata.MemberLovePanel[id].NextPanel = gamedata.MemberLovePanel[nId]
		}
	}
}

func init() {
	addLoadFunc(loadMemberLovePanel)
	addPrequisite(loadMemberLovePanel, loadMember)
	addPrequisite(loadMemberLovePanel, loadMemberLoveLevel)
}
