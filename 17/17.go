package main

import (
	"fmt"
)

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(slice []int, num int) int {
	low, high := 0, len(slice)-1

	// Выполняем поиск, пока диапазон не сократится до единичного значения
	for low <= high {
		// Находим средний индекс
		mid := low + (high-low)/2

		// Если значение найдено, то возвращаем индекс
		if num == slice[mid] {
			return mid
		} else if num > slice[mid] {
			// Переходим к правой половине
			low = mid + 1
		} else if num < slice[mid] {
			// Переходим к левой половине
			high = mid - 1
		}
	}

	// Значение не найдено
	return -1
}

func main() {
	// Объявляем срез:
	slice := make([]int, 10)
	// Заполняем срез:
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	// Задаём число для поиска
	num := 12
	fmt.Println(slice)

	idx := binarySearch(slice, num)
	if idx != -1 {
		fmt.Printf("Element %v found at index %v\n", num, idx)
	} else {
		fmt.Printf("Element %v not found in slice\n", num)
	}
}
