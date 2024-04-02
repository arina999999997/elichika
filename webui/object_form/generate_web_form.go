package object_form

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateWebForm(defaultObjPtr any, formAction, resetText, submitText string) string {
	html := "<form action=\"/webui\" method=\"POST\" enctype=\"multipart/form-data\">\n"
	// html
	ptr := reflect.ValueOf(defaultObjPtr) // pointer
	if ptr.Kind() != reflect.Pointer {
		panic("must pass a pointer to object")
	}

	for i := 0; i < ptr.Elem().Type().NumField(); i++ {
		html += "<div>\n"
		field := ptr.Elem().Type().Field(i)
		{
			label := field.Tag.Get("of_label")
			if label == "" {
				label = field.Name
			}
			html += "<label>" + label + "</label>\n"
		}
		customType := field.Tag.Get("of_type")
		if customType == "select" {
			currentValue := fmt.Sprint(reflect.Indirect(ptr.Elem().Field(i)).Interface())
			html += "<select name=\"" + field.Name + "\">\n"
			optionsString := field.Tag.Get("of_options")
			options := strings.Split(optionsString, "\n")
			n := len(options)
			if n%2 == 1 {
				panic("wrong of_options")
			}
			for i := 0; i < n; i += 2 {
				html += "<option value=\"" + options[i+1] + "\""
				if options[i+1] == currentValue {
					html += " selected"
				}
				html += ">"
				html += options[i]
				html += "</option>\n"
			}
			html += "</select>\n"
		} else {
			html += "<input name=\"" + field.Name + "\" "
			if customType == "time" {
				// second since midnight
				html += "type=\"time\" step=\"1\" value=\""
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
				html += "\""
			} else {
				switch field.Type {
				case reflect.TypeOf((*string)(nil)):
					html += "type=\"text\" value=\""
					html += reflect.Indirect(ptr.Elem().Field(i)).Interface().(string)
					html += "\""
				case reflect.TypeOf((*int32)(nil)):
					html += "type=\"number\" value=\""
					html += fmt.Sprint(reflect.Indirect(ptr.Elem().Field(i)).Interface().(int32))
					html += "\""
				case reflect.TypeOf((*bool)(nil)):
					html += "type=\"checkbox\""
					if reflect.Indirect(ptr.Elem().Field(i)).Interface().(bool) {
						html += " checked"
					}
				default:
					panic("field type not supported")
				}
				extraTags := field.Tag.Get("of_attrs")
				if extraTags != "" {
					html += " " + extraTags
				}
			}
			html += "/>"
		}

		html += "</div>\n"
	}

	html += "<div><input type=\"reset\" value=\""
	html += resetText
	html += "\"/></div>\n"
	html += "<div><input type=\"submit\" value=\""
	html += submitText
	html += "\" formaction=\"" + formAction + "\"/></div>\n"

	html += "</form>\n"
	return html
}
