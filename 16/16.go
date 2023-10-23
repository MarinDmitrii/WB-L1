package main

import (
	"fmt"
	"math/rand"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(slice []int) []int {
	// Если длина среза меньше 2, то он уже отсортирован
	if len(slice) < 2 {
		return slice
	}

	// Выбираем средний элемент
	middleIndex := len(slice) / 2
	middle := slice[middleIndex]

	// Удаляем средний элемент из среза:
	slice = append(slice[:middleIndex], slice[middleIndex+1:]...)

	// Делим срез на 2 части:
	// lesser - для элементов меньше или равных среднему элементу,
	// greater - для элементов больше среднего элемента
	var lesser, greater []int

	for _, value := range slice {
		if value <= middle {
			lesser = append(lesser, value)
		} else {
			greater = append(greater, value)
		}
	}

	// Рекурсивно сортируем срезы lesser и greater:
	lesser = quickSort(lesser)
	greater = quickSort(greater)

	// Объединяем отсортированные срезы со средним элементом и возвращаем результат:
	return append(append(lesser, middle), greater...)
}

func main() {
	// Объявляем срез:
	slice := make([]int, 10)
	// Заполняем срез:
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(20)
	}
	// Исходный срез:
	fmt.Println(slice)
	// Сортированный срез:
	fmt.Println(quickSort(slice))
}
