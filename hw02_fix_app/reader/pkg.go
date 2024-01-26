package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ildar52/hw-test/types"
)

func ReadJSON(
	filePath string,
) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	res := data

	return res, nil
}
