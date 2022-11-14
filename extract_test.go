package gojsonschema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToto(t *testing.T) {
	schemaLoader := NewStringLoader(`
	{
		"type":"object",
		"properties": {
			"id": {"type":"string", "format":"uuid", "extract":"flavorId"},
			"list": {"type":"array", "extract":"flavorList"}
		}
	}
	`)

	documentLoader := NewStringLoader(`
	{
		"id":"f5424aa4-cfbd-4ccb-bd42-289e0426806d",
		"list": [
			"a","b",42
		]
	}
	`)

	result, err := Validate(schemaLoader, documentLoader)
	assert.Nil(t, err)

	fmt.Printf("extraction %d: %+v", len(result.Extracts), result.Extracts)

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		t.Errorf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			t.Errorf("- %s\n", desc)
		}
	}
}
