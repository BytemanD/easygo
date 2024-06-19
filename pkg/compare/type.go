package compare

import (
	"reflect"
)

func IsType[T any](i interface{}) bool {
	switch i.(type) {
	case T:
		return true
	default:
		return false
	}
}

func IsStructOf(structName string, v1 interface{}) bool {
	if v1 == nil {
		return false
	}
	if reflect.ValueOf(v1).Type().Name() == structName {
		return true
	}
	return IsPtrOf(structName, v1)
}

func IsPtrOf(structName string, v1 interface{}) bool {
	if v1 == nil {
		return false
	}
	v1Value := reflect.ValueOf(v1)
	if v1Value.Kind() != reflect.Ptr {
		return false
	}
	v1ElemType := v1Value.Elem().Type()
	if v1ElemType.Kind() != reflect.Struct {
		return false
	}
	return v1ElemType.Name() == structName
}

func EqualsStruct(v1 interface{}, v2 interface{}) bool {
	if v1 == nil {
		return v2 == nil
	}

	v1Value := reflect.ValueOf(v1)
	var v1TypeName string
	if v1Value.Kind() == reflect.Ptr {
		v1TypeName = v1Value.Elem().Type().Name()
	} else {
		v1TypeName = reflect.ValueOf(v1).Type().Name()
	}

	return IsStructOf(v1TypeName, v2)
}
