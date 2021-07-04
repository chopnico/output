package output

import (
	"reflect"
)

func value(t reflect.Value) reflect.Value {
	switch t.Kind() {
	case reflect.Ptr:
		return t.Elem()
	default:
		return t
	}
}
