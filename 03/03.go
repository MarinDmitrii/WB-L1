package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	slice := []int{2, 4, 6, 8, 10}

	Solution(slice)
}

func Solution(slice []int) {
	// Объявляем WaitGroup
	wg := new(sync.WaitGroup)
	sum := 0

	for i := 0; i < len(slice); i++ {
		wg.Add(1)
		// Передаём счётчик i в горутину, чтобы избежать проблемы считывания из общего контекста
		go func(i int) {
			// После завершения горутины передаём информацию об этом путём снижения счётчика WaitGroup
			defer wg.Done()
			result := slice[i] * slice[i]
			sum += result
		}(i)
	}
	// Дожидаемся выполнения всех горутин
	wg.Wait()
	fmt.Printf("Сумма квадратов последовательности %v = %v\n", slice, sum)
}
