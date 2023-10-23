package main

import (
	"fmt"
	"sync"
	"time"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

func main() {
	// Объявляем массив
	array := [100]int{}
	// Заполняем массив
	for i := 0; i < 100; i++ {
		array[i] = i + 1
	}

	// Создаём каналы:
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)

	// Объявляем WaitGroup
	wg := new(sync.WaitGroup)

	wg.Add(1)
	// Записываем числа из массива в первый канал
	go func() {
		defer wg.Done()
		for _, value := range array {
			chan1 <- value
			time.Sleep(100 * time.Millisecond)
		}
		// Закрываем chan1, чтобы указать, что больше данных не будет
		close(chan1)
	}()

	wg.Add(1)
	// Горутина для считывания данных из первого канала и записи во второй
	go func() {
		defer wg.Done()
		for data := range chan1 {
			chan2 <- data * 2
		}
		close(chan2)
	}()

	wg.Add(1)
	// Считываем данные из второго канала и выводим в stdout
	go func() {
		defer wg.Done()
		for result := range chan2 {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
