package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type StoryMainChapter struct {
	// from m_story_main_chapter
	Id int32 `xorm:"pk 'id'"`
	// StoryMainPartMasterId int32 `xorm:"'story_main_part_master_id'"`
	// Title string `xorm:"'title'"`
	// Description string `xorm:"'description'"`
	// ThumbnailAssetPath string `xorm:"'thumbnail_asset_path'"`
	// BackgroundAssetPath string `xorm:"'background_asset_path'"`
	// HardBackgroundAssetPath string `xorm:"'hard_background_asset_path'"`
	// BgmAssetPath string `xorm:"'bgm_asset_path'"`

	// from m_story_main_cell
	LastCellId int32 `xorm:"-"`
}

func (smc *StoryMainChapter) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	cells := []int32{}
	err := masterdata_db.Table("m_story_main_cell").Where("chapter_id = ?", smc.Id).Cols("id").Find(&cells)
	utils.CheckErr(err)
	for _, cellId := range cells {
		if smc.LastCellId < cellId {
			smc.LastCellId = cellId
		}
	}
}

func loadStoryMain(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading StoryMainChapter")
	gamedata.StoryMainChapter = make(map[int32]*StoryMainChapter)
	err := masterdata_db.Table("m_story_main_chapter").Find(&gamedata.StoryMainChapter)
	utils.CheckErr(err)
	for _, storyMainChapter := range gamedata.StoryMainChapter {
		storyMainChapter.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadStoryMain)
}
