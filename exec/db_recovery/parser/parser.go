package parser

import (
	"errors"
	"reflect"
	"strings"
)

const (
	INSERT = 0
	DELETE = -1
	UPDATE = 1
)

type Field struct {
	Key   string
	Value string
}

func FieldFromString(pair string) Field {
	if !strings.HasPrefix(pair, "/*") {
		panic("Unexpected pair format: " + pair)
	}
	pair = pair[2:]
	strs := strings.Split(pair, "*/")
	if len(strs) != 2 {
		panic("Unexpected pair format: " + pair)
	}
	return Field{
		Key:   strs[0],
		Value: strs[1],
	}
}

type ModifierEvent[T any] struct {
	Fields []Field
	Type   int
}

func (this *ModifierEvent[T]) Id() int64 {
	var dummy *T
	return reflect.ValueOf(dummy).MethodByName("Id").Call([]reflect.Value{reflect.ValueOf(this.Fields)})[0].Interface().(int64)
}

func (this *ModifierEvent[T]) String() string {
	var dummy *T
	res := ""
	switch this.Type {
	case INSERT:
		res += "INSERT INTO "

		res += reflect.ValueOf(dummy).MethodByName("Table").Call([]reflect.Value{})[0].Interface().(string)
		res += " VALUES ("
		for i, f := range this.Fields {
			if i > 0 {
				res += ", "
			}
			res += reflect.ValueOf(dummy).MethodByName("Value").Call([]reflect.Value{reflect.ValueOf(f)})[0].Interface().(string)
		}
		res += ");"
	case DELETE:
		res += "DELETE FROM "
		res += reflect.ValueOf(dummy).MethodByName("Table").Call([]reflect.Value{})[0].Interface().(string)
		res += " WHERE "
		res += reflect.ValueOf(dummy).MethodByName("Condition").Call([]reflect.Value{reflect.ValueOf(this.Fields)})[0].Interface().(string)
		res += ";"
	case UPDATE:
		res += "UPDATE "
		res += reflect.ValueOf(dummy).MethodByName("Table").Call([]reflect.Value{})[0].Interface().(string)
		res += " SET "
		for i, f := range this.Fields {
			if i > 0 {
				res += ", "
			}
			res += reflect.ValueOf(dummy).MethodByName("Update").Call([]reflect.Value{reflect.ValueOf(f)})[0].Interface().(string)
		}
		res += " WHERE "
		res += reflect.ValueOf(dummy).MethodByName("Condition").Call([]reflect.Value{reflect.ValueOf(this.Fields)})[0].Interface().(string)
		res += ";"
	default:
		panic("Unknown type")
	}
	return res
}

func TryParse[T any](line string) (ModifierEvent[T], error) {
	res := ModifierEvent[T]{}
	if !(strings.HasPrefix(line, "+INSERT INTO") || strings.HasPrefix(line, "-INSERT INTO")) {
		return res, errors.New("not an event")
	}
	if line[0] == '+' {
		res.Type = INSERT
	} else {
		res.Type = DELETE
	}

	line = line[strings.Index(line, "(")+1:]
	line = line[:len(line)-2]
	fields := strings.Split(line, ", ")
	for _, field := range fields {
		res.Fields = append(res.Fields, FieldFromString(field))
	}
	return res, nil
}

func Parse[T any](input string) map[int64]ModifierEvent[T] {
	lines := strings.Split(input, "\n")
	res := map[int64]ModifierEvent[T]{}
	current := map[int64]ModifierEvent[T]{}
	// parse from top down, so from newest commit to oldest
	// within each commit, - is shown before + so we can just overwrite existing -for an update
	for _, line := range lines {
		if strings.HasPrefix(line, "commit") {
			// merge the current commit
			for key, value := range current {
				_, exists := res[key]
				if exists {
					continue
				}
				res[key] = value
			}
			current = map[int64]ModifierEvent[T]{}
			continue
		}
		event, err := TryParse[T](line)
		if err != nil {
			continue
		}
		current[event.Id()] = event
	}
	for key, value := range current {
		_, exists := res[key]
		if exists {
			continue
		}
		res[key] = value
	}
	return res
}
