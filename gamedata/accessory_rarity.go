package gamedata

import (
	"elichika/utils"

	"xorm.io/xorm"
)

type AccessoryLevelUp struct {
	PlusExp   int32 `xorm:"'plus_exp'"`
	GameMoney int32 `xorm:"'game_money'"`
}

const GRADE_COUNT = 6 // from 0 - 5. Not hardcoding this is pretty messy
type AccessoryRarity struct {
	// from m_accessory_rarity_setting
	RarityType int32 `xorm:"pk 'rarity_type'" enum:"AccessoryRarity"`
	// Name string
	// from m_accessory_grade_up_setting
	GradeUpMoney []int32 `xorm:"-"` // 0 indexed
	// from m_accessory_level_up_setting
	LevelUp []AccessoryLevelUp `xorm:"-"` // 0 indexed
	// from m_accessory_passive_skill_level_up_denominator
	// 0 indexed on the grade access then 1 indexed on the skill level access
	SkillLevelUpDenominator [6]([]int32) `xorm:"-"`
	// from m_accessory_passive_skill_level_up_plus_percent
	SkillLevelUpPlusPercent ([]int32) `xorm:"-"`
	// from m_accessory_passive_skill_level_up_setting
	GradeMaxSkillLevel []int32 `xorm:"-"` // 0 indexed
	// from m_accessory_rarity_up_setting
	RarityUpMoney int32 `xorm:"-"`
}

func (rarity *AccessoryRarity) populate(gamedata *Gamedata) {
	var err error
	{
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_accessory_grade_up_setting").Where("rarity = ?", rarity.RarityType).OrderBy("grade").Cols("game_money").Find(&rarity.GradeUpMoney)
		})
		utils.CheckErr(err)
	}
	{
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_accessory_level_up_setting").Where("rarity = ?", rarity.RarityType).OrderBy("level").Find(&rarity.LevelUp)
		})
		utils.CheckErr(err)
		rarity.LevelUp = append([]AccessoryLevelUp{AccessoryLevelUp{}}, rarity.LevelUp...)
	}

	{
		for grade := 0; grade < GRADE_COUNT; grade++ {
			gamedata.MasterdataDb.Do(func(session *xorm.Session) {
				err = session.Table("m_accessory_passive_skill_level_up_denominator").Where("rarity = ? AND grade = ?", rarity.RarityType, grade).OrderBy("skill_level").Cols("denominator").Find(&rarity.SkillLevelUpDenominator[grade])
			})
			utils.CheckErr(err)
			rarity.SkillLevelUpDenominator[grade] = append([]int32{0}, rarity.SkillLevelUpDenominator[grade]...)
		}
	}

	{
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_accessory_passive_skill_level_up_plus_percent").Where("rarity = ?", rarity.RarityType).OrderBy("skill_level").Cols("plus_percent").Find(&rarity.SkillLevelUpPlusPercent)
		})
		utils.CheckErr(err)
		rarity.SkillLevelUpPlusPercent = append([]int32{0}, rarity.SkillLevelUpPlusPercent...)
	}

	{
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			err = session.Table("m_accessory_passive_skill_level_up_setting").Where("rarity = ?", rarity.RarityType).OrderBy("grade").Cols("max_level").Find(&rarity.GradeMaxSkillLevel)
		})
		utils.CheckErr(err)
	}

	{
		gamedata.MasterdataDb.Do(func(session *xorm.Session) {
			_, err = session.Table("m_accessory_rarity_up_setting").Where("rarity = ?", rarity.RarityType).Cols("game_money").Get(&rarity.RarityUpMoney)
		})
		utils.CheckErr(err)
	}
}

func loadAccessoryRarity(gamedata *Gamedata) {
	gamedata.AccessoryRarity = make(map[int32]*AccessoryRarity)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_accessory_rarity_setting").Find(&gamedata.AccessoryRarity)
	})
	utils.CheckErr(err)
	for _, rarity := range gamedata.AccessoryRarity {
		rarity.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadAccessoryRarity)
}
