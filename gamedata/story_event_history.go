package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type StoryEventHistory struct {
	// from m_story_event_history_detail
	StoryEventId int32 `xorm:"pk 'story_event_id'"`
	// ...
}

func loadStoryEventHistory(gamedata *Gamedata) {
	fmt.Println("Loading StoryEventHistory")
	gamedata.StoryEventHistory = make(map[int32]*StoryEventHistory)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_story_event_history_detail").Find(&gamedata.StoryEventHistory)
	})
	utils.CheckErr(err)
}

func init() {
	addLoadFunc(loadStoryEventHistory)
}
