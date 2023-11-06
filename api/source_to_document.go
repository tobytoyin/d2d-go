package api

import (
	"encoding/json"
	"log"
	"os"
)

// takes in the string content of JSON and parse into TaskSpec
func JsonStringToTaskSpec(content string) TaskSpec {
	spec := new(TaskSpec)
	err := json.Unmarshal([]byte(content), spec)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return *spec
}

// takes in the JSON file and parse into TaskSpec
func JsonFileToTaskSpec(path string) TaskSpec {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("%s is not a valid JSON path", path)
	}

	taskSpec := JsonStringToTaskSpec(string(content))
	return taskSpec
}

type AdapterSpec struct {
	Adapter string
	Options map[string]string
}

type TaskSpec struct {
	Spec map[string]AdapterSpec
}
