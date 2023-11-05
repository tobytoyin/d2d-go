package adapters

import (
	"d2d/contracts"
	"fmt"
	"os"
)

type MockTasks struct {
}

// required methods for the SourceTask interface
func (task *MockTasks) SourceIO(source *contracts.Source) []byte {
	// open the file
	content, err := os.ReadFile(source.Path)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(content)
	return content

}
