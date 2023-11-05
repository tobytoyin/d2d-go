package adapters

import (
	"d2d/contracts"
	"fmt"
	"os"
	"strings"
)

// mock implementation of the Service Catalog
type MockTasks struct{}

// required methods for the SourceTask interface
func (task *MockTasks) SourceIO(source *contracts.Source) []byte {
	// open the file
	content, err := os.ReadFile(source.Path)

	if err != nil {
		fmt.Println(err)
		panic("Cannot Load the contents")
	}

	fmt.Println(content)
	return content

}

// method that return the Uid based on the source
func (task *MockTasks) Uid(source *contracts.Source) string {
	path := source.Path
	filename := strings.Split(path, "/")
	uid := strings.Split(filename[len(filename)-1], ".")[0]

	return uid

}

// method that return the Summary of the SourceIO's contents
func (task *MockTasks) Summary(source *contracts.Source) contracts.DocSummary {
	contents := string(task.SourceIO(source))
	fmt.Println(contents)
	return contracts.DocSummary{Content: "mock summary"}
}

// func (task *MockTasks) Metadata(source *contracts.Source) contracts.DocMetadata {
// 	// reuse sourceIO
// 	contents := string(task.SourceIO(source))
// }
