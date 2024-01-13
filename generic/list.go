// Implement the list types: List_1_<TYPE>
// This is to be used for il2cpp list, which apparently is a c# List<T>
// It is basically a slice but it will be jsonfied to an empty array [] instead of null
// This is better than a slice because for a lot of types, we would need to manually set the slice to an empty slice and not nil

package generic

import (
	"encoding/json"
)

type List[T any] struct {
	Slice []T
}

// this is not cached so avoid calling it in loop
// just how go work I suppose
func (l *List[T]) Size() int {
	return len(l.Slice)
}
func (l *List[T]) Append(item T) {
	l.Slice = append(l.Slice, item)
}

func (l *List[T]) Copy() List[T] {
	var res List[T]
	res.Slice = append(res.Slice, l.Slice...)
	return res
}

func (l *List[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	return json.Unmarshal(data, &l.Slice)
}

func (l List[T]) MarshalJSON() ([]byte, error) {
	if l.Slice == nil {
		return []byte("[]"), nil
	}
	bytes, err := json.Marshal(&l.Slice)
	return bytes, err
}
