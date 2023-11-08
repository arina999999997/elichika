package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveMemberMapping = map[int]LiveMemberMappingMember
type LiveMemberMappingMember struct {
	// from m_live_member_mapping
	// - mapping ID is repeated for the whole mapping
	// - MemberMasterID marked as pk to help loading for Live
	MappingID      *int `xorm:"pk 'mapping_id'"`
	Position       int  `xorm:"pk 'position'"`
	MemberMasterID int  `xorm:"'member_master_id'"`
	IsCenter       bool `xorm:"'is_center'"`
	CardPosition   int  `xorm:"'card_position'"`
	SuitMasterID   *int `xorm:"'suit_master_id'"`
}

func loadLiveMemberMapping(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveMemberMapping")
	memberMappings := []LiveMemberMappingMember{}
	err := masterdata_db.Table("m_live_member_mapping").Find(&memberMappings)
	utils.CheckErr(err)
	gamedata.LiveMemberMapping = make(map[int]LiveMemberMapping)
	for _, memberMapping := range memberMappings {
		_, exists := gamedata.LiveMemberMapping[*memberMapping.MappingID]
		if !exists {
			gamedata.LiveMemberMapping[*memberMapping.MappingID] = make(LiveMemberMapping)
		}
		gamedata.LiveMemberMapping[*memberMapping.MappingID][memberMapping.MemberMasterID] = memberMapping
	}
}

func init() {
	addLoadFunc(loadLiveMemberMapping)
}
