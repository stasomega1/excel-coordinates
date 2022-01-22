# excel-coordinates

### Library to easily generate map of excel-cell/value

***

#### Examples of usage:

```go
package main

import (
	"fmt"
	"strconv"

	coordinates "github.com/stasomega1/excel-coordinates"
)

type ExcelStruct struct {
	FirstName  string `excol:"A"`
	MiddleName string `excol:"B"`
	LastName   string `excol:"C"`
}

func main() {
	MyStructSlice := make([]*ExcelStruct, 0)
	for i := 0; i < 10; i++ {
		MyStruct := &ExcelStruct{
			FirstName:  "FirstName" + strconv.Itoa(i),
			MiddleName: "MiddleName" + strconv.Itoa(i),
			LastName:   "LastName" + strconv.Itoa(i),
		}
		MyStructSlice = append(MyStructSlice, MyStruct)
	}

	myMap, err := coordinates.GetCellMapWithRow(MyStructSlice, 1)
	if err != nil {
		panic(err)
	}

	for cell, value := range myMap {
		fmt.Printf("%v is located at %s cell\n", value, cell)
		//myExcelFile.Write(cell, value)
	}
}

```