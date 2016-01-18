package digger

const pathSeparator = "/"

// Digger interface for accesing config properties
type Digger interface {
	GetString(string) (string, error)
	GetNumber(string) (float64, error)
	GetBool(string) (bool, error)
}
