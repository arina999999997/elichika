package client

import (
	"elichika/generic"
)

type AdvScriptStruktur struct {
	V generic.Nullable[string] `json:"v" xorm:"pk"`
}

// for xorm
func (a *AdvScriptStruktur) FromDB(data []byte) error {
	if string(data) == "" {
		return nil
	}
	a.V = generic.NewNullable(string(data))
	return nil
}
func (a *AdvScriptStruktur) ToDB() ([]byte, error) {
	return []byte(a.V.Value), nil
}
