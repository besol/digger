package digger

import "fmt"

// MultiDigger digs in a set of diggers: ftw!
type MultiDigger struct {
	diggers []Digger
}

// NewMultiDigger builds a new MultiDigger object
func NewMultiDigger(diggers ...Digger) (Digger, error) {
	if diggers == nil {
		return nil, fmt.Errorf("nil argument")
	}
	return MultiDigger{diggers}, nil
}

// GetString digs in and brings you a string (or an error if the path doesn't lead to one)
func (d MultiDigger) GetString(path string) (string, error) {
	for _, digger := range d.diggers {
		val, err := digger.GetString(path)
		if err == nil {
			return val, nil
		}
	}
	return "", fmt.Errorf("Could not read string from '%s'", path)
}

// GetNumber digs in and brings you a number (or an error if the path doesn't lead to one)
func (d MultiDigger) GetNumber(path string) (float64, error) {
	for _, digger := range d.diggers {
		val, err := digger.GetNumber(path)
		if err == nil {
			return val, nil
		}
	}
	return 0, fmt.Errorf("Could not read number from '%s'", path)
}

// GetBool digs in and brings you a boolean (or an error if the path doesn't lead to one)
func (d MultiDigger) GetBool(path string) (bool, error) {
	for _, digger := range d.diggers {
		val, err := digger.GetBool(path)
		if err == nil {
			return val, nil
		}
	}
	return false, fmt.Errorf("Could not read bool from '%s'", path)
}

// Get digs in and brings you anything (or an error if the path doesn't lead to a value)
func (d MultiDigger) Get(path string) (interface{}, error) {
	for _, digger := range d.diggers {
		val, err := digger.Get(path)
		if err == nil {
			return val, nil
		}
	}
	return false, fmt.Errorf("Could not read bool from '%s'", path)
}
