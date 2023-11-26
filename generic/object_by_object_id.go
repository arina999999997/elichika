package generic

import (
	"elichika/model"
	"elichika/utils"

	"reflect"

	"encoding/json"
)

// TODO: there might be some optimization that can be done here

// Common pattern of [object_1_id, object_1, object_2_id, object_2...]

type ObjectByObjectIDList[T any] struct {
	Length  int // length
	Objects []T // slice of items
}

// Note that this only work on objects, not pointer.
// This is a limitation but also by design.
// For example, somethings use the (id, null) pattern to delete an existing items
// one design would be to have a null object
// but a null object can't contain ID, so we actually have to have a wrapper or some marking field
// by default we check for IsNull, if it is true then the object get rendered as nil

// Unmarshal: from JSON bytes to value
// require method SetID for the values
// return an empty array if data is null
func (oboid *ObjectByObjectIDList[T]) UnmarshalJSON(data []byte) error {
	oboid.Objects = []T{}
	oboid.Length = 0
	if string(data) == "null" {
		return nil
	}
	// TODO: remove this once done with placeholders
	if reflect.TypeOf(oboid.Objects).Elem() == reflect.ValueOf(0).Type() {
		return nil
	}
	arr := []json.RawMessage{}
	err := json.Unmarshal(data, &arr) // first unmarshal into an array of raw json
	utils.CheckErr(err)

	for i := range arr {
		if i%2 == 1 { // this is an object
			continue
		}
		// this is ID, we create a new value
		var id int64
		var obj T
		err = json.Unmarshal(arr[i], &id)
		if err != nil {
			return err
		}
		err = json.Unmarshal(arr[i+1], &obj)
		if err != nil {
			return err
		}
		setID := reflect.ValueOf(&obj).MethodByName("SetID")
		if setID.IsValid() {
			// if there's a SetID method then call it
			// this should only be done for structs where the ID are not present in json
			reflect.ValueOf(&obj).MethodByName("SetID").Call([]reflect.Value{reflect.ValueOf(id)})
		} else {
			// make sure we have things correctly by calling ID
			if id != reflect.ValueOf(&obj).MethodByName("ID").Call([]reflect.Value{})[0].Interface().(int64) {
				panic("ID doesn't match list provided id")
			}
		}
		oboid.Objects = append(oboid.Objects, obj) // append the pointer of the original type
		oboid.Length++
	}
	return nil
}

// Convert object to ID
// require method ID to get ID for the values
func (oboid ObjectByObjectIDList[T]) MarshalJSON() ([]byte, error) {
	arr := []any{}
	for _, object := range oboid.Objects {
		id := reflect.ValueOf(&object).MethodByName("ID").Call([]reflect.Value{})[0].Interface().(int64)
		arr = append(arr, id)
		isNull := reflect.ValueOf(object).FieldByName("IsNull")
		if isNull.IsValid() && isNull.Interface().(bool) {
			arr = append(arr, nil)
		} else {
			arr = append(arr, object)
		}
	}
	bytes, err := json.Marshal(&arr)
	return bytes, err
}

// append an object
func (oboid *ObjectByObjectIDList[T]) PushBack(obj T) {
	oboid.Length++
	oboid.Objects = append(oboid.Objects, obj)
}

// append a zero valued object and return a pointer to said object
func (oboid *ObjectByObjectIDList[T]) AppendNew() *T {
	var object T
	oboid.Objects = append(oboid.Objects, object)
	oboid.Length++
	return &oboid.Objects[oboid.Length-1]
}

// append a zero valued object and return a pointer to said object
func (oboid *ObjectByObjectIDList[T]) AppendNewWithID(id int64) *T {
	var object T
	reflect.ValueOf(&object).MethodByName("SetID").Call([]reflect.Value{reflect.ValueOf(id)})
	oboid.Objects = append(oboid.Objects, object)
	oboid.Length++
	return &oboid.Objects[oboid.Length-1]
}

func (oboid *ObjectByObjectIDList[T]) ToContents() []model.Content {
	contents := []model.Content{}
	for i := range oboid.Objects {
		contents = append(contents, reflect.ValueOf(&oboid.Objects[i]).MethodByName("ToContent").
			Call([]reflect.Value{})[0].Interface().(model.Content))
	}
	return contents
}

func (oboid *ObjectByObjectIDList[T]) SetUserID(uid int) {
	for i := range oboid.Objects {
		rUserID := reflect.Indirect(reflect.ValueOf(&oboid.Objects[i])).FieldByName("UserID")
		if rUserID.IsValid() {
			rUserID.Set(reflect.ValueOf(uid))
		} else {
			// fmt.Println("skip UserID for",  reflect.ValueOf(&oboid.Objects[i]).Type())
			return
		}
	}
}

// handler for an array object, use a map to map to the value for easier selection / tracking
// note that we don't store the object in the map itself because that lead to complication with xorm, as xorm can't use the ID function and rely on pk mapping
type ObjectByObjectIDMapping[T any] struct {
	List *ObjectByObjectIDList[T]
	Map  map[int64]int
}

func (m *ObjectByObjectIDMapping[T]) SetList(list *ObjectByObjectIDList[T]) *ObjectByObjectIDMapping[T] {
	if m.List != list {
		// new list, reset the map, and recalculate the mapping too (needed for importing account)
		m.List = list
		m.Map = make(map[int64]int)
		for i := range list.Objects {
			m.Map[reflect.ValueOf(&list.Objects[i]).MethodByName("ID").Call([]reflect.Value{})[0].Interface().(int64)] = i
		}
	}
	return m
}

// create a new list if there is none
func (m *ObjectByObjectIDMapping[T]) NewList() *ObjectByObjectIDMapping[T] {
	if m.List == nil {
		m.List = new(ObjectByObjectIDList[T])
		m.Map = make(map[int64]int)
	}
	return m
}

// update or insert an object, require ID
// copy the object and return the new pointer
func (m *ObjectByObjectIDMapping[T]) Update(object T) {
	id := reflect.ValueOf(&object).MethodByName("ID").Call([]reflect.Value{})[0].Interface().(int64)
	pos, exist := m.Map[id]
	if exist {
		m.List.Objects[pos] = object
	} else {
		m.List.PushBack(object)
		m.Map[id] = m.List.Length - 1
	}
}

// insert by ID and return the pointer to the object, require SetID
func (m *ObjectByObjectIDMapping[T]) InsertNew(id int64) *T {
	ptr := m.List.AppendNewWithID(id)
	m.Map[id] = m.List.Length - 1
	return ptr
}

func (m *ObjectByObjectIDMapping[T]) GetObject(id int64) *T {
	pos, exist := m.Map[id]
	if !exist {
		panic("Item doesn't exist")
	}
	return &m.List.Objects[pos]
}
