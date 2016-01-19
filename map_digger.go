package digger

import (
	"fmt"
	"reflect"
	"strings"
)

// MapDigger digs in string maps to give you what you look for
type MapDigger struct {
	props map[string]interface{}
}

// NewMapDigger builds a new MapDigger object
func NewMapDigger(props map[string]interface{}) (Digger, error) {
	if props == nil {
		return nil, fmt.Errorf("input map must not be null")
	}
	return &MapDigger{props: props}, nil
}

// GetString digs in and brings you a string (or an error if the path doesn't lead to one)
func (d MapDigger) GetString(path string) (string, error) {
	// split the path
	nestedKeys := strings.Split(path, pathSeparator)
	// just the last element
	lastProp := nestedKeys[len(nestedKeys)-1]
	// remove it
	nestedKeys = nestedKeys[:len(nestedKeys)-1]
	// get last map
	lastMap, err := d.dig(nestedKeys)
	if err != nil {
		return "", err
	}
	target := lastMap[lastProp]
	if target == nil {
		return "", fmt.Errorf("Incorrect path (%s) : %s property does not exist or has nil value", path, lastProp)
	}
	switch target.(type) {
	case string:
		return target.(string), nil
	default:
		return "", fmt.Errorf("Incorrect type: %s property is not a string but a %v", path, reflect.TypeOf(target))
	}
}

// GetNumber digs in and brings you a number (or an error if the path doesn't lead to one)
func (d MapDigger) GetNumber(path string) (float64, error) {
	// split the path
	nestedKeys := strings.Split(path, pathSeparator)
	// just the last element
	lastProp := nestedKeys[len(nestedKeys)-1]
	// remove it
	nestedKeys = nestedKeys[:len(nestedKeys)-1]
	// get last map
	lastMap, err := d.dig(nestedKeys)
	if err != nil {
		return 0, err
	}
	target := lastMap[lastProp]
	if target == nil {
		return 0, fmt.Errorf("Incorrect path (%s) : %s property does not exist or has nil value", path, lastProp)
	}
	switch target.(type) {
	case float64:
		return target.(float64), nil
	case int:
		return float64(target.(int)), nil
	default:
		return 0, fmt.Errorf("Incorrect type: %s property is not a number but a %v", path, reflect.TypeOf(target))
	}
}

// GetBool digs in and brings you a boolean (or an error if the path doesn't lead to one)
func (d MapDigger) GetBool(path string) (bool, error) {
	// split the path
	nestedKeys := strings.Split(path, pathSeparator)
	// just the last element
	lastProp := nestedKeys[len(nestedKeys)-1]
	// remove it
	nestedKeys = nestedKeys[:len(nestedKeys)-1]
	// get last map
	lastMap, err := d.dig(nestedKeys)
	if err != nil {
		return false, err
	}
	target := lastMap[lastProp]
	if target == nil {
		return false, fmt.Errorf("Incorrect path (%s) : %s property does not exist or has nil value", path, lastProp)
	}
	switch target.(type) {
	case bool:
		return target.(bool), nil
	default:
		return false, fmt.Errorf("Incorrect type: %s property is not a string but a %v", path, reflect.TypeOf(target))
	}
}

func (d MapDigger) dig(path []string) (map[string]interface{}, error) {
	// traverse the map
	current := d.props
	for _, key := range path {
		if current[key] == nil {
			return nil, fmt.Errorf("Incorrect path (%s) : %s property does not exist or has nil value", path, key)
		}

		switch current[key].(type) {
		case map[string]interface{}:
			current = current[key].(map[string]interface{})
		default:
			return nil, fmt.Errorf("Incorrect path (%s) : %s has no nested objects", path, key)
		}
	}

	return current, nil
}
