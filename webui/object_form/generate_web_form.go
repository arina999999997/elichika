package object_form

import (
	"fmt"
	htmllib "html"
	"reflect"
	"strings"
)

var externalOptions = map[string][]string{}

func SetExternalOptions(key string, options []string) {
	externalOptions[key] = options
}

func GenerateWebForm(defaultObjPtr any, formId, buttonAction, resetText, submitText string) string {
	html := `<form id="` + formId + `" method="POST"  enctype="multipart/form-data" onkeydown="if(event.keyCode === 13) {
		alert('You have pressed Enter key, use submit button instead'); 
		return false;
	}">` + "\n"
	// html
	ptr := reflect.ValueOf(defaultObjPtr) // pointer
	if ptr.Kind() != reflect.Pointer {
		panic("must pass a pointer to object")
	}

	for i := 0; i < ptr.Elem().Type().NumField(); i++ {
		html += `<div>` + "\n"
		field := ptr.Elem().Type().Field(i)
		{
			label := field.Tag.Get("of_label")
			if label == "" {
				label = field.Name
			}
			html += `<label>` + label + `</label>` + "\n"
		}
		customType := field.Tag.Get("of_type")
		if customType == "select" {
			currentValue := fmt.Sprint(reflect.Indirect(ptr.Elem().Field(i)).Interface())
			html += `<select name="` + field.Name + `">` + "\n"
			optionsString := field.Tag.Get("of_options")
			var options []string
			if optionsString == "" {
				// kind of a hack but it allow us to set options at runtime
				externalKey := field.Tag.Get("of_options_external")
				if externalKey == "" {
					panic("no option of external key provided")
				}
				options = externalOptions[externalKey]
			} else {
				options = strings.Split(optionsString, "\n")
			}
			n := len(options)
			if n%2 == 1 {
				panic("wrong of_options")
			}
			for i := 0; i < n; i += 2 {
				html += `<option value="` + options[i+1] + `"`
				if options[i+1] == currentValue {
					html += " selected"
				}
				html += ">"
				html += htmllib.EscapeString(options[i])
				html += `</option>` + "\n"
			}
			html += `</select>` + "\n"
		} else {
			html += `<input name="` + field.Name + `" `
			if customType == "time" {
				// second since midnight
				html += `type="time" step="1" value="`
				switch field.Type {
				case reflect.TypeOf((*int32)(nil)):
					value := reflect.Indirect(ptr.Elem().Field(i)).Interface().(int32)
					html += fmt.Sprintf("%02d:%02d:%02d", value/3600, value%3600/60, value%60)
				case reflect.TypeOf((*string)(nil)):
					value := reflect.Indirect(ptr.Elem().Field(i)).Interface().(string)
					html += value
				default:
					panic("field type not supported")
				}
				html += `"`
			} else {
				switch field.Type {
				case reflect.TypeOf((*string)(nil)):
					html += `type="text" value="`
					if customType != "password" {
						html += reflect.Indirect(ptr.Elem().Field(i)).Interface().(string)
					}
					html += `"`
				case reflect.TypeOf((*int32)(nil)):
					html += `type="number" value="`
					html += fmt.Sprint(reflect.Indirect(ptr.Elem().Field(i)).Interface().(int32))
					html += `"`
				case reflect.TypeOf((*bool)(nil)):
					html += `type="checkbox"`
					if reflect.Indirect(ptr.Elem().Field(i)).Interface().(bool) {
						html += ` checked`
					}
				case reflect.TypeOf((string)("")):
					html += `type="text" value="`
					if customType != "password" {
						html += ptr.Elem().Field(i).Interface().(string)
					}
					html += `"`
				case reflect.TypeOf((int32)(0)):
					html += `type="number" value="`
					html += fmt.Sprint(ptr.Elem().Field(i).Interface().(int32))
					html += `"`
				case reflect.TypeOf((bool)(false)):
					html += `type="checkbox"`
					if ptr.Elem().Field(i).Interface().(bool) {
						html += ` checked`
					}
				default:
					panic("field type not supported")
				}
				extraTags := field.Tag.Get("of_attrs")
				if extraTags != "" {
					html += ` ` + extraTags
				}
			}
			html += `/>`
		}

		html += "</div>\n"
	}
	if resetText != "" {
		html += `<div><input type="reset" value="`
		html += resetText
		html += `"/></div>` + "\n"
	}

	html += `<div><input type="button" value="`
	html += submitText + `"`
	html += buttonAction + `/></div>` + "\n"

	html += `</form>` + "\n"
	return html
}
