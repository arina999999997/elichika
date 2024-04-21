package object_form

import (
	"errors"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseForm(ctx *gin.Context, defaultObjPtr any) error {
	form := ctx.MustGet("form").(*multipart.Form)
	ptr := reflect.ValueOf(defaultObjPtr) // pointer
	if ptr.Kind() != reflect.Pointer {
		return errors.New("must pass a pointer to object")
	}
	// TODO: this only works for pointer fields for now
	for i := 0; i < ptr.Elem().Type().NumField(); i++ {

		field := ptr.Elem().Type().Field(i)
		if field.Type.Kind() == reflect.Pointer {
			ptr.Elem().Field(i).Set(reflect.New(ptr.Elem().Field(i).Type().Elem()))
		}
		if (field.Type == reflect.TypeOf((*bool)(nil))) || (field.Type == reflect.TypeOf((bool)(false))) {
			onString, on := form.Value[field.Name]
			if on && (onString[0] != "on") {
				return errors.New("explicit off checkbox?")
			}
			if field.Type.Kind() == reflect.Pointer {
				reflect.Indirect(ptr.Elem().Field(i)).Set(reflect.ValueOf(on))
			} else {
				ptr.Elem().Field(i).Set(reflect.ValueOf(on))
			}
			continue
		}

		stringValue := form.Value[field.Name][0]
		customType := field.Tag.Get("of_type")

		if customType == "select" || customType == "" {
			// there is no special type, just parse directly
			switch field.Type {
			case reflect.TypeOf((*string)(nil)):
				reflect.Indirect(ptr.Elem().Field(i)).Set(reflect.ValueOf(stringValue))
			case reflect.TypeOf((*int32)(nil)):
				value, err := strconv.ParseInt(stringValue, 10, 32)
				if err != nil {
					return err
				}
				reflect.Indirect(ptr.Elem().Field(i)).Set(reflect.ValueOf(int32(value)))
			default:
				return errors.New("field type not supported")
			}
		} else if customType == "time" {
			switch field.Type {
			case reflect.TypeOf((*string)(nil)):
				reflect.Indirect(ptr.Elem().Field(i)).Set(reflect.ValueOf(stringValue))
			case reflect.TypeOf((*int32)(nil)):
				parts := strings.Split(stringValue, ":")
				ints := []int32{}
				for j := 0; j < 3; j++ {
					value, err := strconv.ParseInt(parts[j], 10, 32)
					if err != nil {
						return err
					}
					ints = append(ints, int32(value))
				}
				seconds := ints[0]*3600 + ints[1]*60 + ints[2]
				reflect.Indirect(ptr.Elem().Field(i)).Set(reflect.ValueOf(int32(seconds)))
			}

		}
	}
	return nil
}
