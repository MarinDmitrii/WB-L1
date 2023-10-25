package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// Вариант 1:
func search1(newStr string) bool {
	for _, char := range newStr {
		// Считаем сколько раз символ встречается в строке:
		if strings.Count(newStr, string(rune(char))) > 1 {
			// Если символ встречается более одного раза, то возвращаем false
			return false
		}
	}

	return true
}

// Вариант 2:
func search2(newStr string) bool {
	// Объявляем map для отслеживания уникальных символов:
	unique := make(map[rune]bool)

	for _, char := range newStr {
		// Если символ уже есть в map, значит он не уникален
		if unique[char] {
			return false
		}

		// Если символ отсутствует, то добавляем его
		unique[char] = true
	}

	// Если цикл завершился, и нет повторяющихся символов, то возвращаем true
	return true
}

func main() {
	// Объявляем строку:
	str := "abcd"
	// Преобразуем строку в нижний регистр:
	newStr := strings.ToLower(str)

	result1 := search1(newStr)
	fmt.Printf("%v - %v\n", str, result1)

	result2 := search2(newStr)
	fmt.Printf("%v - %v\n", str, result2)
}
