package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Live struct {
	// from m_live
	LiveID int `xorm:"pk 'live_id'"`
	// Is2DLive bool `xorm:"'is_2d_live'"`
	// MusicID *int `xorm:"'music_id'"`
	// BgmPath string `xorm:"'bgm_path'"`
	// ChorusBgmPath string `xorm:"'chorus_bgm_path'"`
	LiveMemberMapping   LiveMemberMapping `xorm:"-"`
	LiveMemberMappingID *int              `xorm:"'live_member_mapping_id'"`

	Name string `xorm:"'name'"`
	// Pronunciation string
	MemberGroup int  `xorm:"'member_group'"`
	MemberUnit  *int `xorm:"'member_unit'"`
	// OriginalDeckName string
	// Copyright string
	// Source *string
	// JacketAssetPath string
	// BackgroundAssetPath string
	// DisplayOrder int

	// from m_live_difficulty
	LiveDifficulties []*LiveDifficulty `xorm:"-"`

	// from m_live, m_live_difficulty, and dictionary

	Gamedata *Gamedata `xorm:"-"`
}

func init() {
	addLoadFunc(loadLive)
	addPrequisite(loadLive, loadLiveMemberMapping)
}

func loadLive(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Live")
	gamedata.Live = make(map[int]*Live)
	err := masterdata_db.Table("m_live").Find(&gamedata.Live)
	utils.CheckErr(err)
	for _, live := range gamedata.Live {
		live.LiveMemberMapping = gamedata.LiveMemberMapping[*live.LiveMemberMappingID]
		live.Name = dictionary.Resolve(live.Name)
		live.Gamedata = gamedata
	}
}
