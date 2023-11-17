package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type LiveDifficulty struct {
	// from m_live_difficulty
	LiveDifficultyID int   `xorm:"pk 'live_difficulty_id'"`
	LiveID           *int  `xorm:"'live_id'"`
	Live             *Live `xorm:"-"`
	// Live3DAssetMasterID *int
	LiveDifficultyType int `xorm:"'live_difficulty_type'"`
	UnlockPattern      int `xorm:"'unlock_pattern'"`
	// DefaultAttribute int
	TargetVoltage int `xorm:"'target_voltage'"`
	ConsumedLP    int `xorm:"'consumed_lp'"`
	RewardUserExp int `xorm:"'reward_user_exp'"`
	// JudgeID int
	NoteDropGroupID *int `xorm:"'note_drop_group_id'"`
	// NoteDropGroup *NoteDropGroup `xorm:"-"`
	DropChooseCount    int  `xorm:"'drop_choose_count'"`
	RateDropRate       int  `xorm:"'rare_drop_rate'"`
	DropContentGroupID *int `xorm:"'drop_content_group_id'"`
	// DropContentGroup *DropContentGroup `xorm:"-"`
	RareDropContentGroupID *int `xorm:"'rare_drop_content_group_id'"`
	// RareDropContentGroup *RareDropContentGroup `xorm:"-"`
	AdditionalDropContentGroupID *int `xorm:"'additional_drop_content_group_id'"`
	// AdditionalDropContentGroup *AdditionalDropContentGroup `xorm:"-"`
	// ?????
	BottomTechnique              int `xorm:"'bottom_technique'"`
	AdditionalDropDecayTechnique int `xorm:"'additional_drop_decay_technique'"`

	RewardBaseLovePoint int `xorm:"'reward_base_love_point'"`
	EvaluationSScore    int `xorm:"'evaluation_s_score'"`
	EvaluationAScore    int `xorm:"'evaluation_a_score'"`
	EvaluationBScore    int `xorm:"'evaluation_b_score'"`
	EvaluationCScore    int `xorm:"'evaluation_c_score'"`
	// UpdatedAt int `xorm:"'updated_at'"`
	LoseAtDeath bool `xorm:"'lose_at_death'"`
	// AutoplayRequirementID *int `xorm:"'autoplay_requirement_id'"`
	SkipMasterID *int `xorm:"'skip_master_id'"`
	// StaminaVoltageGroupID int
	// ComboVoltageGroupID int
	// DifficultyConstMasterID int
	IsCountTarget bool `xorm:"'is_count_target'"`
	// InsufficentRate int

	// from m_live_difficulty_mission
	Missions []LiveDifficultyMission `xorm:"-"`
}

func (liveDifficulty *LiveDifficulty) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {

	liveDifficulty.Live = gamedata.Live[*liveDifficulty.LiveID]
	// 2-way links
	liveDifficulty.Live.LiveDifficulties = append(liveDifficulty.Live.LiveDifficulties, liveDifficulty)
	liveDifficulty.LiveID = &gamedata.Live[*liveDifficulty.LiveID].LiveID
	err := masterdata_db.Table("m_live_difficulty_mission").Where("live_difficulty_master_id = ?", liveDifficulty.LiveDifficultyID).
		OrderBy("position").Find(&gamedata.LiveDifficulty[liveDifficulty.LiveDifficultyID].Missions)
	utils.CheckErr(err)
}

func loadLiveDifficulty(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading LiveDifficulty")
	gamedata.LiveDifficulty = make(map[int]*LiveDifficulty)
	err := masterdata_db.Table("m_live_difficulty").Find(&gamedata.LiveDifficulty)
	utils.CheckErr(err)
	for _, liveDifficulty := range gamedata.LiveDifficulty {
		liveDifficulty.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadLiveDifficulty)
	addPrequisite(loadLiveDifficulty, loadLive)
}
