package client

// yep this is the real name in the code
// Strucktur is structure in deutsch but it might as well be engrish
type TextureStruktur struct {
	V string `json:"v" xorm:"pk"`
}

// for xorm
func (ts *TextureStruktur) FromDB(data []byte) error {
	ts.V = string(data)
	return nil
}
func (ts *TextureStruktur) ToDB() ([]byte, error) {
	return []byte(ts.V), nil
}
