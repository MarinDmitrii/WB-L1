package main

import (
	"fmt"
	"math/rand"
)

// Реализовать пересечение двух неупорядоченных множеств.

func mix(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Перебираем элементы первого множества:
	for elem := range set1 {
		// Если элемент присутствует во втором множестве, то добавляем его в результат:
		if set2[elem] {
			result[elem] = true
		}
	}

	return result
}

func randomBool() bool {
	random := rand.Intn(2)

	return random == 1
}

func main() {
	// Создаём два неупорядоченных множества:
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// Наполняем неупорядоченные множества:
	for i := 0; i < 30; i++ {
		value := randomBool()
		if value {
			set1[i] = value
		}
	}
	for i := 0; i < 30; i++ {
		value := randomBool()
		if value {
			set2[i] = value
		}
	}

	// Вызываем функцию для нахождения пересечения множеств:
	result := mix(set1, set2)

	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(result)
}
