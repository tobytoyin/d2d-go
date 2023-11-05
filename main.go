package main

import (
	"d2d/adapters"
	"d2d/contracts"
)

func main() {
	source := &contracts.Source{"./sample_text.txt", "mock", nil}

	task := new(adapters.MockTasks)
	buffer := task.SourceIO(source)

	// buffer := make([]byte, 12)
	// n, _ := reader.Read(buffer)

	println(string(string(buffer)))

}
