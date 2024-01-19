package printer

import (
	"fmt"

	"github.com/ildar52/hw-test/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		str = fmt.Sprintf("UserID: %s; Age: %d; Name: %s; DepartmentID: %d", staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(i)
	}

	fmt.Println(str)
}
