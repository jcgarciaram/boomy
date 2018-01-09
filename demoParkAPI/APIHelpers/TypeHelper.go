package APIHelpers

// TypeHelper will be used to be able to use shared common functions across objects
type TypeHelper interface {
	Save() error
	Get(string) error
	Validate() error
}

// SliceTypeHelper will be used to be able to use shared common functions across slice objects
type SliceTypeHelper interface {
	GetAll() error
}
