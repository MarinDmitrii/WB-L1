package main

import (
	"fmt"
	"math"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func XOR(a int64, i int) {
	mask := int64(1) << i
	result := a ^ mask

	fmt.Println(result)
}

func main() {
	var s string
	var a int64
	var i int

	for {
		fmt.Println("Введите значение переменной int64:")
		fmt.Scan(&s)

		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Println("Введено нечисловое значение или значение вне допустимого диапазона int64")
			continue
		}

		a = num
		break
	}

	for {
		fmt.Println("Введите номер бита, который нужно изменить:")
		_, err := fmt.Scan(&i)
		if err != nil {
			fmt.Println("Введено нечисловое значение")
			continue
		}

		if i < 0 || i > 63 {
			fmt.Println("Индекс бита вне допустимого диапазона int64")
		} else {
			break
		}
	}

	if a == math.MinInt64 && i == 63 {
		fmt.Println("Операция не выполнена, так как приведёт к выходу из диапазона int64")
	} else {
		XOR(a, i)
	}
}
