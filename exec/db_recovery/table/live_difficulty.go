package table

import (
	"elichika/exec/db_recovery/parser"
	"elichika/utils"

	"strconv"
)

var (
	recoveredLiveDifficulty map[int64]bool
)

type LiveDifficulty struct {
}

func (*LiveDifficulty) Table() string {
	return "m_live_difficulty"
}
func (*LiveDifficulty) Id(fields []parser.Field) int64 {
	if fields[0].Key != "live_difficulty_id" {
		panic("wrong field order")
	}
	id, err := strconv.ParseInt(fields[0].Value, 10, 64)
	utils.CheckErr(err)
	return id
}
func (*LiveDifficulty) Value(field parser.Field) string {
	if field.Key == "live_3d_asset_master_id" || field.Key == "autoplay_requirement_id" {
		if field.Value == "\"\"" {
			return "NULL"
		}
	}
	return field.Value
}
func (ld *LiveDifficulty) Update(field parser.Field) string {
	return field.Key + "=" + ld.Value(field)
}
func (ld *LiveDifficulty) Condition(fields []parser.Field) string {
	return ld.Update(fields[0])
}

func handleLiveDifficultyEvent(event parser.ModifierEvent[LiveDifficulty]) {
	var dummy LiveDifficulty
	if event.Type == parser.DELETE { // if deleted then we can add it back
		event.Type = parser.INSERT
		recoveredLiveDifficulty[dummy.Id(event.Fields)] = true
	} else if event.Type == parser.INSERT { // check if the unlock pattern is 3, then we set it to 1
		// return for now
		return
		// if event.Fields[4].Value == "3" {
		// 	event.Type = parser.UPDATE
		// 	event.Fields[4].Value = "1"
		// } else { // this already exists
		// 	return
		// }
	}
	output += event.String() + "\n"
}

func handleLiveDifficulty() {
	recoveredLiveDifficulty = make(map[int64]bool)
	var dummy LiveDifficulty
	events := parser.Parse[LiveDifficulty](readGitChange(dummy.Table()))
	for _, event := range events {
		handleLiveDifficultyEvent(event)
	}
}

func init() {
	addHandler(handleLiveDifficulty)
	addPrequisite(handleLiveDifficulty, handleLive)
	addPrequisite(handleLiveDifficulty, handleLiveDifficultyConst)
}
