package client

import (
	"elichika/generic"

	"fmt"
)

type LiveStage struct {
	LiveDifficultyId int32                                                      `xorm:"pk" json:"live_difficulty_id"`
	LiveNotes        generic.Array[LiveNoteSetting]                             `json:"live_notes"`
	LiveWaveSettings generic.Array[LiveWaveSetting]                             `json:"live_wave_settings"`
	NoteGimmicks     generic.Array[NoteGimmick]                                 `json:"note_gimmicks"`
	StageGimmickDict generic.Dictionary[int32, generic.Array[LiveStageGimmick]] `json:"stage_gimmick_dict"`
}

func (ls *LiveStage) IsSame(other *LiveStage) bool {
	if ls.LiveDifficultyId != other.LiveDifficultyId {
		return false
	}
	if ls.LiveNotes.Size() != other.LiveNotes.Size() {
		return false
	}
	for i := range ls.LiveNotes.Slice {
		if !ls.LiveNotes.Slice[i].IsSame(&other.LiveNotes.Slice[i]) {
			fmt.Println(ls.LiveNotes.Slice[i])
			fmt.Println(other.LiveNotes.Slice[i])
			return false
		}
	}

	if ls.LiveWaveSettings.Size() != other.LiveWaveSettings.Size() {
		return false
	}
	for i := range ls.LiveWaveSettings.Slice {
		if ls.LiveWaveSettings.Slice[i] != other.LiveWaveSettings.Slice[i] {
			fmt.Println(ls.LiveWaveSettings.Slice[i])
			fmt.Println(other.LiveWaveSettings.Slice[i])
			return false
		}
	}
	if ls.NoteGimmicks.Size() != other.NoteGimmicks.Size() {
		return false
	}
	for i := range ls.NoteGimmicks.Slice {
		if !ls.NoteGimmicks.Slice[i].IsSame(&other.NoteGimmicks.Slice[i]) {
			return false
		}
	}

	if ls.StageGimmickDict.Size() != other.StageGimmickDict.Size() {
		return false
	}
	if ls.StageGimmickDict.Size() > 0 {
		for _, key := range ls.StageGimmickDict.OrderedKey {
			thisArray := ls.StageGimmickDict.GetOnly(key)
			otherArray, exist := other.StageGimmickDict.Get(key)
			if !exist {
				return false
			}
			if thisArray.Size() != otherArray.Size() {
				return false
			}
			for i := range thisArray.Slice {
				if thisArray.Slice[i] != otherArray.Slice[i] {
					return false
				}
			}
		}
	}
	// fmt.Println("Stage Gimmick OK")
	return true
}

func (ls *LiveStage) Copy() LiveStage {
	return LiveStage{
		LiveDifficultyId: ls.LiveDifficultyId,
		LiveNotes:        ls.LiveNotes.Copy(),
		LiveWaveSettings: ls.LiveWaveSettings.Copy(),
		NoteGimmicks:     ls.NoteGimmicks.Copy(),
		StageGimmickDict: ls.StageGimmickDict.Copy(),
	}
}
