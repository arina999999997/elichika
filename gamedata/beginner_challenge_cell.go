package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type BeginnerChallengeCell struct {
	// from m_challenge_cell
	Id int32 `xorm:"pk 'id'"`
	// SetMId int32 `xorm:"'set_m_id'"`
	ChallengeId int32 `xorm:"-"`
	// Title string `xorm:"'title'"`
	// Summary string `xorm:"'summary'"`
	// SceneTransitionLink int32 `json:"scene_transition_link"`
	// SceneTransitionParam int32  `json:"scene_transition_param"`
	// DisplayOrder int32 `xorm:"'display_order'"`
	MissionClearConditionType   int32  `xorm:"'mission_clear_condition_type'" enum:"MissionClearConditionType"`
	MissionClearConditionCount  int32  `xorm:"'mission_clear_condition_count'"`
	MissionClearConditionParam1 *int32 `xorm:"'mission_clear_condition_param1'"`
	MissionClearConditionParam2 *int32 `xorm:"'mission_clear_condition_param2'"` // always null
	// StartCount 					int32  `xorm:"'start_count'"` // always 1

	// from m_challenge_reward
	Rewards []client.Content `xorm:"-"`
}

func (c *BeginnerChallengeCell) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := masterdata_db.Table("m_challenge_reward").Where("cell_m_id = ?", c.Id).OrderBy("display_order").
		Find(&c.Rewards)
	gamedata.BeginnerChallengeCellByClearConditionType[c.MissionClearConditionType] =
		append(gamedata.BeginnerChallengeCellByClearConditionType[c.MissionClearConditionType], c)
	utils.CheckErr(err)
}

func loadBeginnerChallengeCell(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading BeginnerChallengeCell")
	gamedata.BeginnerChallengeCell = make(map[int32]*BeginnerChallengeCell)
	gamedata.BeginnerChallengeCellByClearConditionType = make(map[int32][]*BeginnerChallengeCell)
	err := masterdata_db.Table("m_challenge_cell").Find(&gamedata.BeginnerChallengeCell)
	utils.CheckErr(err)
	for _, cell := range gamedata.BeginnerChallengeCell {
		cell.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadBeginnerChallengeCell)
}
