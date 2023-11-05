package adapters

import (
	mock "d2d/adapters/mock"
	"d2d/contracts"
)

// adapter Mapper
func GetAdapter(adapterName string) *contracts.SourceTask {

	adapterMap := map[string]contracts.SourceTask{
		"mock": new(mock.MockTasks),
	}

	adapter := adapterMap[adapterName]
	return &adapter
}
