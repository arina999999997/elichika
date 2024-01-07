package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryLevelUpItem struct {
	// frm m_accessory_level_up_item
	Id int `xorm:"pk 'id'"`
	// Name string
	// ThumbnailAssetPath
	Rarity    int `xorm:"'rarity'"`
	Attribute int `xorm:"'attribute'"`
	PlusExp   int `xorm:"'plus_exp'"`
	GameMoney int `xorm:"'game_money'"`
	// Description string
	ItemListCategoryType int `xorm:"'item_list_category_type'"`
	// SceneId int
	// BannerImageAssetPath
	// DisplayOrder int
}

func loadAccessoryLevelUpItem(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryLevelUpItem = make(map[int]*AccessoryLevelUpItem)
	err := masterdata_db.Table("m_accessory_level_up_item").Find(&gamedata.AccessoryLevelUpItem)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryLevelUpItem)
}
