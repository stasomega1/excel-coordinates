package excel_coordinates

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotASlice = errors.New("input data is not a slice")
)

func GetCellMapWithRow(data interface{}, fromRow int) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return nil, ErrNotASlice
	}

	s := reflect.ValueOf(data)
	for i := 0; i < s.Len(); i++ {
		dataValue := s.Index(i)
		if dataValue.Kind() == reflect.Ptr {
			dataValue = dataValue.Elem()
		}
		typeData := dataValue.Type()
		for i := 0; i < typeData.NumField(); i++ {
			field := typeData.Field(i)
			cell := fmt.Sprintf("%s%d", field.Tag.Get("excol"), fromRow)
			result[cell] = fmt.Sprintf("%v", dataValue.Field(i).Interface())
		}
		fromRow++
	}

	return result, nil
}
