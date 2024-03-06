package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 interface{} = 42
	var v2 interface{} = "hello"
	var v3 interface{} = true
	var v4 interface{} = make(chan int)

	checkType(v1)
	checkType(v2)
	checkType(v3)
	checkType(v4)
}

// checkType - функция для проверки типа
func checkType(v interface{}) {
	t := reflect.TypeOf(v) // Получаем тип переменной v
	switch t.Kind() {
	case reflect.Int: // Тип переменной v - int
		fmt.Printf("%v - это int\n", v)
	case reflect.String: // Тип переменной v - string
		fmt.Printf("%v - это string\n", v)
	case reflect.Bool: // Тип переменной v - bool
		fmt.Printf("%v - это bool\n", v)
	case reflect.Chan: // Тип переменной v - chan
		fmt.Printf("%v - это chan\n", v)
	default: // Неизвестный тип
		fmt.Println("Неизвестный тип")
	}
}
