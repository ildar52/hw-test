package main

import (
	"fmt"

	"github.com/ildar52/hw-test/printer"
	"github.com/ildar52/hw-test/reader"
	"github.com/ildar52/hw-test/types"
)

func main() {
	path := "data.json"
	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
