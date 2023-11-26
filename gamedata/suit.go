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
	// Name string `xorm:"'name'"`
	// ThumbnailImageAssetPath string `xorm:"'thumbnail_image_asset_path'"`
	SuitReleaseRoute int `xorm:"'suit_release_route'"`
	// ModelAssetPath string `xorm:"'model_asset_path'"`
	// DisplayOrder int `xorm:"'display_order'"`
}

func (suit *Suit) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	suit.Member = gamedata.Member[*suit.MemberMID]
	suit.MemberMID = &suit.Member.ID
	// suit.Name = dictionary.Resolve(suit.Name)
	// fmt.Println(suit.ID, "\t", *suit.MemberMID, "\t", suit.Name, "\t", suit.ThumbnailImageAssetPath, "\t", suit.ModelAssetPath)
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
