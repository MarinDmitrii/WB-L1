package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	// Объявляем строку
	str := "главрыба"
	// Преобразовываем строку в массив рун
	runes := []rune(str)

	// Переворачиваем массив рун
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}

	fmt.Printf("%v - %v\n", str, string(runes))
}
