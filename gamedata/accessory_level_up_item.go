package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryLevelUpItem struct {
	// frm m_accessory_level_up_item
	Id int32 `xorm:"pk 'id'"`
	// Name string
	// ThumbnailAssetPath
	Rarity    int32 `xorm:"'rarity'"`
	Attribute int32 `xorm:"'attribute'"`
	PlusExp   int32 `xorm:"'plus_exp'"`
	GameMoney int32 `xorm:"'game_money'"`
	// Description string
	ItemListCategoryType int32 `xorm:"'item_list_category_type'" enum:"ItemListCategoryType"`
	// SceneId int
	// BannerImageAssetPath
	// DisplayOrder int
}

func loadAccessoryLevelUpItem(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryLevelUpItem = make(map[int32]*AccessoryLevelUpItem)
	err := masterdata_db.Table("m_accessory_level_up_item").Find(&gamedata.AccessoryLevelUpItem)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadAccessoryLevelUpItem)
}
