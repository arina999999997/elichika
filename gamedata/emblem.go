package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Emblem struct {
	Id int32 `xorm:"pk 'id'"`
	// Name string `xorm:"'name'"`
	// Description string `xorm:"'description'"`
	// EmblemType int32 `xorm:"'emblem_type'"`
	Grade int32 `xorm:"'grade'"`
	// EmblemAssetPath string `xorm:"'emblem_asset_path'"`
	// EmblemSubAssetPath string `xorm:"'emblem_sub_asset_path'"`
	// EmblemClearConditionType int32 `xorm:"'emblem_clear_condition_type'"`
	// EmblemClearConditionParam int32 `xorm:"'emblem_clear_condition_param'"`
	// IsEmblemSecretCondition int32 `xorm:"'is_emblem_secret_condition'"`
	// IsEventEmblem int32 `xorm:"'is_event_emblem'"`
	// ReleasedAt int32 `xorm:"'released_at'"`
	// DisplayOrder int32 `xorm:"'display_order'"`
}

func loadEmblem(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Emblem")
	gamedata.Emblem = make(map[int32]*Emblem)
	err := masterdata_db.Table("m_emblem").Find(&gamedata.Emblem)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadEmblem)
}
