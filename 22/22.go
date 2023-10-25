package main

import (
	"fmt"
	"math"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	var o string

	for {
		fmt.Println("Укажите символ операции, которую необходимо выполнить: +, -, *, /")
		fmt.Scan(&o)

		if o == "+" || o == "-" || o == "*" || o == "/" {
			break
		} else {
			fmt.Println("Введён недопустимый символ")
		}
	}

	// Вариант 1:
	a := math.Pow(2, 65)
	b := math.Pow(2, 37)

	switch o {
	case "+":
		fmt.Printf("%v + %v = %v\n", a, b, a+b)
	case "-":
		fmt.Printf("%v - %v = %v\n", a, b, a-b)
	case "*":
		fmt.Printf("%v * %v = %v\n", a, b, a*b)
	case "/":
		fmt.Printf("%v / %v = %v\n", a, b, a/b)
	}

	// Вариант 2:
	// Устанавливаем значения a и b (здесь используются строки для больших чисел)
	stringA := "36893488147419103232" // 2 ^ 65
	stringB := "137438953472"         // 2 ^ 37

	a2 := new(big.Int)
	b2 := new(big.Int)

	a2.SetString(stringA, 10)
	b2.SetString(stringB, 10)

	result := new(big.Int)

	switch o {
	case "+":
		fmt.Printf("%v + %v = %v\n", a2, b2, result.Add(a2, b2))
	case "-":
		fmt.Printf("%v - %v = %v\n", a2, b2, result.Sub(a2, b2))
	case "*":
		fmt.Printf("%v * %v = %v\n", a2, b2, result.Mul(a2, b2))
	case "/":
		fmt.Printf("%v / %v = %v\n", a2, b2, result.Div(a2, b2))
	}
}
