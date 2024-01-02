package table

import (
	"elichika/exec/db_recovery/parser"
	"elichika/utils"

	"strconv"
)

type LiveDifficultyGimmick struct {
}

func (_ *LiveDifficultyGimmick) Table() string {
	return "m_live_difficulty_gimmick"
}
func (_ *LiveDifficultyGimmick) ID(fields []parser.Field) int64 {
	if fields[0].Key != "id" {
		panic("wrong field order")
	}
	id, err := strconv.ParseInt(fields[0].Value, 10, 64)
	utils.CheckErr(err)
	return id
}
func (_ *LiveDifficultyGimmick) Value(field parser.Field) string {
	return field.Value
}
func (this *LiveDifficultyGimmick) Update(field parser.Field) string {
	return field.Key + "=" + this.Value(field)
}
func (this *LiveDifficultyGimmick) Condition(fields []parser.Field) string {
	return this.Update(fields[0])
}

func handleLiveDifficultyGimmickEvent(event parser.ModifierEvent[LiveDifficultyGimmick]) {
	if event.Type == parser.DELETE { // if deleted then we can add it back
		event.Type = parser.INSERT
	} else if event.Type == parser.INSERT {
		return
	}
	output += event.String() + "\n"
}

func handleLiveDifficultyGimmick() {
	var dummy LiveDifficultyGimmick
	events := parser.Parse[LiveDifficultyGimmick](readGitChange(dummy.Table()))
	for _, event := range events {
		handleLiveDifficultyGimmickEvent(event)
	}
}

func init() {
	addHandler(handleLiveDifficultyGimmick)
	addPrequisite(handleLiveDifficultyGimmick, handleLiveDifficulty)
}
