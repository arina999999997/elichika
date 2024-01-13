// Implement the array type
// <TYPE>__Array
// This type is usually used for list that won't be too long in the client(?)
// The implementation store 32 pointers, maybe they're used for dynamic growth?
// Not too sure, here we just use the same implementation and interface as a list

package generic

import (
	"encoding/json"
)

type Array[T any] struct {
	Slice []T
}

// this is not cached so avoid calling it in loop
// just how go work I suppose
func (l *Array[T]) Size() int {
	return len(l.Slice)
}
func (l *Array[T]) Append(item T) {
	l.Slice = append(l.Slice, item)
}

func (l *Array[T]) Copy() Array[T] {
	var res Array[T]
	res.Slice = append(res.Slice, l.Slice...)
	return res
}

func (l *Array[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	return json.Unmarshal(data, &l.Slice)
}

func (l Array[T]) MarshalJSON() ([]byte, error) {
	if l.Slice == nil {
		return []byte("[]"), nil
	}
	bytes, err := json.Marshal(&l.Slice)
	return bytes, err
}
