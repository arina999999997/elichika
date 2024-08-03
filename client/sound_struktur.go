package client

import (
	"elichika/generic"
)

type SoundStruktur struct {
	V generic.Nullable[string] `json:"v" xorm:"pk"`
}

// for xorm
func (ss *SoundStruktur) FromDB(data []byte) error {
	if string(data) == "" {
		return nil
	}
	ss.V = generic.NewNullable(string(data))
	return nil
}
func (ss *SoundStruktur) ToDB() ([]byte, error) {
	return []byte(ss.V.Value), nil
}
