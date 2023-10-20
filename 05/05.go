package main

import (
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func writer(run chan<- int, stop <-chan struct{}) {
	value := 1
	for {
		select {
		case <-stop:
			close(run)
			return
		default:
			run <- value
			value++
			time.Sleep(1 * time.Second)
		}
	}
}

func workers(id int, wg *sync.WaitGroup, run <-chan int) {
	defer wg.Done()

	for {
		select {
		case value, ok := <-run:
			if !ok {
				fmt.Printf("Worker %v: Channel closed\n", id)
				return
			}
			fmt.Printf("Worker %v: %v\n", id, value)
		}
	}
}

func main() {
	var n int
	// Задаём время работы программы
	fmt.Print("Enter the program run time in seconds: ")
	fmt.Scan(&n)

	// Создаём буферизованный канал
	run := make(chan int, 5)

	// Создаём канал для остановки записи данных
	stop := make(chan struct{})

	// Реализуем постоянную запись данных в канал run
	go writer(run, stop)

	// Объявляем WaitGroup
	wg := new(sync.WaitGroup)

	// Запуск воркеров
	wg.Add(1)
	go workers(1, wg, run)

	// Отправляем сигнал об остановке в горутину
	time.Sleep(time.Duration(n) * time.Second)
	close(stop)

	// Дожидаемся выполнения всех горутин
	wg.Wait()

	fmt.Println("Программа завершена")
}
