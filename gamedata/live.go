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

	// Name string
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
	for id, _ := range gamedata.Live {
		gamedata.Live[id].LiveMemberMapping = gamedata.LiveMemberMapping[*gamedata.Live[id].LiveMemberMappingID]
	}
}
