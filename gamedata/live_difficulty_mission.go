package gamedata

import (
	"elichika/client"
)

type LiveDifficultyMission struct {
	// from m_live_difficulty_mission
	// LiveDifficultyMasterId int
	Position   int32 `xorm:"'position'"`
	TargetType int32 `xorm:"'target_type'" enum:"LiveMissionType"`
	// TargetValue int32            `xorm:"'target_value'"` // this field is wrong for some songs, so we use the LiveDifficulty's value instead
	Reward client.Content `xorm:"extends"`
}
