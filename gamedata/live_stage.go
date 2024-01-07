package gamedata

import (
	"elichika/model"
	"elichika/utils"

	"strconv"
)

type SimpleLiveStage struct {
	LiveDifficultyId int                     `json:"live_difficulty_id"`
	LiveNotes        []model.LiveNote        `json:"live_notes"`
	LiveWaveSettings []model.LiveWaveSetting `json:"live_wave_settings"`
	Original         *int                    `json:"original"`
}

type LiveDifficultyGimmick struct {
	Id                     int `xorm:"pk 'id'"`
	LiveDifficultyMasterId int `xorm:"live_difficulty_master_id"`
	TriggerType            int `xorm:"'trigger_type'"`
	ConditionMasterId1     int `xorm:"'condition_master_id1'"`
	ConditionMasterId2     int `xorm:"'condition_master_id2'"`
	SkillMasterId          int `xorm:"'skill_master_id'"`
}

type LiveDifficultyNoteGimmick struct {
	LiveDifficultyId    int    `xorm:"pk 'live_difficulty_id'"`
	NoteId              int    `xorm:"pk 'note_id'"`
	NoteGimmickType     int    `xorm:"note_gimmick_type"`
	NoteGimmickIconType int    `xorm:"note_gimmick_icon_type"`
	SkillMasterId       int    `xorm:"skill_master_id"`
	Name                string `xorm:"name"`
	Id                  int    `xorm:"-"` // must be extracted from name
}

type LiveNoteWaveGimmickGroup struct {
	LiveDifficultyId int `xorm:"pk 'live_difficulty_id'"`
	WaveId           int `xorm:"pk 'wave_id'"`
	State            int
	SkillId          int
}

func (ldng *LiveDifficultyNoteGimmick) populate() {
	var err error
	ldng.Id, err = strconv.Atoi(ldng.Name[len(ldng.Name)-8:])
	utils.CheckErr(err)
}
