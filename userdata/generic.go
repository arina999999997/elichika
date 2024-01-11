package userdata

import (
	"elichika/utils"

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

	genericTableFieldPopulators map[string]string // map from a database table name to a UserModel field
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

func addGenericTableFieldPopulator(tableName, fieldName string) {
	if genericTableFieldPopulators == nil {
		genericTableFieldPopulators = make(map[string]string)
	}
	genericTableFieldPopulators[fieldName] = tableName
}

func newGenericTableFieldPopulator(session *Session) {
	// TODO(refactor): These can be init at the start or something
	rModel := reflect.ValueOf(&session.UserModel)
	for i := 0; i < rModel.Type().Elem().NumField(); i++ {
		rFieldType := rModel.Type().Elem().Field(i)
		tableName := rFieldType.Tag.Get("table")
		keyColumn := rFieldType.Tag.Get("key")
		if tableName == "" {
			genericTableFieldPopulator(session, rFieldType.Name)
			continue
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

func genericTableFieldPopulator(session *Session, fieldName string) {
	tableName, exists := genericTableFieldPopulators[fieldName]
	if !exists {
		return
	}
	rModel := reflect.ValueOf(&session.UserModel)
	rField := rModel.Elem().FieldByName(fieldName)
	if !rField.IsValid() {
		fmt.Println("Invalid table field pair: ", tableName, "->", fieldName)
		return
	}
	fmt.Println(tableName, fieldName)
	if fieldName == "UserMemberByMemberId" {
		// TODO(refactor): This is temporary

	} else {
		err := session.Db.Table(tableName).Where("user_id = ?", session.UserId).
			Find(rField.FieldByName("Objects").Addr().Interface())
		utils.CheckErr(err)
	}
}

func init() {
	// addPopulator(genericTableFieldPopulator)
	addPopulator(newGenericTableFieldPopulator)
}
