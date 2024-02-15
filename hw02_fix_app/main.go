package main

import (
	"fmt"

	"github.com/ildar52/hw-test/hw02_fix_app/printer"
	"github.com/ildar52/hw-test/hw02_fix_app/reader"
	"github.com/ildar52/hw-test/hw02_fix_app/types"
)

func main() {
	path := "data.json"
	fmt.Println("Enter data file path: ")

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
