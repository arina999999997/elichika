// implement the dictionary types
// Dictionary_2_<KEY_TYPE>_DotUnder_<VALUE_TYPE>
// requirement:
// - Convertible to and from jsons [key_1, value_1, key_2, value_2, key_3, value_3, ...]
// - The keys are sorted increasingly.
//   - This is the behaviour of the official server.
//   - The client can sometime handle things if it's not sorted, but other time it might cause problems
//   - for example the order of decks
//
// - If the map is nil or empty, return []
// Note that sometime the value can be null, if this is the case, it's up to the user to set the VALUE_TYPE to Nullable

// To read and write the map from database, tag the table and map key like so
// type UserModel struct {
//     ...
//     MapField generic.Dictionary[KeyType, ValueType]  `table:"table_name" key:"col_name"`
// }
// The system will see if col_name exists in ValueType, and if it doesn't exist then it will use a different interface.
// Reading is done by passing the db, the user id, the table name and the table key since those exist outside of the mapping types
// Writing is still done using finalizer because it has to handle update and insert and all that.

package generic

import (
	"elichika/utils"

	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"

	"xorm.io/xorm"
)

type Dictionary[K int32 | int64, V any] struct {
	Map map[K]V
}

func (d *Dictionary[K, V]) Get(key K) (V, bool) {
	value, exist := d.Map[key]
	return value, exist
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	if d.Map == nil {
		d.Map = make(map[K]V)
	}
	d.Map[key] = value
}

func (d *Dictionary[K, V]) SetZero(key K) {
	var value V
	d.Set(key, value)
}

func (d *Dictionary[K, V]) Has(key K) bool {
	_, exist := d.Map[key]
	return exist
}

func (d *Dictionary[K, V]) Remove(key K) {
	delete(d.Map, key)
}

func (d *Dictionary[K, V]) GetOnly(key K) V {
	return d.Map[key]
}

func (d *Dictionary[K, V]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	arr := []json.RawMessage{}
	err := json.Unmarshal(data, &arr) // first unmarshal into an array of raw json
	if err != nil {
		return err
	}
	n := len(arr)
	if n%2 != 0 {
		return errors.New("array isn't a dictionary")
	}
	for i := 0; i < n; i += 2 {
		var key K
		var value V
		err = json.Unmarshal(arr[i], &key)
		if err != nil {
			return err
		}
		err = json.Unmarshal(arr[i+1], &value)
		if err != nil {
			return err
		}
		if d.Has(key) {
			return errors.New("key already exists")
		}
		d.Set(key, value)
	}
	return nil
}

func (d Dictionary[K, V]) MarshalJSON() ([]byte, error) {
	arr := []any{}
	keys := []K{}
	for key := range d.Map {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, key := range keys {
		arr = append(arr, key)
		arr = append(arr, d.Map[key])
	}
	bytes, err := json.Marshal(&arr)
	return bytes, err
}

// Load this from a xorm session
func (d *Dictionary[K, V]) LoadFromDb(db *xorm.Session, userId int, table, mapKey string) {
	// TODO(refactor) cache this
	var valueDummy V
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
	var values []V
	if valueHasKey {
		err := db.Table(table).Where("user_id = ?", userId).Find(&values)
		utils.CheckErr(err)
		for _, v := range values {
			d.Set(reflect.ValueOf(v).Field(keyField).Interface().(K), v)
		}
	} else {
		// V is likely some sort of wrapper, we will use V's interface to generate the value and key arrays
		// use reflect because this might not exist
		var keys []any
		rLoadFromDB := reflect.ValueOf(valueDummy).MethodByName("LoadFromDb")
		if rLoadFromDB.IsValid() {
			rLoadFromDB.Call([]reflect.Value{reflect.ValueOf(db), reflect.ValueOf(userId),
				reflect.ValueOf(table), reflect.ValueOf(mapKey), reflect.ValueOf(&keys), reflect.ValueOf(&values)})
			for i := range values {
				d.Set(keys[i].(K), values[i])
			}
		} else {
			panic(fmt.Sprint("Not supported yet, table: ", table, ", key: ", mapKey))
		}
	}
}
