package chatbot

import (
	"reflect"
	"runtime"
)

type responseHandlerMethod struct {
	methodName string
}

var (
	methodMap = make(map[string]reflect.Value)
)

// RegisterMethod adds the method to the methodMap of the chatbot to be able to be called later on
func RegisterMethod(method func(interface{}, string) error) {
	methodName := runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()

	methodMap[methodName] = reflect.ValueOf(method)
}
