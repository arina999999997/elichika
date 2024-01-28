package userdata

// import (
// 	"elichika/client"

// 	"fmt"
// 	"reflect"
// )

// this is used for loading exported account data
// TODO(now)
func (session *Session) populateGenericContentDiffFromUserModel() {
	// rModel := reflect.ValueOf(&session.UserModel)
	// for contentType, fieldName := range userModelField {
	// 	rDictionary := rModel.Elem().FieldByName(fieldName)
	// 	if !rDictionary.IsValid() {
	// 		fmt.Println("Invalid field: ", contentType, "->", fieldName)
	// 		continue
	// 	}
	// 	rDictionaryPtrType := reflect.PointerTo(rDictionary.Type())
	// 	rDictionaryToContents, ok := rDictionaryPtrType.MethodByName("ToContents")
	// 	if !ok {
	// 		panic(fmt.Sprintln("Type ", rDictionaryPtrType, " must have method ToContents"))
	// 	}
	// 	contents := rDictionaryToContents.Func.Call([]reflect.Value{rDictionary.Addr()})[0].Interface().([]any)
	// 	for _, content := range contents {
	// 		session.UpdateUserContent(content.(client.Content))
	// 	}
	// }

}
