package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

func typeswitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Printf("Type of %v is int\n", value)
	case string:
		fmt.Printf("Type of %v is string\n", value)
	case bool:
		fmt.Printf("Type of %v is bool\n", value)
	case chan struct{}:
		fmt.Printf("Type of %v is chan struct{}\n", value)
	default:
		fmt.Println("Unidentified type")
	}
}

func typeof(value interface{}) {
	fmt.Printf("Type of %v is %v\n", value, reflect.TypeOf(value))
}

func main() {
	var number, str, boolean, channel interface{}
	number = 42
	str = "Hello, Gopher!"
	boolean = true
	channel = make(chan struct{})

	array := []interface{}{number, str, boolean, channel}

	for _, value := range array {
		// Вариант через type switch
		typeswitch(value)
		// Вариант через reflect.TypeOf()
		typeof(value)
	}
}
