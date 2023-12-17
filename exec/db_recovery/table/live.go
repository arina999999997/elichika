package table

import (
	"elichika/exec/db_recovery/parser"
	"elichika/utils"

	"strconv"
)

type Live struct {
}

func (_ *Live) Table() string {
	return "m_live"
}
func (_ *Live) ID(fields []parser.Field) int64 {
	if fields[0].Key != "live_id" {
		panic("wrong field order")
	}
	id, err := strconv.ParseInt(fields[0].Value, 10, 64)
	utils.CheckErr(err)
	return id
}
func (_ *Live) Value(field parser.Field) string {
	if field.Key == "member_unit" || field.Key == "original_deck_name" || field.Key == "source" {
		if field.Value == "\"\"" {
			return "NULL"
		}
	}
	return field.Value
}
func (this *Live) Update(field parser.Field) string {
	return field.Key + "=" + this.Value(field)
}
func (this *Live) Condition(fields []parser.Field) string {
	return this.Update(fields[0])
}

func handleLiveEvent(event parser.ModifierEvent[Live]) {
	if event.Type == parser.DELETE { // if deleted then we can add it back
		event.Type = parser.INSERT
	} else if event.Type == parser.INSERT {
		return
	}
	output += event.String() + "\n"
}

func handleLive() {
	var dummy Live
	events := parser.Parse[Live](readGitChange(dummy.Table()))
	for _, event := range events {
		handleLiveEvent(event)
	}
}

func init() {
	addHandler(handleLive)
}
