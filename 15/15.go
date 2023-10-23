package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Рекомендуется избегать использования глобальных переменных
// justString - строка размером 100 байт. Так как символы могут занимать разное
// количество байт в зависимости от кодировки, то можем потерять последний символ.

// Пример корректной реалзиации:
func createHugeString(i int) string {
	return strings.Repeat("мg", i)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Преобразуем строку в руны, чтобы не потерять символ
	// В таком случае будет создана строка в 100 символов, а не 100 байт
	runes := []rune(v)
	justString := runes[:100]

	fmt.Println(string(justString))
}

func main() {
	someFunc()
}
