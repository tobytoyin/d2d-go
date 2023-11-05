package adapters

import (
	adapters "d2d/adapters/mock"
	"d2d/contracts"
	"testing"
)

func GetMockSource() contracts.Source {
	source := &contracts.Source{"../../examples/sample_text.txt", "mock", nil}

	return *source
}

func GetMockTaskCatalog() contracts.SourceTask {
	return new(adapters.MockTasks)
}

func TestMockSourceIO(t *testing.T) {

	source := GetMockSource()
	tasks := GetMockTaskCatalog()

	contents := string(tasks.SourceIO(&source))

	if contents != "hello world" {
		t.Fatalf("SourceIO is not loading contents correctly")
	}

}

func TestUid(t *testing.T) {
	source := GetMockSource()
	tasks := GetMockTaskCatalog()
	uid := tasks.Uid(&source)

	if uid != "sample_text" {
		t.Fatalf("Uid is not not loading UID correctly")
	}

}

func TestSummary(t *testing.T) {
	source := GetMockSource()
	task := GetMockTaskCatalog()

	summary := task.Summary(&source)

	if summary.Content != "mock summary" {
		t.Fatalf("Summary is not loading contents correctly")
	}
}

func TestMetadata(t *testing.T) {
	source := GetMockSource()
	tasks := GetMockTaskCatalog()

	metadata := tasks.Metadata(&source)

	if metadata.Type != "mock" {
		t.Fatalf("Metadata.Type is not loading correctly")
	}

	if metadata.Properties == nil {
		t.Fatalf("Metadata.Properties is not load correctly")
	}
}
