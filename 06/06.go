package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	// Способ 1: channel
	// Создаём канал для завершения работы горутин:
	stop := make(chan struct{})

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				fmt.Println("Goroutine 1 are stopped")
				return
			default:
				fmt.Println("Goroutine 1 are working")
				time.Sleep(1 * time.Second)
			}
		}
	}(stop)

	time.Sleep(2 * time.Second)
	// Отправляем сигнал об остановке в горутину
	close(stop)

	// Способ 2: context
	// Создаём контекст для прерывания выполнения горутин
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():

			default:
				fmt.Println("Goroutine 2 are working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	// Отправляем сигнал об остановке в горутину
	cancel()
	fmt.Println("Goroutine 2 are stopped")

	// Способ 3: sync.WaitGroup
	// Объявляем WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		fmt.Println("Goroutine 3 are working")
		wg.Done()
	}()

	// Дожидаемся выполнения горутины
	wg.Wait()
	fmt.Println("Goruotine 3 are stopped")

	// Способ 4: runtime.Goexit()
	go func() {
		fmt.Println("Goroutine 4 are working")
		runtime.Goexit()
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Goruotine 4 are stopped")
}
