package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

// Вариант 1:
func reverse1(str string) {
	// Преобразуем строку в руны:
	runes := []rune(str)
	var newRunes []rune

	// Меняем слова местами:
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == ' ' {
			newRunes = append(newRunes, runes[i+1:]...)
			newRunes = append(newRunes, runes[i])
			runes = runes[:i]
		}
	}
	newRunes = append(newRunes, runes...)

	fmt.Printf("%v - %v\n", str, string(newRunes))
}

// Вариант 2:
func reverse2(str string) {
	// Делим строку на слова:
	words := strings.Split(str, " ")

	// Меняем слова местами:
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем массив строк:
	newStr := strings.Join(words, " ")

	fmt.Printf("%v - %v\n", str, newStr)
}

func main() {
	str := "snow dog sun"

	reverse1(str)
	reverse2(str)
}
