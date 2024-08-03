package parser

import (
	"elichika/client"
	"elichika/utils"

	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type CsvContext struct {
	// inclusively parse the field in [StartField, EndField)
	// if EndField is 0, parse from start to the end
	StartField int
	EndField   int
	HasHeader  bool
}

func ParseCsv[T any](path string, result *[]T, ctx *CsvContext) {
	if ctx == nil {
		ctx = &CsvContext{
			StartField: 0,
			EndField:   0,
			HasHeader:  false,
		}
	}
	var dummy T
	dummyPtr := &dummy
	ptr := reflect.ValueOf(dummyPtr)
	ioReader, err := os.Open(path)
	utils.CheckErr(err)
	reader := csv.NewReader(ioReader)
	if ctx.EndField == 0 {
		ctx.EndField = ptr.Elem().Type().NumField()
	}
	reader.FieldsPerRecord = ctx.EndField - ctx.StartField

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		utils.CheckErr(err)
		if ctx.HasHeader {
			ctx.HasHeader = false
			continue
		}
		for i, str := range record {
			field := ptr.Elem().Type().Field(i + ctx.StartField)
			switch field.Type {
			case reflect.TypeOf(int32(0)):
				valueInt, err := strconv.Atoi(str)
				utils.CheckErr(err)
				value := int32(valueInt)
				ptr.Elem().Field(i + ctx.StartField).Set(reflect.ValueOf(value))
			case reflect.TypeOf(false):
				valueInt, err := strconv.Atoi(str)
				utils.CheckErr(err)
				value := (valueInt != 0)
				ptr.Elem().Field(i + ctx.StartField).Set(reflect.ValueOf(value))
			case reflect.TypeOf(string("")):
				ptr.Elem().Field(i + ctx.StartField).Set(reflect.ValueOf(str))
			case reflect.TypeOf(client.LocalizedText{}):
				ptr.Elem().Field(i + ctx.StartField).Set(reflect.ValueOf(client.LocalizedText{
					DotUnderText: str,
				}))
			default:
				fmt.Println(field.Type)
				panic("field type not supported")
			}
		}
		*result = append(*result, dummy)
	}
}
