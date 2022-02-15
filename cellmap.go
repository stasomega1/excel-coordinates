package excel_coordinates

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotASlice = errors.New("input data is not a slice")
)

const (
	colTag = "excol"
)

func GetCellMapWithRow(data interface{}, fromRow int) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return nil, ErrNotASlice
	}

	s := reflect.ValueOf(data)
	//iterate over a slice
	for i := 0; i < s.Len(); i++ {
		dataValue := s.Index(i)
		//if slice is a pointer - get its value
		if dataValue.Kind() == reflect.Ptr {
			dataValue = dataValue.Elem()
		}
		typeData := dataValue.Type()
		//iterate over fields
		for i := 0; i < typeData.NumField(); i++ {
			field := typeData.Field(i)
			column := field.Tag.Get(colTag)
			if column == "" {
				continue
			}
			cell := fmt.Sprintf("%s%d", column, fromRow)
			dataField := dataValue.Field(i)
			//if field is a pointer - get its value
			if dataField.Kind() == reflect.Ptr {
				dataField = dataField.Elem()
			}
			if dataField.IsValid() {
				result[cell] = fmt.Sprintf("%v", dataField.Interface())
			} else {
				result[cell] = ""
			}
		}
		fromRow++
	}

	return result, nil
}
