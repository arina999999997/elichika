package user_content

import (
	"elichika/client"
	"elichika/userdata"

	"fmt"
	"reflect"
)

// this is used for account importing
func PopulateGenericContentDiffFromUserModel(session *userdata.Session) {
	rModel := reflect.ValueOf(&session.UserModel)
	for contentType, fieldName := range userModelField {
		rDictionary := rModel.Elem().FieldByName(fieldName)
		if !rDictionary.IsValid() {
			fmt.Println("Invalid field: ", contentType, "->", fieldName)
			continue
		}
		rDictionaryPtrType := reflect.PointerTo(rDictionary.Type())
		rDictionaryToContents, ok := rDictionaryPtrType.MethodByName("ToContents")
		if !ok {
			panic(fmt.Sprintln("Type ", rDictionaryPtrType, " must have method ToContents"))
		}
		contents := rDictionaryToContents.Func.Call([]reflect.Value{rDictionary.Addr()})[0].Interface().([]any)
		for _, content := range contents {
			UpdateUserContent(session, content.(client.Content))
		}
	}
}
