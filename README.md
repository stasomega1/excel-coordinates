# excel-coordinates

### Library to easily generate map of excel-cell/value

***

#### Examples of usage:

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	coordinates "github.com/stasomega1/excel-coordinates"
	"github.com/xuri/excelize/v2"
)

type ExcelStruct struct {
	FirstName  *string `excol:"A"`
	MiddleName int     `excol:"B"`
	LastName   bool    `excol:"C"`
}

func newStrPtr(str string) *string {
	return &str
}

func main() {
	MyStructSlice := make([]*ExcelStruct, 0)
	for i := 0; i < 10; i++ {
		MyStruct := &ExcelStruct{
			FirstName:  newStrPtr("FirstName" + strconv.Itoa(i)),
			MiddleName: i,
			LastName:   false,
		}
		MyStructSlice = append(MyStructSlice, MyStruct)
	}

	myMap, err := coordinates.GetCellMapWithRow(MyStructSlice, 1)
	if err != nil {
		panic(err)
	}
	//
	exFile := excelize.NewFile()
	sheet := "Sheet1"
	exFile.NewSheet(sheet)
	for cell, value := range myMap {
		fmt.Printf("%v is located at %s cell\n", value, cell)
		exFile.SetCellValue(sheet, cell, value)
	}

	file, _ := os.Create("filename.xlsx")
	defer file.Close()
	exFile.WriteTo(file)
}
```