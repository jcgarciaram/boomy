package utils

import (
	"reflect"
)

// GetType is used to get the name of a struct
func GetType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}
