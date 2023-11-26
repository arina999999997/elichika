package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type StoryMember struct {
	// from m_story_member
	ID int `xorm:"pk 'id'"`
	// MemberMasterID int `xorm:"member_master_id"`
	// StoryNo int
	// LoveLevel int
	// Title string
	// Description string
	// ScenarioScriptAssetPath string
	// CardImageAssetPath string
	// DisplayOrder int
	UnlockLiveID *int `xorm:"'unlock_live_id'"`

	// from m_story_member_rewards
	Reward *model.Content `xorm:"-"`
}

func (story *StoryMember) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	reward := model.Content{}
	exist, err := masterdata_db.Table("m_story_member_rewards").Where("story_member_master_id = ?", story.ID).Get(&reward)
	utils.CheckErr(err)
	if exist {
		story.Reward = &reward
	}
}

func loadStoryMember(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading StoryMember")
	gamedata.StoryMember = make(map[int]*StoryMember)
	err := masterdata_db.Table("m_story_member").Find(&gamedata.StoryMember)
	utils.CheckErr(err)
	for _, storyMember := range gamedata.StoryMember {
		storyMember.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadStoryMember)
}
