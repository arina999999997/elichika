package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveMemberMapping = map[int32]LiveMemberMappingMember
type LiveMemberMappingMember struct {
	// from m_live_member_mapping and m_live_override_member_mapping
	// - mapping Id is repeated for the whole mapping
	// - MemberMasterId marked as pk to help loading for Live
	MappingId *int32 `xorm:"pk 'mapping_id'"`
	Position  int32  `xorm:"pk 'position'"`
	// MemberMasterId *int32  `xorm:"'member_master_id'"`
	// MemberNonPlayableMasterId *int32  `xorm:"'member_non_playable_master_id'"`
	IsCenter     bool  `xorm:"'is_center'"`
	CardPosition int32 `xorm:"'card_position'"`
	// SuitMasterId   *int32 `xorm:"'suit_master_id'"`
}

func loadLiveMemberMapping(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveMemberMapping")
	gamedata.LiveMemberMapping = make(map[int32]LiveMemberMapping)

	tables := []string{"m_live_member_mapping", "m_live_override_member_mapping"}
	for _, table := range tables {
		memberMappings := []LiveMemberMappingMember{}
		err := masterdata_db.Table(table).Find(&memberMappings)
		utils.CheckErr(err)
		for _, memberMapping := range memberMappings {
			_, exist := gamedata.LiveMemberMapping[*memberMapping.MappingId]
			if !exist {
				gamedata.LiveMemberMapping[*memberMapping.MappingId] = make(LiveMemberMapping)
			}
			gamedata.LiveMemberMapping[*memberMapping.MappingId][memberMapping.Position] = memberMapping
		}
	}

}

func init() {
	addLoadFunc(loadLiveMemberMapping)
}
