// implement the dictionary types
// Dictionary_2_<KEY_TYPE>_DotUnder_<VALUE_TYPE>
// requirement:
// - Convertible to and from jsons [key_1, value_1, key_2, value_2, key_3, value_3, ...]
// - Ordered, in the sense that key_1, key_2, will be in the order they are first inserted
//   - Sometime this doesn't matter, sometime it doesn't.
// - If the map is nil or empty, return [] for json
// implementation choices:
// - The value are stored using pointer, this is consistent with the client
//   - The pointer can be null, in which case it should be jsonfied as null
// - the dictionary type is insert only
// - User can do unordered iteration using d.Map
// - Or ordered iteration using d.Order
// - User can access the inner pointer directly

// Pretty sure the implementation is Dictionary<TKey, TValue> from c# with some modification, but we will have our own interface for now.

// To read and write the map from database, tag the table and map key like so
// type UserModel struct {
//     ...
//     MapField generic.Dictionary[KeyType, ValueType]  `table:"table_name" key:"col_name"`
// }
// Then call Dictionary.LoadFromDb to load the relevant info.
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
	Map   map[K]*V
	Order []K
}

func (d *Dictionary[K, V]) Get(key K) (*V, bool) {
	value, exist := d.Map[key]
	return value, exist
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	if d.Map == nil {
		d.Map = make(map[K]*V)
	}
	if !d.Has(key) {
		d.Order = append(d.Order, key)
	}
	d.Map[key] = new(V)
	*d.Map[key] = value
}

func (d *Dictionary[K, V]) SetNull(key K) {
	if d.Map == nil {
		d.Map = make(map[K]*V)
	}
	if !d.Has(key) {
		d.Order = append(d.Order, key)
	}
	d.Map[key] = nil
}

func (d *Dictionary[K, V]) Has(key K) bool {
	_, exist := d.Map[key]
	return exist
}

func (d *Dictionary[K, V]) Sort() {
	sort.Slice(d.Order, func(i, j int) bool {
		return d.Order[i] < d.Order[j]
	})
}

func (d *Dictionary[K, V]) GetOnly(key K) *V {
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
	for _, key := range d.Order { // because the map is insert only, this is guaranteed to exist
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
		err := db.Table(table).Where("user_id = ?", userId).OrderBy(mapKey).Find(&values)
		utils.CheckErr(err)
		for _, v := range values {
			d.Set(reflect.ValueOf(v).Field(keyField).Interface().(K), v)
		}
	} else {
		rLoadFromDB := reflect.ValueOf(valueDummy).MethodByName("LoadFromDb")
		if rLoadFromDB.IsValid() {
			// V is a special wrapper like Nullable, we will use V's method to generate the value and key arrays
			var keys []any
			rLoadFromDB.Call([]reflect.Value{reflect.ValueOf(db), reflect.ValueOf(userId),
				reflect.ValueOf(table), reflect.ValueOf(mapKey), reflect.ValueOf(&keys), reflect.ValueOf(&values)})
			for i := range values {
				d.Set(keys[i].(K), values[i])
			}
		} else {
			fmt.Println("use default system", table, mapKey)
			// V is a raw type but the key doesn't exists
			// we fetch the object and then the key separately, both time ordered
			err := db.Table(table).Where("user_id = ?", userId).OrderBy(mapKey).Find(&values)
			utils.CheckErr(err)
			var keys []K
			err = db.Table(table).Where("user_id = ?", userId).OrderBy(mapKey).Cols(mapKey).Find(&keys)
			utils.CheckErr(err)
			for i := range keys {
				d.Set(keys[i], values[i])
			}
		}
	}
}

func (d *Dictionary[K, V]) ToContents() []any {
	contents := []any{}
	// TODO(refactor): This rely on the ID of the item, change it
	for _, content := range d.Map {
		contents = append(contents, reflect.ValueOf(content).MethodByName("ToContent").
			Call([]reflect.Value{})[0].Interface())
	}
	return contents
}
