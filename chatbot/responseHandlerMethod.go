package chatbot

import (
	"reflect"
	"runtime"

	"github.com/jcgarciaram/boomy/utils"
)

// ResponseHandlerMethod stores the method name used to handle a response from the client
type ResponseHandlerMethod struct {
	MethodName string
}

var (
	methodMap = make(map[string]reflect.Value)
)

// RegisterMethod adds the method to the methodMap of the chatbot to be able to be called later on
func RegisterMethod(method func(utils.Conn, interface{}, string) error) string {
	methodName := runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()
	methodValue := reflect.ValueOf(method)

	methodMap[methodName] = methodValue

	return methodName
}
