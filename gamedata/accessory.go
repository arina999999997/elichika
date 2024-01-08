package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryGrade struct {
	// Grade int
	MaxLevel              int32  `xorm:"'max_level'"`
	PassiveSkill1MasterId *int32 `xorm:"'accessory_passive_skill_1_master_id'"`
	PassiveSkill2MasterId *int32 `xorm:"'accessory_passive_skill_2_master_id'"` // always null
}

type AccessoryRarityUp struct {
	AfterAccessoryMasterId *int32                  `xorm:"'after_accessory_master_id'"`
	AfterAccessory         *Accessory              `xorm:"-"`
	RarityUpGroupMasterId  *int32                  `xorm:"'accessory_rarity_up_group_master_id'"`
	RarityUpGroup          *AccessoryRarityUpGroup `xorm:"-"`
}

func (rarityUp *AccessoryRarityUp) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	rarityUp.AfterAccessory = gamedata.Accessory[*rarityUp.AfterAccessoryMasterId]
	rarityUp.RarityUpGroup = gamedata.AccessoryRarityUpGroup[*rarityUp.RarityUpGroupMasterId]
}

type Accessory struct {
	// from m_accessory
	Id int32 `xorm:"pk 'id'"`
	// Name string
	// AccessoryNo int
	// ThumbnailAssetPath string
	// AccessoryType int // always 1
	// MemberMasterId *int // always null
	RarityType int32            `xorm:"'rarity_type'" enum:"AccessoryRarity"`
	Rarity     *AccessoryRarity `xorm:"-"`
	Attribute  int32            `xorm:"'attribute'" enum:"CardAttribute"`
	Role       int              `xorm:"'role'"`
	MaxGrade   int              `xorm:"'max_grade'"`

	// from m_accessory_grade_up
	Grade []AccessoryGrade `xorm:"-"` // 0 indexed
	// from m_accessory_level_up
	LevelExp []int32 `xorm:"-"` // 1 indexed, total amount of exp

	// from m_accessory_melt
	MeltGroupMasterIds []int32               `xorm:"-"` // 0 indexed
	MeltGroup          []*AccessoryMeltGroup `xorm:"-"` // 0 indexed

	// from m_accessory_rarity_up
	RarityUp *AccessoryRarityUp `xorm:"-"`
}

func (accessory *Accessory) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	accessory.Rarity = gamedata.AccessoryRarity[accessory.RarityType]

	{
		err := masterdata_db.Table("m_accessory_grade_up").Where("accessory_master_id = ?", accessory.Id).OrderBy("grade").Find(&accessory.Grade)
		utils.CheckErr(err)
	}

	{
		err := masterdata_db.Table("m_accessory_level_up").Where("accessory_master_id = ?", accessory.Id).OrderBy("level").Cols("exp").Find(&accessory.LevelExp)
		utils.CheckErr(err)
		accessory.LevelExp = append([]int32{0}, accessory.LevelExp...)
	}

	{
		err := masterdata_db.Table("m_accessory_melt").Where("accessory_master_id = ?", accessory.Id).OrderBy("grade").
			Cols("accessory_melt_group_master_id").Find(&accessory.MeltGroupMasterIds)
		utils.CheckErr(err)
		for _, meltGroupId := range accessory.MeltGroupMasterIds {
			accessory.MeltGroup = append(accessory.MeltGroup, gamedata.AccessoryMeltGroup[meltGroupId])
		}
	}

	{
		rarityUp := AccessoryRarityUp{}
		exist, err := masterdata_db.Table("m_accessory_rarity_up").Where("accessory_master_id = ?", accessory.Id).Get(&rarityUp)
		utils.CheckErr(err)
		if exist {

			rarityUp.populate(gamedata, masterdata_db, serverdata_db, dictionary)
			accessory.RarityUp = &rarityUp
		}
	}
}

func loadAccessory(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.Accessory = make(map[int32]*Accessory)
	err := masterdata_db.Table("m_accessory").Find(&gamedata.Accessory)
	utils.CheckErr(err)
	for _, accessory := range gamedata.Accessory {
		accessory.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}

}

func init() {
	addLoadFunc(loadAccessory)
	addPrequisite(loadAccessory, loadAccessoryRarity)
	addPrequisite(loadAccessory, loadAccessoryMeltGroup)
	addPrequisite(loadAccessory, loadAccessoryRarityUpGroup)
}
