package generic

import (
	"elichika/utils"

	"encoding/json"
	"fmt"
	"reflect"

	"xorm.io/xorm"
)

type Nullable[T any] struct {
	Value    T    `xorm:"-"`
	HasValue bool `xorm:"-"` // zero value mean HasValue = false and thus the zero value is empty
}

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		Value:    value,
		HasValue: true,
	}
}

func NewNullableFromPointer[T any](pointer *T) Nullable[T] {
	if pointer == nil {
		return Nullable[T]{
			HasValue: false,
		}
	} else {
		return Nullable[T]{
			Value:    *pointer,
			HasValue: true,
		}
	}
}

func (n *Nullable[T]) ToPointer() *T {
	if n.HasValue {
		return &n.Value
	} else {
		return nil
	}
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

// if used in database, mark the column type as json
// for example:
// - Item Nullable[int] `xorm:"json 'item'"`
// there might be a better way to do this but this is good enough for now

// for dictionary of Nullable only, load the Nullable from a table
func (Nullable[T]) LoadFromDb(db *xorm.Session, userId int32, table, mapKey string, keyResult *[]any, result *[]Nullable[T]) {

	var valueDummy T
	rValueType := reflect.TypeOf(valueDummy)
	valueHasKey := false
	var keyField int
	for i := 0; i < rValueType.NumField(); i++ {
		f := rValueType.Field(i)
		if mapKey == f.Tag.Get("json") {
			valueHasKey = true
			keyField = i
			break
		}
	}
	if !valueHasKey {
		panic(fmt.Sprint("Not supported yet, table: ", table, ", key: ", mapKey))
	}
	var items []T
	err := db.Table(table).Where("user_id = ?", userId).Find(&items)
	utils.CheckErr(err)
	for _, item := range items {
		*keyResult = append(*keyResult, reflect.ValueOf(item).Field(keyField).Interface())
		*result = append(*result, NewNullable(item))
	}
}
