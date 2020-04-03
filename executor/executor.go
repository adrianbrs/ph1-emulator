package executor

import (
	"reflect"
	"strings"
)

func ExecuteOperation(obj interface{}, name string, args ...interface{}) bool {
	method := reflect.ValueOf(obj).MethodByName("ExecuteOp" + strings.ToUpper(name))

	if !method.IsValid() {
		return false
	}

	operationArgs := make([]reflect.Value, method.Type().NumIn())

	for i := 0; i < method.Type().NumIn(); i++ {
		operationArgs[i] = reflect.ValueOf(args[i])
	}

	method.Call(operationArgs)

	return true
}
