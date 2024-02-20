package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/generic/drop"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LessonMenu struct {
	// from m_lesson_menu
	Id int32 `xorm:"pk"`
	// PassiveSkillDropGroupId int32
	// Name string
	// ThumbnailMAssetPath string
	// ThumbnailSAssetPath string
	// BackgroundImagePath string
	// BgmPath string
	DefaultDrop *drop.WeightedDropList[client.LessonDropItem]           `xorm:"-"`
	Drop        map[int32]*drop.WeightedDropList[client.LessonDropItem] `xorm:"-"`
}

func (lm *LessonMenu) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {

	type LessonDropContent struct {
		ContentType   int32
		ContentId     int32
		ContentAmount int32
		Weight        int32
		Rarity        int32
	}

	contents := []LessonDropContent{}
	err := masterdata_db.Table("m_lesson_drop_content").Where("lesson_menu_master_id == ?", lm.Id).Find(&contents)
	utils.CheckErr(err)
	lm.DefaultDrop = &drop.WeightedDropList[client.LessonDropItem]{}
	for _, content := range contents {
		lm.DefaultDrop.AddItem(client.LessonDropItem{
			ContentType:   content.ContentType,
			ContentId:     content.ContentId,
			ContentAmount: content.ContentAmount,
			DropRarity:    content.Rarity,
		}, content.Weight)
	}

	type LessonEnhancingItemDropRate struct {
		LessonEnhancingItemId int32
		TargetRarity          int32
		MagnificationWeight   int32
	}
	enhancingItems := []LessonEnhancingItemDropRate{}
	err = masterdata_db.Table("m_lesson_enhancing_item_effect_drop_rate").Find(&enhancingItems)
	utils.CheckErr(err)
	lm.Drop = map[int32]*drop.WeightedDropList[client.LessonDropItem]{}
	for _, rate := range enhancingItems {
		if lm.Drop[rate.LessonEnhancingItemId] == nil {
			lm.Drop[rate.LessonEnhancingItemId] = &drop.WeightedDropList[client.LessonDropItem]{}
		}
		for _, content := range contents {
			if content.Rarity == rate.TargetRarity {
				lm.Drop[rate.LessonEnhancingItemId].AddItem(client.LessonDropItem{
					ContentType:   content.ContentType,
					ContentId:     content.ContentId,
					ContentAmount: content.ContentAmount,
					DropRarity:    content.Rarity,
				}, content.Weight*rate.MagnificationWeight/10000)
			}
		}
	}
}

func loadLessonMenu(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LessonMenu")
	gamedata.LessonMenu = make(map[int32]*LessonMenu)
	err := masterdata_db.Table("m_lesson_menu").Find(&gamedata.LessonMenu)
	utils.CheckErr(err)
	for _, lessonMenu := range gamedata.LessonMenu {
		lessonMenu.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadLessonMenu)
}
