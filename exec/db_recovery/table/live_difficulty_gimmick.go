package table

import (
	"elichika/exec/db_recovery/parser"
	"elichika/utils"

	"strconv"
)

type LiveDifficultyGimmick struct {
}

func (*LiveDifficultyGimmick) Table() string {
	return "m_live_difficulty_gimmick"
}
func (*LiveDifficultyGimmick) Id(fields []parser.Field) int64 {
	if fields[0].Key != "id" {
		panic("wrong field order")
	}
	id, err := strconv.ParseInt(fields[0].Value, 10, 64)
	utils.CheckErr(err)
	return id
}
func (*LiveDifficultyGimmick) Value(field parser.Field) string {
	return field.Value
}
func (ldg *LiveDifficultyGimmick) Update(field parser.Field) string {
	return field.Key + "=" + ldg.Value(field)
}
func (ldg *LiveDifficultyGimmick) Condition(fields []parser.Field) string {
	return ldg.Update(fields[0])
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
