package printer

import (
	"fmt"

	"github.com/ildar52/hw-test/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf("UserID: %d; Age: %d; Name: %s; DepartmentID: %d", staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(str)
	}
}
