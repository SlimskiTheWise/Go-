package reflect

import (
	"reflect"
)

type Person struct {
	name string
	age  int
}

func ReflectT(args interface{}) reflect.Type {
	return reflect.TypeOf(args)
}

func GetValueOfPointer(num *int) interface{} {
	return reflect.ValueOf(num).Elem().Int()
}
