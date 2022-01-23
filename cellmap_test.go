package excel_coordinates

import (
	"testing"
)

type StructToExcel struct {
	Field1 string  `excol:"A"`
	Field2 int     `excol:"B"`
	Field3 float64 `excol:"C"`
	Field4 *string `excol:"D"`
	Field5 bool    `excol:"F"`
}

func getPtrStr(value string) *string {
	return &value
}

func TestGetCellMapWithRow(t *testing.T) {
	type TestStruct struct {
		have    []StructToExcel
		havePtr []*StructToExcel
		want    map[string]string
	}

	test := TestStruct{
		have: []StructToExcel{
			{
				Field1: "strField1",
				Field2: 1,
				Field3: 0.1,
				Field4: getPtrStr("strPtrField1"),
				Field5: true,
			},
			{
				Field1: "strField2",
				Field2: 2,
				Field3: 0.2,
				Field4: getPtrStr("strPtrField2"),
				Field5: true,
			},
			{
				Field1: "strField3",
				Field2: 3,
				Field3: 3.0,
				Field4: getPtrStr("strPtrField3"),
				Field5: true,
			},
		},
		havePtr: []*StructToExcel{
			{
				Field1: "strField1",
				Field2: 1,
				Field3: 0.1,
				Field4: getPtrStr("strPtrField1"),
				Field5: true,
			},
			{
				Field1: "strField2",
				Field2: 2,
				Field3: 0.2,
				Field4: getPtrStr("strPtrField2"),
				Field5: true,
			},
			{
				Field1: "strField3",
				Field2: 3,
				Field3: 3.0,
				Field4: getPtrStr("strPtrField3"),
				Field5: true,
			},
		},
		want: map[string]string{
			"A1": "strField1",
			"A2": "strField2",
			"A3": "strField3",
			"B1": "1",
			"B2": "2",
			"B3": "3",
			"C1": "0.1",
			"C2": "0.2",
			"C3": "3",
			"D1": "strPtrField1",
			"D2": "strPtrField2",
			"D3": "strPtrField3",
			"F1": "true",
			"F2": "true",
			"F3": "true",
		},
	}

	result, err := GetCellMapWithRow(test.have, 1)
	if err != nil {
		t.Error(err)
	}

	resultPtr, err := GetCellMapWithRow(test.havePtr, 1)
	if err != nil {
		t.Error(err)
	}

	for cell, value := range test.want {
		if result[cell] != value {
			t.Errorf("result failed, cell: %s, want %s, have %s", cell, value, result[cell])
		}
		if resultPtr[cell] != value {
			t.Errorf("resultPtr failed, cell: %s, want %s, have %s", cell, value, result[cell])
		}
	}
}
