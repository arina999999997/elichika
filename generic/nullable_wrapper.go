package generic

import (
	"encoding/json"
)

type Nullable[T any] struct {
	Object T    `xorm:"extends"`
	IsNull bool `xorm:"-"`
}

// Unmarshal: from JSON bytes to value
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.IsNull = true
		return nil
	}
	n.IsNull = false
	return json.Unmarshal(data, n.Object)
}
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.IsNull {
		return []byte("null"), nil
	}
	bytes, err := json.Marshal(n.Object)
	return bytes, err
}

func NewNullable[T any](obj T) Nullable[T] {
	return Nullable[T]{
		Object: obj,
		IsNull: false,
	}
}
