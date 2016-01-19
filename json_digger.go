package digger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// NewJSONDigger builds a new Digger object from a JSON stream
func NewJSONDiggerFromFile(jsonFilePath string) (Digger, error) {
	jsonBytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file '%s' : %v", jsonFilePath, err)
	}
	return NewJSONDigger(jsonBytes)
}
