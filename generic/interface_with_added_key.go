package generic

import (
	"reflect"
)

// Add the key to the interface and then return the interface
// currently only support int for INTEGER and string for TEXT
// The key will appear in the order of the arrays, before the existing fields, and it can contain pk and such too
func InterfaceWithAddedKey[keyType int32 | string](original interface{}, names []string, xormTags []reflect.StructTag) interface{} {
	var key keyType
	fields := []reflect.StructField{}
	for i, tag := range xormTags {
		fields = append(fields, reflect.StructField{
			Name: names[i],
			Type: reflect.TypeOf(key),
			Tag:  tag,
		})
	}
	fields = append(fields, reflect.VisibleFields(reflect.TypeOf(original))...)
	return reflect.New(reflect.StructOf(fields)).Elem().Interface()
}
