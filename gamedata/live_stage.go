package gamedata

import (
	"elichika/client"
	"elichika/utils"

	"strconv"
)

type SimpleLiveStage struct {
	LiveDifficultyId int32                    `json:"live_difficulty_id"`
	LiveNotes        []client.LiveNoteSetting `json:"live_notes"`
	LiveWaveSettings []client.LiveWaveSetting `json:"live_wave_settings"`
	Original         *int32                   `json:"original"`
}

type LiveDifficultyGimmick struct {
	Id                     int32 `xorm:"pk 'id'"`
	LiveDifficultyMasterId int32 `xorm:"live_difficulty_master_id"`
	TriggerType            int32 `xorm:"'trigger_type'"`
	ConditionMasterId1     int32 `xorm:"'condition_master_id1'"`
	ConditionMasterId2     int32 `xorm:"'condition_master_id2'"`
	SkillMasterId          int32 `xorm:"'skill_master_id'"`
}

type LiveDifficultyNoteGimmick struct {
	LiveDifficultyId    int32  `xorm:"pk 'live_difficulty_id'"`
	NoteId              int32  `xorm:"pk 'note_id'"`
	NoteGimmickType     int32  `xorm:"note_gimmick_type"`
	NoteGimmickIconType int32  `xorm:"note_gimmick_icon_type"`
	SkillMasterId       int32  `xorm:"skill_master_id"`
	Name                string `xorm:"name"`
	Id                  int32  `xorm:"-"` // must be extracted from name
}

type LiveNoteWaveGimmickGroup struct {
	LiveDifficultyId int32 `xorm:"pk 'live_difficulty_id'"`
	WaveId           int32 `xorm:"pk 'wave_id'"`
	State            int32
	SkillId          int32
}

func (ldng *LiveDifficultyNoteGimmick) populate() {
	id, err := strconv.Atoi(ldng.Name[len(ldng.Name)-8:])
	utils.CheckErr(err)
	ldng.Id = int32(id)
}
