package gamedata

import (
	"elichika/client"
)

type LiveDifficultyMission struct {
	// from m_live_difficulty_mission
	// LiveDifficultyMasterId int
	Position    int            `xorm:"'position'"`
	TargetType  int            `xorm:"'target_type'"`
	TargetValue int            `xorm:"'target_value'"`
	Reward      client.Content `xorm:"extends"`
}
