package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	mu    sync.Mutex     // Мьютекс для обеспечения безопасности доступа к счетчику
	value int            // Значение счетчика
	wg    sync.WaitGroup // WaitGroup для ожидания завершения всех горутин
}

// NewCounter создает новый экземпляр счетчика
func NewCounter() *Counter {
	return &Counter{}
}

// Increment инкрементирует значение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()         // Блокируем доступ к счетчику
	defer c.mu.Unlock() // Освобождаем блокировку после завершения функции
	c.value++           // Инкрементируем значение счетчика
}

// GetValue возвращает текущее значение счетчика
func (c *Counter) GetValue() int {
	c.mu.Lock()         // Блокируем доступ к счетчику
	defer c.mu.Unlock() // Освобождаем блокировку после завершения функции
	return c.value      // Возвращаем значение счетчика
}

func main() {
	// Создаем новый счетчик
	counter := NewCounter()

	// Количество горутин для инкрементации счетчика
	numGoroutines := 100

	// Добавляем в WaitGroup количество горутин, которые будут инкрементировать счетчик
	counter.wg.Add(numGoroutines)

	// Запускаем горутины для инкрементации счетчика
	for i := 0; i < numGoroutines; i++ {
		go func() {
			counter.Increment() // Инкрементируем счетчик
			counter.wg.Done()   // Уменьшаем счетчик WaitGroup после завершения работы горутины
		}()
	}

	// Ожидаем завершения всех горутин
	counter.wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
