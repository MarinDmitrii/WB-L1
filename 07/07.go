package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Реализовать конкурентную запись данных в map.

// Запись в map с использованием Mutex
func WriteMutexMap(key int, value string, mutexMap map[int]string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	// Блокируем Mutex, чтобы остальные горутины не могли писать в map
	mu.Lock()
	// Записываем данные
	mutexMap[key] = value
	// Разблокируем Mutex
	mu.Unlock()
}

// Запись в map с использованием sync.Map
func WriteSyncMap(key int, value string, syncMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	syncMap.Store(key, value)
}

func main() {
	// Объявляем пустую map для примера с использованием Mutex
	mutexMap := make(map[int]string)
	// Объявляем пустую map для примера с использованием sync.Map
	syncMap := sync.Map{}
	// Объявляем Mutex
	mu := sync.Mutex{}
	// Объявляем WaitGroup
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		// Генерируем значение для записи в map
		value := "Num" + strconv.Itoa(i)
		wg.Add(1)
		if i < 25 {
			go WriteMutexMap(i, value, mutexMap, &mu, &wg)
		} else {
			go WriteSyncMap(i, value, &syncMap, &wg)
		}
	}

	// Дожидаемся выполнения всех горутин
	wg.Wait()
	fmt.Println(mutexMap)

	syncMap.Range(func(key, value any) bool {
		fmt.Printf("%v:%v ", key, value)
		return true
	})
}
