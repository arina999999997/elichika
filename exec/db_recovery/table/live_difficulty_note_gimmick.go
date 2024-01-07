package table

import (
	"elichika/exec/db_recovery/parser"
	"elichika/utils"

	"strconv"
)

type LiveDifficultyNoteGimmick struct {
}

func (*LiveDifficultyNoteGimmick) Table() string {
	return "m_live_difficulty_note_gimmick"
}
func (*LiveDifficultyNoteGimmick) Id(fields []parser.Field) int64 {
	if fields[0].Key != "live_difficulty_id" {
		panic("wrong field order")
	}
	if fields[1].Key != "note_id" {
		panic("wrong field order")
	}
	liveId, err := strconv.ParseInt(fields[0].Value, 10, 64)
	utils.CheckErr(err)
	noteId, err := strconv.ParseInt(fields[1].Value, 10, 64)
	utils.CheckErr(err)
	return liveId*1000 + noteId
}
func (*LiveDifficultyNoteGimmick) Value(field parser.Field) string {
	return field.Value
}
func (ldng *LiveDifficultyNoteGimmick) Update(field parser.Field) string {
	return field.Key + "=" + ldng.Value(field)
}
func (ldng *LiveDifficultyNoteGimmick) Condition(fields []parser.Field) string {
	return ldng.Update(fields[0])
}

func handleLiveDifficultyNoteGimmickEvent(event parser.ModifierEvent[LiveDifficultyNoteGimmick]) {
	var dummy LiveDifficultyNoteGimmick
	if event.Type == parser.DELETE {
		if recoveredLiveDifficulty[dummy.Id(event.Fields)/1000] { // only recover the notes for deleted map
			event.Type = parser.INSERT
		} else {
			return
		}
	} else if event.Type == parser.INSERT {
		return
	}
	output += event.String() + "\n"
}

func handleLiveDifficultyNoteGimmick() {
	var dummy LiveDifficultyNoteGimmick
	events := parser.Parse[LiveDifficultyNoteGimmick](readGitChange(dummy.Table()))
	for _, event := range events {
		handleLiveDifficultyNoteGimmickEvent(event)
	}
}

func init() {
	addHandler(handleLiveDifficultyNoteGimmick)
	addPrequisite(handleLiveDifficultyNoteGimmick, handleLiveDifficulty)
}
