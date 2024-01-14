package userdata

import (
	"fmt"
	"reflect"
)

// there are 2 types of generic handlers
// - populators will be used to load login data into the session
// - finalizers will be used to finalize data and write to database
// - it is possible to have a populator but not a finalizer for a data type and vice versa
// - other than that, there might be other generic type of handler but they should be handled by another system
// calling order of handlers are not guranteed, so they have to be implemented to accomodate for that
type handler = func(*Session)

var (
	populators map[uintptr]handler
	finalizers map[uintptr]handler
)

func addPopulator(p handler) {
	if populators == nil {
		populators = make(map[uintptr]handler)
		finalizers = make(map[uintptr]handler)
	}
	populators[reflect.ValueOf(p).Pointer()] = p
}

func addFinalizer(f handler) {
	if finalizers == nil {
		populators = make(map[uintptr]handler)
		finalizers = make(map[uintptr]handler)
	}
	finalizers[reflect.ValueOf(f).Pointer()] = f
}

// TODO(refactor): This is kinda ugly
func (session *Session) PopulateUserModelField(fieldName string) {
	rModel := reflect.ValueOf(&session.UserModel)
	for i := 0; i < rModel.Type().Elem().NumField(); i++ {
		rFieldType := rModel.Type().Elem().Field(i)
		if rFieldType.Name != fieldName {
			continue
		}
		tableName := rFieldType.Tag.Get("table")
		keyColumn := rFieldType.Tag.Get("key")
		rField := rModel.Elem().Field(i)
		rMethod := rField.Addr().MethodByName("LoadFromDb")
		if rMethod.IsValid() {
			rMethod.Call([]reflect.Value{reflect.ValueOf(session.Db), reflect.ValueOf(session.UserId),
				reflect.ValueOf(tableName), reflect.ValueOf(keyColumn)})
		} else {
			panic(fmt.Sprint("Tagged but not supported: ", i, rField, rMethod, tableName, keyColumn))
		}
	}
}

func genericTableFieldPopulator(session *Session) {
	// TODO(refactor): These can be init at the start or something
	rModel := reflect.ValueOf(&session.UserModel)
	for i := 0; i < rModel.Type().Elem().NumField(); i++ {
		rFieldType := rModel.Type().Elem().Field(i)
		tableName := rFieldType.Tag.Get("table")
		keyColumn := rFieldType.Tag.Get("key")
		if rFieldType.Name == "UserStatus" || tableName == "u_resource" {
			continue
		} else if tableName == "" {
			panic(rFieldType.Name)
		}
		rField := rModel.Elem().Field(i)
		rMethod := rField.Addr().MethodByName("LoadFromDb")
		if rMethod.IsValid() {
			rMethod.Call([]reflect.Value{reflect.ValueOf(session.Db), reflect.ValueOf(session.UserId),
				reflect.ValueOf(tableName), reflect.ValueOf(keyColumn)})
		} else {
			panic(fmt.Sprint("Tagged but not supported: ", i, rField, rMethod, tableName, keyColumn))
		}
	}
}

func init() {
	addPopulator(genericTableFieldPopulator)
}
