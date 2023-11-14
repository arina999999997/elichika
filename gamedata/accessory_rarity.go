package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryLevelUp struct {
	PlusExp   int `xorm:"'plus_exp'"`
	GameMoney int `xorm:"'game_money'"`
}

const GRADE_COUNT = 6 // from 0 - 5. Not hardcoding this is pretty messy
type AccessoryRarity struct {
	// from m_accessory_rarity_setting
	RarityType int `xorm:"pk 'rarity_type'"`
	// Name string
	// from m_accessory_grade_up_setting
	GradeUpMoney []int `xorm:"-"` // 0 indexed
	// from m_accessory_level_up_setting
	LevelUp []AccessoryLevelUp `xorm:"-"` // 0 indexed
	// from m_accessory_passive_skill_level_up_denominator
	// 0 indexed on the grade access then 1 indexed on the skill level access
	SkillLevelUpDenominator [6]([]int) `xorm:"-"`
	// from m_accessory_passive_skill_level_up_plus_percent
	SkillLevelUpPlusPercent ([]int) `xorm:"-"`
	// from m_accessory_passive_skill_level_up_setting
	GradeMaxSkillLevel []int `xorm:"-"` // 0 indexed
	// from m_accessory_rarity_up_setting
	RarityUpMoney int `xorm:"-"`
}

func (rarity *AccessoryRarity) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	{
		err := masterdata_db.Table("m_accessory_grade_up_setting").Where("rarity = ?", rarity.RarityType).OrderBy("grade").
			Cols("game_money").Find(&rarity.GradeUpMoney)
		utils.CheckErr(err)
	}
	{
		err := masterdata_db.Table("m_accessory_level_up_setting").Where("rarity = ?", rarity.RarityType).OrderBy("level").Find(&rarity.LevelUp)
		utils.CheckErr(err)
		rarity.LevelUp = append([]AccessoryLevelUp{AccessoryLevelUp{}}, rarity.LevelUp...)
	}

	{
		for grade := 0; grade < GRADE_COUNT; grade++ {
			err := masterdata_db.Table("m_accessory_passive_skill_level_up_denominator").Where("rarity = ? AND grade = ?", rarity.RarityType, grade).
				OrderBy("skill_level").Cols("denominator").Find(&rarity.SkillLevelUpDenominator[grade])
			utils.CheckErr(err)
			rarity.SkillLevelUpDenominator[grade] = append([]int{0}, rarity.SkillLevelUpDenominator[grade]...)
		}
	}

	{
		err := masterdata_db.Table("m_accessory_passive_skill_level_up_plus_percent").Where("rarity = ?", rarity.RarityType).
			OrderBy("skill_level").Cols("plus_percent").Find(&rarity.SkillLevelUpPlusPercent)
		utils.CheckErr(err)
		rarity.SkillLevelUpPlusPercent = append([]int{0}, rarity.SkillLevelUpPlusPercent...)
	}

	{
		err := masterdata_db.Table("m_accessory_passive_skill_level_up_setting").Where("rarity = ?", rarity.RarityType).
			OrderBy("grade").Cols("max_level").Find(&rarity.GradeMaxSkillLevel)
		utils.CheckErr(err)
	}

	{
		_, err := masterdata_db.Table("m_accessory_rarity_up_setting").Where("rarity = ?", rarity.RarityType).Cols("game_money").Get(&rarity.RarityUpMoney)
		utils.CheckErr(err)
	}
}

func loadAccessoryRarity(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	gamedata.AccessoryRarity = make(map[int]*AccessoryRarity)
	err := masterdata_db.Table("m_accessory_rarity_setting").Find(&gamedata.AccessoryRarity)
	utils.CheckErr(err)
	for _, rarity := range gamedata.AccessoryRarity {
		rarity.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadAccessoryRarity)
}
