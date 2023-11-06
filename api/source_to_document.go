package api

import (
	"encoding/json"
	"log"
)

type AdapterSpec struct {
	Adapter string
	Options map[string]string
}

type TaskSpec struct {
	Spec map[string]AdapterSpec
}

type JSONSpecHandler struct{}

func (h JSONSpecHandler) PayloadToStruct(content string) TaskSpec {
	spec := new(TaskSpec)
	err := json.Unmarshal([]byte(content), spec)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return *spec
}

// func (h JSONSpecHandler) LoadJson(path string) {
// 	content, err := io.ReadFile(path)

// 	if err != nil {
// 		log.Fatal("Error when opening file: ", err)
// 	}

// 	// parse json
// 	json.Unmarshal(content)
// }
