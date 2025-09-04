package main

import (
	"fmt"
	"reflect"
)

func EvaluateType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип int")
	case bool:
		fmt.Println("Тип bool")
	case string:
		fmt.Println("Тип string")
	default:
		if reflect.ValueOf(v).Kind() == reflect.Chan {
			fmt.Println("Тип chan")
		} else {
			fmt.Println("Тип не опознан")
		}
	}
}

func main() {
	EvaluateType(123)

	EvaluateType("string")

	EvaluateType(true)

	ch := make(chan int)
	EvaluateType(ch)
}
