package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Suit struct {
	// from m_suit
	ID        int     `xorm:"pk 'id'"`
	MemberMID *int    `xorm:"'member_m_id'"`
	Member    *Member `xorm:"-"`
	// Name string
	// ThumbnailImageAssetPath string
	SuitReleaseRoute int `xorm:"'suit_release_route'"`
	// ModelAssetPath string
	// DisplayOrder int `xorm:"'display_order'"`
}

func (suit *Suit) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	suit.Member = gamedata.Member[*suit.MemberMID]
	suit.MemberMID = &suit.Member.ID
}

func loadSuit(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Suit")
	gamedata.Suit = make(map[int]*Suit)
	err := masterdata_db.Table("m_suit").Find(&gamedata.Suit)
	utils.CheckErr(err)
	for _, suit := range gamedata.Suit {
		suit.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadSuit)
	addPrequisite(loadSuit, loadMember)
}
