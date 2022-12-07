package main

import (
	"fmt"
	"reflect"
)

type Field[T any] struct {
	Value T
}

func typeOf(v any) reflect.Type {
	return reflect.TypeOf(&v).Elem()
}

func genericTypeOf[T any](v T) reflect.Type {
	return reflect.TypeOf(&v).Elem()
}

func main() {
	myVar := Field[string]{
		Value: "string",
	}
	fmt.Println(typeOf(myVar))
	fmt.Println(genericTypeOf(myVar))
}
