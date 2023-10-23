package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree),
// создать для нее собственное множество.

func main() {
	// Объявляем последовательность строк:
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	// Объявляем множество:
	set := make(map[string]bool)

	// Перебираем элементы массива:
	for _, value := range array {
		// Если элемент отсутствует во множестве, то добавляем его:
		if !set[value] {
			set[value] = true
		}
	}

	fmt.Println(set)
}
