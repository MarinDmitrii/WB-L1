package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров

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
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Создание канала для обработки сигнала завершения программы
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT)

	var n int
	// Задаём количество воркеров
	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&n)

	// Создаём буферизованный канал
	run := make(chan int, n)

	// Создаём канал для остановки записи данных
	stop := make(chan struct{})

	// Реализуем постоянную запись данных в канал run
	go writer(run, stop)

	// Объявляем WaitGroup
	wg := new(sync.WaitGroup)

	// Запуск воркеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		// Передаём счётчик i в горутину, чтобы избежать проблемы считывания из общего контекста
		go workers(i, wg, run)
	}

	// Ожидание сигналов об остановке (ctrl+c)
	<-sig

	// Отправляем сигнал об остановке в горутину
	close(stop)

	// Дожидаемся выполнения всех горутин
	wg.Wait()

	fmt.Println("Программа завершена")
}
