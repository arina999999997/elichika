package client

import (
	"elichika/generic"
)

// yep this is the real name in the code
// Strucktur is structure in deutsch but it might as well be engrish
// TextureStruktur can be null, in which case V should be null, not the whole struct
type TextureStruktur struct {
	V generic.Nullable[string] `json:"v" xorm:"pk"`
}

// for xorm
func (ts *TextureStruktur) FromDB(data []byte) error {
	if string(data) == "" {
		return nil
	}
	ts.V = generic.NewNullable(string(data))
	return nil
}
func (ts *TextureStruktur) ToDB() ([]byte, error) {
	return []byte(ts.V.Value), nil
}
