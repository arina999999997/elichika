package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Suit struct {
	// from m_suit
	Id        int32   `xorm:"pk 'id'"`
	MemberMId *int32  `xorm:"'member_m_id'"`
	Member    *Member `xorm:"-"`
	// Name string `xorm:"'name'"`
	// ThumbnailImageAssetPath string `xorm:"'thumbnail_image_asset_path'"`
	SuitReleaseRoute int32 `xorm:"'suit_release_route'" enum:"SuitReleaseRoute"`
	// ModelAssetPath string `xorm:"'model_asset_path'"`
	// DisplayOrder int `xorm:"'display_order'"`
}

func (suit *Suit) populate(gamedata *Gamedata) {
	suit.Member = gamedata.Member[*suit.MemberMId]
	suit.MemberMId = &suit.Member.Id
	// suit.Name = gamedata.Dictionary.Resolve(suit.Name)
	// fmt.Println(suit.Id, "\t", *suit.MemberMId, "\t", suit.Name, "\t", suit.ThumbnailImageAssetPath, "\t", suit.ModelAssetPath)
}

func loadSuit(gamedata *Gamedata) {
	fmt.Println("Loading Suit")
	gamedata.Suit = make(map[int32]*Suit)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_suit").Find(&gamedata.Suit)
	})
	utils.CheckErr(err)
	for _, suit := range gamedata.Suit {
		suit.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadSuit)
	addPrequisite(loadSuit, loadMember)
}
