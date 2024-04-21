package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type StoryEventHistory struct {
	// from m_story_event_history_detail
	StoryEventId int32 `xorm:"pk 'story_event_id'"`
	// ...
}

func loadStoryEventHistory(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading StoryEventHistory")
	gamedata.StoryEventHistory = make(map[int32]*StoryEventHistory)
	err := masterdata_db.Table("m_story_event_history_detail").Find(&gamedata.StoryEventHistory)
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadStoryEventHistory)
}
