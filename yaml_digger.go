package digger

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// NewYAMLDigger builds a new Digger object from a YAML stream
func NewYAMLDigger(yamlBytes []byte) (Digger, error) {
	var yamlMap interface{}
	err := yaml.Unmarshal(yamlBytes, &yamlMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %v", err)
	}
	normalizedMap, err := normalizeValue(yamlMap)
	if err != nil {
		return nil, fmt.Errorf("error normalizing YAML map: %v", err)
	}
	return NewMapDigger(normalizedMap.(map[string]interface{}))
}

// NewYAMLDiggerFromFile builds a new Digger object from a YAML file
func NewYAMLDiggerFromFile(yamlFilePath string) (Digger, error) {
	yamlBytes, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file '%s' : %v", yamlFilePath, err)
	}
	return NewYAMLDigger(yamlBytes)
}
