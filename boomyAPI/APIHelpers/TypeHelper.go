package APIHelpers

import "github.com/jcgarciaram/boomy/utils"

// TypeHelper will be used to be able to use shared common functions across objects
type TypeHelper interface {
	Create(utils.Conn) error
	Update(utils.Conn) error
	First(utils.Conn, uint) error
	Validate() error
}

// SliceTypeHelper will be used to be able to use shared common functions across slice objects
type SliceTypeHelper interface {
	Find(utils.Conn) error
}
