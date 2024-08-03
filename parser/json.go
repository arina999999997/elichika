package parser

import (
	"elichika/utils"

	"encoding/json"
)

func ParseJson(path string, result any) {
	text := utils.ReadAllText(path)
	err := json.Unmarshal([]byte(text), result)
	utils.CheckErr(err)
}
