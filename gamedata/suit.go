package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Suit struct {
	// from m_suit
	Id        int32     `xorm:"pk 'id'"`
	MemberMId *int32  `xorm:"'member_m_id'"`
	Member    *Member `xorm:"-"`
	// Name string `xorm:"'name'"`
	// ThumbnailImageAssetPath string `xorm:"'thumbnail_image_asset_path'"`
	SuitReleaseRoute int32 `xorm:"'suit_release_route'" enum:"SuitReleaseRoute"`
	// ModelAssetPath string `xorm:"'model_asset_path'"`
	// DisplayOrder int `xorm:"'display_order'"`
}

func (suit *Suit) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	suit.Member = gamedata.Member[*suit.MemberMId]
	suit.MemberMId = &suit.Member.Id
	// suit.Name = dictionary.Resolve(suit.Name)
	// fmt.Println(suit.Id, "\t", *suit.MemberMId, "\t", suit.Name, "\t", suit.ThumbnailImageAssetPath, "\t", suit.ModelAssetPath)
}

func loadSuit(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Suit")
	gamedata.Suit = make(map[int32]*Suit)
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
