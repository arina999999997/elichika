package gamedata

import (
	"elichika/model"
)

type LiveDifficultyMission struct {
	// from m_live_difficulty_mission
	// LiveDifficultyMasterID int
	Position    int           `xorm:"'position'"`
	TargetType  int           `xorm:"'target_type'"`
	TargetValue int           `xorm:"'target_value'"`
	Reward      model.Content `xorm:"extends"`
}
