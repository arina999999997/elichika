package model

import (
	"elichika/client"
	"elichika/generic"

	"sort"
)

type UserCommunicationMemberDetailBadge struct {
	MemberMasterId     int  `xorm:"pk 'member_master_id'" json:"member_master_id"`
	IsStoryMemberBadge bool `xorm:"'is_story_member_badge'" json:"is_story_member_badge"`
	IsStorySideBadge   bool `xorm:"'is_story_side_badge'" json:"is_story_side_badge"`
	IsVoiceBadge       bool `xorm:"'is_voice_badge'" json:"is_voice_badge"`
	IsThemeBadge       bool `xorm:"'is_theme_badge'" json:"is_theme_badge"`
	IsCardBadge        bool `xorm:"'is_card_badge'" json:"is_card_badge"`
	IsMusicBadge       bool `xorm:"'is_music_badge'" json:"is_music_badge"`
}

func (ucmdb *UserCommunicationMemberDetailBadge) Id() int64 {
	return int64(ucmdb.MemberMasterId)
}

// UserMember ...

type MemberPublicInfo struct {
	MemberMasterId       int `xorm:"pk 'member_master_id'" json:"member_master_id"`
	LoveLevel            int `json:"love_level"`
	LovePointLimit       int `json:"love_point_limit"`
	OwnedCardCount       int `json:"owned_card_count"`
	AllTrainingCardCount int `json:"all_training_activated_card_count"`
}

// Bond board (Love Panel)
type UserMemberLovePanel struct {
	// love panel level = m_member_love_panel[id] / 1000
	// SELECT * FROM m_member_love_panel_cell WHERE id != (member_love_panel_master_id / 1000 - 1) * 10000 + (panel_index + 1) * 1000 + (member_love_panel_master_id % 1000); -> 0
	MemberId                  int32   `xorm:"pk <- 'member_master_id'" json:"member_id"` // member_love_panel_master_id % 1000
	MemberLovePanelCellIds    []int32 `xorm:"-" json:"member_love_panel_cell_ids"`
	LovePanelLevel            int32   `json:"-"` // member_love_panel_master_id / 1000
	LovePanelLastLevelCellIds []int32 `xorm:"'love_panel_last_level_cell_ids'" json:"-"`
	// there is no ambiguous representation
	// - When the last level is filled, if there is a next level then LovePanelLevel is increased, and LovePanelLastLevelCellIds is cleared
	// - otherwise, LovePanelLevel stay the same, and LovePanelLastLevelCellIds has 5 tiles.
}

func (x *UserMemberLovePanel) Normalize() {
	if x.LovePanelLevel > 0 { // already calculated
		return
	}
	l := len(x.MemberLovePanelCellIds)
	if l == 0 {
		return
	}
	sort.Slice(x.MemberLovePanelCellIds, func(i, j int) bool { return x.MemberLovePanelCellIds[i] < x.MemberLovePanelCellIds[j] })
	i := l - l%5
	if i == l {
		i -= 5
	}
	x.LovePanelLevel = int32((i / 5) + 1)
	x.LovePanelLastLevelCellIds = x.MemberLovePanelCellIds[i:l]
}

func (x *UserMemberLovePanel) Fill() {
	x.MemberLovePanelCellIds = []int32{} // [] instead of null
	for l := int32(1); l < x.LovePanelLevel; l++ {
		for cell := int32(1000); cell <= 5000; cell += 1000 {
			x.MemberLovePanelCellIds = append(x.MemberLovePanelCellIds, (l-1)*10000+cell+x.MemberId)
		}
	}
	x.MemberLovePanelCellIds = append(x.MemberLovePanelCellIds, x.LovePanelLastLevelCellIds...)
}

func (x *UserMemberLovePanel) LevelUp() {
	if len(x.LovePanelLastLevelCellIds) != 5 {
		panic("incorrect level up")
	}
	x.LovePanelLastLevelCellIds = []int32{}
	x.LovePanelLevel++
}

func init() {

	type DbMember struct {
		client.UserMember         `xorm:"extends"`
		LovePanelLevel            int   `xorm:"'love_panel_level' default 1"`
		LovePanelLastLevelCellIds []int `xorm:"'love_panel_last_level_cell_ids' default '[]'"`
	}
	TableNameToInterface["u_member"] = generic.UserIdWrapper[DbMember]{}
	TableNameToInterface["u_communication_member_detail_badge"] = generic.UserIdWrapper[UserCommunicationMemberDetailBadge]{}
}
