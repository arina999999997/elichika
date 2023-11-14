package gamedata

import (
	"elichika/utils"
	"elichika/dictionary"

	"xorm.io/xorm"
)

type AccessoryGrade struct {
	// Grade int
	MaxLevel                       int  `xorm:"'max_level'"`
	PassiveSkill1MasterID *int `xorm:"'accessory_passive_skill_1_master_id'"`
	PassiveSkill2MasterID *int `xorm:"'accessory_passive_skill_2_master_id'"` // always null
}

type AccessoryRarityUp struct {
	AfterAccessoryMasterID *int                    `xorm:"'after_accessory_master_id'"`
	AfterAccessory         *Accessory              `xorm:"-"`
	RarityUpGroupMasterID  *int                    `xorm:"'accessory_rarity_up_group_master_id'"`
	RarityUpGroup          *AccessoryRarityUpGroup `xorm:"-"`
}

func (rarityUp *AccessoryRarityUp) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	rarityUp.AfterAccessory = gamedata.Accessory[*rarityUp.AfterAccessoryMasterID]
	rarityUp.RarityUpGroup = gamedata.AccessoryRarityUpGroup[*rarityUp.RarityUpGroupMasterID]
}

type Accessory struct {
	// from m_accessory
	ID int `xorm:"pk 'id'"`
	// Name string
	// AccessoryNo int
	// ThumbnailAssetPath string
	// AccessoryType int // always 1
	// MemberMasterID *int // always null
	RarityType int              `xorm:"'rarity_type'"`
	Rarity     *AccessoryRarity `xorm:"-"`
	Attribute  int              `xorm:"'attribute'"`
	Role       int              `xorm:"'role'"`
	MaxGrade   int              `xorm:"'max_grade'"`

	// from m_accessory_grade_up
	Grade []AccessoryGrade `xorm:"-"` // 0 indexed
	// from m_accessory_level_up
	LevelExp []int `xorm:"-"` // 1 indexed, total amount of exp

	// from m_accessory_melt
	MeltGroupMasterIDs []int                 `xorm:"-"` // 0 indexed
	MeltGroup          []*AccessoryMeltGroup `xorm:"-"` // 0 indexed

	// from m_accessory_rarity_up
	RarityUp *AccessoryRarityUp `xorm:"-"`
}

func (accessory *Accessory) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	accessory.Rarity = gamedata.AccessoryRarity[accessory.RarityType]

	{
		err := masterdata_db.Table("m_accessory_grade_up").Where("accessory_master_id = ?", accessory.ID).OrderBy("grade").Find(&accessory.Grade)
		utils.CheckErr(err)
	}

	{
		err := masterdata_db.Table("m_accessory_level_up").Where("accessory_master_id = ?", accessory.ID).OrderBy("level").Cols("exp").Find(&accessory.LevelExp)
		utils.CheckErr(err)
		accessory.LevelExp = append([]int{0}, accessory.LevelExp...)
	}

	{
		err := masterdata_db.Table("m_accessory_melt").Where("accessory_master_id = ?", accessory.ID).OrderBy("grade").
		Cols("accessory_melt_group_master_id").Find(&accessory.MeltGroupMasterIDs)
		utils.CheckErr(err)
		for _, meltGroupID := range accessory.MeltGroupMasterIDs {
			accessory.MeltGroup = append(accessory.MeltGroup, gamedata.AccessoryMeltGroup[meltGroupID])
		}
	}

	{
		rarityUp := AccessoryRarityUp{}
		exists, err := masterdata_db.Table("m_accessory_rarity_up").Where("accessory_master_id = ?", accessory.ID).Get(&rarityUp)
		utils.CheckErr(err)
		if exists {
			
			rarityUp.populate(gamedata, masterdata_db, serverdata_db, dictionary)
			accessory.RarityUp =&rarityUp
		}
	}
}

func loadAccessory(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.Accessory = make(map[int]*Accessory)
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
