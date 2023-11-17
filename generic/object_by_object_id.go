package generic

import (
	"elichika/utils"

	"reflect"

	"encoding/json"
)

// Common pattern of [object_1_id, object_1, object_2_id, object_2...]
// if null give an empty array
// need to pass the pointer type of a type because I don't know golang

type ObjectByObjectIDReadInterface interface { // only unmarshal
	SetID(int64)
}

type ObjectByObjectIDRead[T ObjectByObjectIDReadInterface] struct {
	Objects []T
}

func (oboid *ObjectByObjectIDRead[T]) UnmarshalJSON(data []byte) error {
	oboid.Objects = []T{}
	if string(data) == "null" {
		return nil
	}
	arr := []any{}
	err := json.Unmarshal(data, &arr)
	utils.CheckErr(err)
	for i, _ := range arr {
		if i%2 == 1 {
			continue
		}
		bytes, err := json.Marshal(arr[i+1])
		utils.CheckErr(err)
		ptp := new(T)                         // pointer to pointer of the original type
		err = json.Unmarshal(bytes, ptp)      // unmarshal into the pointer of pointer the original type, create a pointer and then an object
		(*ptp).SetID(int64(arr[i].(float64))) // set the id
		utils.CheckErr(err)
		oboid.Objects = append(oboid.Objects, *ptp) // append the pointer of the original type
	}
	return nil
}

type ObjectByObjectIDWriteInterface interface { // only marshal
	ID() int64
}

type ObjectByObjectIDWrite[T ObjectByObjectIDWriteInterface] struct {
	Length  int
	Objects []T
}

func (oboid ObjectByObjectIDWrite[T]) MarshalJSON() ([]byte, error) {
	arr := []any{}
	for _, object := range oboid.Objects {
		arr = append(arr, object.ID())
		arr = append(arr, object)
	}
	bytes, err := json.Marshal(&arr)
	utils.CheckErr(err)
	return bytes, nil
}

// append a pointer
func (oboid *ObjectByObjectIDWrite[T]) PushBack(ptr T) {
	oboid.Length += 1
	oboid.Objects = append(oboid.Objects, ptr)
}

// append a new pointer at the end, the pointer point to a zero value
// then return the pointer and the index
func (oboid *ObjectByObjectIDWrite[T]) AppendNew() T {
	ptp := new(T) // pointer to pointer, that hold nothing
	// a bit dirty but it works
	err := json.Unmarshal([]byte("{}"), ptp) // unmarshal into the pointer of pointer the original type, create a pointer and then an object
	utils.CheckErr(err)

	oboid.Objects = append(oboid.Objects, *ptp) // append the pointer and return it
	oboid.Length += 1
	return *ptp
}

type ObjectByObjectIDInterface interface { // only unmarshal
	SetID(int64)
	ID() int64
}

type ObjectByObjectID[T any] struct {
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
func (oboid *ObjectByObjectID[T]) UnmarshalJSON(data []byte) error {
	oboid.Objects = []T{}
	oboid.Length = 0
	if string(data) == "null" {
		return nil
	}
	arr := []json.RawMessage{}
	err := json.Unmarshal(data, &arr) // first unmarshal into an array of raw json
	utils.CheckErr(err)

	for i, _ := range arr {
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
		reflect.ValueOf(&obj).MethodByName("SetID").Call([]reflect.Value{reflect.ValueOf(&id)})
		if err != nil {
			return err
		}
		oboid.Objects = append(oboid.Objects, obj) // append the pointer of the original type
		oboid.Length += 1
	}
	return nil
}

// Convert object to ID
// require method ID to get ID for the values
func (oboid ObjectByObjectID[T]) MarshalJSON() ([]byte, error) {
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
func (oboid *ObjectByObjectID[T]) PushBack(obj T) {
	oboid.Length += 1
	oboid.Objects = append(oboid.Objects, obj)
}

// append a zero valued object and return a pointer to said object
func (oboid *ObjectByObjectID[T]) AppendNew() *T {
	var dummy T
	oboid.Objects = append(oboid.Objects, dummy)
	oboid.Length += 1
	return &oboid.Objects[oboid.Length-1]
}
