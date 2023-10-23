package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 4
	b := 8
	fmt.Printf("Исходные числа:\n%v и %v\n", a, b)

	a, b = b, a
	fmt.Printf("Числа после перестановки:\n%v и %v", a, b)
}
