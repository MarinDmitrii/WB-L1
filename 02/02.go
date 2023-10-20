package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	Solution(array)
}

func Solution(array [5]int) {
	// Объявляем WaitGroup
	wg := new(sync.WaitGroup)

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		// Передаём счётчик i в горутину, чтобы избежать проблемы считывания из общего контекста
		go func(i int) {
			// После завершения горутины передаём информацию об этом путём снижения счётчика WaitGroup
			defer wg.Done()
			result := array[i] * array[i]
			fmt.Printf("Квадрат числа %v - %v\n", array[i], result)
		}(i)
	}
	// Дожидаемся выполнения всех горутин
	wg.Wait()
}

// func ChannelSolution(array [5]int) {
// 	// Объявляем WaitGroup
// 	wg := new(sync.WaitGroup)
// 	// Объявляем канал
// 	resultCh := make(chan int)

// 	for _, value := range array {
// 		wg.Add(1)
// 		go func(value int) {
// 			defer wg.Done()
// 			result := value * value
// 			resultCh <- result
// 		}(value)
// 	}
// 	wg.Wait()
// 	fmt.Println("it's ok")
// 	close(resultCh)

// 	for res := range resultCh {
// 		fmt.Printf("Квадрат числа - %v\n", res)
// 	}

// }
