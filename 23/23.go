package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	// Объявляем слайс:
	var i int
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	// Задаём элемент, который нужно удалить:
	for {
		fmt.Printf("Укажите номер элемента, который необходимо удалить из среза длиной %v:\n", len(slice))
		fmt.Scan(&i)
		if i >= 0 && i < len(slice) {
			break
		} else {
			continue
		}
	}

	fmt.Println(slice)
	newSlice := append(slice[:i], slice[i+1:]...)
	slice = newSlice

	fmt.Println(slice)
}
