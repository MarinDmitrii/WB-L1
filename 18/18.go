package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Структура Counter представляет счетчик с методом для инкрементации
type Counter struct {
	value int
	mu    sync.Mutex
}

// Метод Increment увеличивает значение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	// Объявляем экземпляр счётчика
	c := Counter{}
	wg := new(sync.WaitGroup)

	// Создаем горутины для инкрементации счетчика в конкурентной среде
	for i := 0; i < 274296; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	// Ожидаем завершения всех горутин:
	wg.Wait()
	// Выводим результат
	fmt.Println(c.value)
}
