package adapters

import (
	"d2d/adapters"
	mock "d2d/adapters/mock"
	"reflect"
	"testing"
)

// test if the factory can return the struct given a string value
func TestAdapterFactory(t *testing.T) {

	adapter := adapters.GetAdapter("mock")
	expected := new(mock.MockTasks)

	if reflect.TypeOf(adapter) == reflect.TypeOf(expected) {
		t.Fatalf("adapter is not returning SourceTask catalog")
	}
}
