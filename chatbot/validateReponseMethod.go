package chatbot

import (
	"reflect"

	"github.com/jcgarciaram/demoPark/utils"
)

type validateResponseMethod struct {
	typeName   string
	methodName string
}

var (
	typeMap = make(map[string]reflect.Value)
)

// RegisterType adds the type to the typeMap of the chatbot which will be used later on
func RegisterType(o interface{}) {
	typeName, isPointer := utils.GetTypePointer(o)

	if isPointer {
		typeMap[typeName] = reflect.ValueOf(o)
		return
	}

	typeMap[typeName] = reflect.ValueOf(&o)
}
