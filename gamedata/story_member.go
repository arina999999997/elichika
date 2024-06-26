package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type StoryMember struct {
	// from m_story_member
	Id        int32 `xorm:"pk 'id'"`
	MemberMId int32 `xorm:"'member_m_id'"`
	// StoryNo int
	LoveLevel int32 `xorm:"'love_level'"`
	// Title string
	// Description string
	// ScenarioScriptAssetPath string
	// CardImageAssetPath string
	// DisplayOrder int
	UnlockLiveId *int32 `xorm:"'unlock_live_id'"`

	// from m_story_member_rewards
	Reward *client.Content `xorm:"-"`
}

func (story *StoryMember) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	reward := client.Content{}
	exist, err := masterdata_db.Table("m_story_member_rewards").Where("story_member_master_id = ?", story.Id).Get(&reward)
	utils.CheckErr(err)
	if exist {
		story.Reward = &reward
	}
}

func loadStoryMember(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading StoryMember")
	gamedata.StoryMember = make(map[int32]*StoryMember)
	err := masterdata_db.Table("m_story_member").Find(&gamedata.StoryMember)
	utils.CheckErr(err)
	for _, storyMember := range gamedata.StoryMember {
		storyMember.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadStoryMember)
}
