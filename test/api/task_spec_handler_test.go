package api

import (
	"d2d/api"
	"testing"
)

func TestJsonToStructSingleNoOptions(t *testing.T) {
	samplePayload := `
	{
		"spec": {
			"summary": {
				"adapter": "mock"
			}
		}
	}
	`
	res := api.JsonStringToTaskSpec(samplePayload)

	if res.Spec == nil {
		t.Fatalf("spec is incorrect")
	}

	if res.Spec["summary"].Adapter != "mock" {
		t.Fatalf("adapter key is incorrect")
	}

	if res.Spec["summary"].Options != nil {
		t.Fatalf("options key is incorrect")
	}
}

func TestJsonToStructSingleWithOptions(t *testing.T) {
	samplePayload := `
	{
		"spec": {
			"summary": {
				"adapter": "mock",
				"options": {
					"arg1": "value1"
				}
			}
		}
	}
	`

	res := api.JsonStringToTaskSpec(samplePayload)

	// access and test the "options" key
	if res.Spec["summary"].Options["arg1"] != "value1" {
		t.Fatalf("options.arg1 value is incorrect")
	}

}

func TestJsonToStructMixture(t *testing.T) {
	samplePayload := `
	{
		"spec": {
			"summary": {
				"adapter": "mock"
			},
			"metadata": {
				"adapter": "mock", 
				"options": {
					"arg1": "value1",
					"arg2": "value2"
				}
			}
		}
	}
	`
	res := api.JsonStringToTaskSpec(samplePayload)

	if res.Spec["summary"].Adapter != "mock" {
		t.Fatalf("summary key isn't loading correctly")
	}

	if res.Spec["summary"].Options != nil {
		t.Fatalf("summary options isn't loading correctly")
	}

	if res.Spec["metadata"].Adapter != "mock" {
		t.Fatalf("metadata key isn't loading correctly")
	}

	if res.Spec["metadata"].Options == nil {
		t.Fatalf("metadata options isn't loading correctly")
	}

	if res.Spec["metadata"].Options["arg1"] != "value1" {
		t.Fatalf("metadata options elements isn't loading correctly")
	}

}

func TestJsonFileToTaskSpec(t *testing.T) {
	path := "../../examples/task_spec.json"

	res := api.JsonFileToTaskSpec(path)

	// same tests as TestJsonToStructMixture
	if res.Spec["summary"].Adapter != "mock" {
		t.Fatalf("summary key isn't loading correctly")
	}

	if res.Spec["summary"].Options != nil {
		t.Fatalf("summary options isn't loading correctly")
	}

	if res.Spec["metadata"].Adapter != "mock" {
		t.Fatalf("metadata key isn't loading correctly")
	}

	if res.Spec["metadata"].Options == nil {
		t.Fatalf("metadata options isn't loading correctly")
	}

	if res.Spec["metadata"].Options["arg1"] != "value1" {
		t.Fatalf("metadata options elements isn't loading correctly")
	}
}
