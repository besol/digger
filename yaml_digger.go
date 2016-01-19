package digger

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// NewYAMLDigger builds a new Digger object from a JSON stream
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

// normalizeValue will build a map[string]interface{} out of a map[interface{}]interface{}
// (based on https://github.com/moraes/config/blob/master/config.go)
func normalizeValue(value interface{}) (interface{}, error) {
	switch value := value.(type) {
	case map[interface{}]interface{}:
		node := make(map[string]interface{}, len(value))
		for k, v := range value {
			key, ok := k.(string)
			if !ok {
				return nil, fmt.Errorf("Unsupported map key: %#v", k)
			}
			item, err := normalizeValue(v)
			if err != nil {
				return nil, fmt.Errorf("Unsupported map value: %#v", v)
			}
			node[key] = item
		}
		return node, nil
	case bool, float64, int, string:
		return value, nil
	}
	return nil, fmt.Errorf("Unsupported type: %T", value)
}
