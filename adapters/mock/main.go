package adapters

import (
	"d2d/contracts"
	"fmt"
	"os"
)

// type alias for private functions handle SourceIO
type SourceIOFunc func(string) interface{}

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

// wrapper function to allow other functions to source SourceIO returned string
func (task *MockTasks) withSourceIO(source *contracts.Source, fn SourceIOFunc) interface{} {
	contents := string(task.SourceIO(source))

	return fn(contents)
}

// private function to extract the summary from string
func summary(content string) interface{} {
	return contracts.DocSummary(content)
}

func (task *MockTasks) Summary(source *contracts.Source) interface{} {
	return task.withSourceIO(*source, summary)
}

// 	return contracts.DocSummary(contents)
// }

// func (task *MockTasks) Metadata(source *contracts.Source) contracts.DocMetadata {
// 	// reuse sourceIO
// 	contents := string(task.SourceIO(source))
// }
