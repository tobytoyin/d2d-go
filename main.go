package main

import (
	adapters "d2d/adapters/mock"
	"d2d/contracts"
)

func main() {

	source := &contracts.Source{"./examples/sample_text.txt", "mock", nil}

	task := new(adapters.MockTasks)
	buffer := task.SourceIO(source)

	// buffer := make([]byte, 12)
	// n, _ := reader.Read(buffer)

	println(string(string(buffer)))

	summary := task.Summary(source)
	println(summary)

}
