package digger

import (
	"encoding/json"
	"fmt"
)

// NewJSONDigger builds a new Digger object from a JSON stream
func NewJSONDigger(jsonBytes []byte) (Digger, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return NewMapDigger(jsonMap)
}
