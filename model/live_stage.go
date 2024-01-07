package model

import (
	"elichika/utils"

	"encoding/json"
	"fmt"
)

type LiveStage struct {
	LiveDifficultyId int               `json:"live_difficulty_id"`
	LiveNotes        []LiveNote        `json:"live_notes"`
	LiveWaveSettings []LiveWaveSetting `json:"live_wave_settings"`
	NoteGimmicks     []NoteGimmick     `json:"note_gimmicks"`
	StageGimmickDict []any             `json:"stage_gimmick_dict"`
}

func (ls *LiveStage) Copy() LiveStage {
	result := LiveStage{
		LiveDifficultyId: ls.LiveDifficultyId,
		LiveNotes:        []LiveNote{},
		LiveWaveSettings: []LiveWaveSetting{},
		NoteGimmicks:     []NoteGimmick{},
		StageGimmickDict: []any{},
	}
	result.LiveNotes = append(result.LiveNotes, ls.LiveNotes...)
	result.LiveWaveSettings = append(result.LiveWaveSettings, ls.LiveWaveSettings...)
	result.NoteGimmicks = append(result.NoteGimmicks, ls.NoteGimmicks...)
	result.StageGimmickDict = append(result.StageGimmickDict, ls.StageGimmickDict...)
	return result
}

func (ls *LiveStage) IsSame(other *LiveStage) bool {
	if ls.LiveDifficultyId != other.LiveDifficultyId {
		return false
	}
	if len(ls.LiveNotes) != len(other.LiveNotes) {
		return false
	}
	for i := range ls.LiveNotes {
		if !ls.LiveNotes[i].IsSame(&other.LiveNotes[i]) {
			fmt.Println(ls.LiveNotes[i])
			fmt.Println(other.LiveNotes[i])
			return false
		}
	}
	// fmt.Println("Notes OK")
	if len(ls.LiveWaveSettings) != len(other.LiveWaveSettings) {
		return false
	}
	for i := range ls.LiveWaveSettings {
		if ls.LiveWaveSettings[i] != other.LiveWaveSettings[i] {
			fmt.Println(ls.LiveWaveSettings[i])
			fmt.Println(other.LiveWaveSettings[i])
			return false
		}
	}
	// fmt.Println("Waves OK")
	if len(ls.NoteGimmicks) != len(other.NoteGimmicks) {
		return false
	}
	for i := range ls.NoteGimmicks {
		if !ls.NoteGimmicks[i].IsSame(&other.NoteGimmicks[i]) {
			return false
		}
	}
	// fmt.Println("Note Gimmicks OK")
	if len(ls.StageGimmickDict) != len(other.StageGimmickDict) {
		return false
	}
	if len(ls.StageGimmickDict) > 0 {
		lsDict, err := json.Marshal(ls.StageGimmickDict[0])
		utils.CheckErr(err)
		lsId := 0
		err = json.Unmarshal(lsDict, &lsId)
		utils.CheckErr(err)
		if lsId != other.StageGimmickDict[0].(int) {
			return false
		}

		lsDict, err = json.Marshal(ls.StageGimmickDict[1].([]any)[0])
		utils.CheckErr(err)
		lsGimmick := StageGimmick{}
		err = json.Unmarshal(lsDict, &lsGimmick)
		utils.CheckErr(err)
		otherGimmick := other.StageGimmickDict[1].([]StageGimmick)[0]
		if lsGimmick != otherGimmick {
			fmt.Println(lsGimmick)
			fmt.Println(otherGimmick)
			return false
		}
	}
	// fmt.Println("Stage Gimmick OK")
	return true
}

type LiveNote struct {
	Id                  int `json:"id"`
	CallTime            int `json:"call_time"`
	NoteType            int `json:"note_type"`
	NotePosition        int `json:"note_position"`
	GimmickId           int `json:"gimmick_id"`
	NoteAction          int `json:"note_action"`
	WaveId              int `json:"wave_id"`
	NoteRandomDropColor int `json:"note_random_drop_color"`
	AutoJudgeType       int `json:"auto_judge_type"`
}

func (ln *LiveNote) IsSame(other *LiveNote) bool {
	same := true
	same = same && (ln.Id == other.Id)
	same = same && (ln.CallTime == other.CallTime)
	same = same && (ln.NoteType == other.NoteType)
	same = same && (ln.NotePosition == other.NotePosition)
	same = same && (ln.GimmickId == other.GimmickId)
	same = same && (ln.NoteAction == other.NoteAction)
	same = same && (ln.WaveId == other.WaveId)
	return same
}

type LiveWaveSetting struct {
	Id            int `json:"id"`
	WaveDamage    int `json:"wave_damage"`
	MissionType   int `json:"mission_type"`
	Arg1          int `json:"arg_1"`
	Arg2          int `json:"arg_2"`
	RewardVoltage int `json:"reward_voltage"`
}

type NoteGimmick struct {
	UniqId          int `json:"uniq_id"`
	Id              int `json:"id"`
	NoteGimmickType int `json:"note_gimmick_type"`
	Arg1            int `json:"arg_1"`
	Arg2            int `json:"arg_2"`
	EffectMId       int `json:"effect_m_id"`
	IconType        int `json:"icon_type"`
}

func (ng *NoteGimmick) IsSame(other *NoteGimmick) bool {
	same := true
	same = same && (ng.UniqId == other.UniqId)
	same = same && (ng.NoteGimmickType == other.NoteGimmickType)
	same = same && (ng.Arg1 == other.Arg1)
	same = same && (ng.Arg2 == other.Arg2)
	same = same && (ng.EffectMId == other.EffectMId)
	if !same {
		return false
	}
	if ng.IconType == other.IconType {
		return true
	}
	if ng.IconType == 5 && other.IconType == 25 { // there was a db update that change this
		return true
	}
	if ng.IconType == 8 && other.IconType == 9 { // there was a db update that change this
		return true
	}
	return false
}

type StageGimmick struct {
	GimmickMasterId    int `json:"gimmick_master_id"`
	ConditionMasterId1 int `json:"condition_master_id_1"`
	ConditionMasterId2 int `json:"condition_master_id_2"`
	SkillMasterId      int `json:"skill_master_id"`
	UniqId             int `json:"uniq_id"`
}
