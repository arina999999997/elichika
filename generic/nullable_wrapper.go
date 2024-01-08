package generic

import (
	// "fmt"
	"encoding/json"
)

type Nullable[T any] struct {
	Value    T    `xorm:"-"`
	HasValue bool `xorm:"-"` // zero value mean HasValue = false and thus the zero value is empty
}

// Unmarshal: from JSON bytes to value
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.HasValue = false
		return nil
	}
	n.HasValue = true
	return json.Unmarshal(data, &n.Value)
}
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.HasValue {
		bytes, err := json.Marshal(n.Value)
		return bytes, err
	} else {
		return []byte("null"), nil
	}
}

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		Value:    value,
		HasValue: true,
	}
}

// if used in database, mark the column type as json
// for example:
// - Item Nullable[int] `xorm:"json 'item'"`
// there might be a better way to do this but this is good enough for now
