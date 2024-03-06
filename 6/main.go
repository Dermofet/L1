package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Пример использования WaitGroup
	wgExample()

	// Пример использования канала для завершения горутин
	chanExample()

	// Пример использования контекста с отменой
	contextWithCancelExample()

	// Пример использования контекста с таймаутом
	contextWithTimeoutExample()

	// Пример использования runtime.Goexit
	goexitExample()
}

// Пример использования WaitGroup
func wgExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer fmt.Println("Горутина завершена по ожиданию WaitGroup")
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина в работе")
			time.Sleep(time.Second)
		}
	}(&wg)
	time.Sleep(2 * time.Second) // Некоторая задержка
	wg.Done()                   // Сигнализируем горутине о завершении работы
	wg.Wait()                   // Ожидаем завершения горутины
}

// Пример использования канала для завершения горутин
func chanExample() {
	stopChan := make(chan struct{})
	go func(stopChan <-chan struct{}) {
		defer fmt.Println("Горутина завершена по сигналу из канала")
		for {
			select {
			case <-stopChan:
				return
			default:
				fmt.Println("Горутина в работе")
				time.Sleep(time.Second)
			}
		}
	}(stopChan)
	time.Sleep(2 * time.Second) // Некоторая задержка
	close(stopChan)             // Закрываем канал для завершения работы горутины
}

// Пример использования контекста с отменой
func contextWithCancelExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		defer fmt.Println("Горутина завершена по сигналу отмены контекста")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина в работе")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	time.Sleep(2 * time.Second) // Некоторая задержка
	cancel()                    // Отменяем выполнение по контексту с отменой
}

// Пример использования контекста с таймаутом
func contextWithTimeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		defer fmt.Println("Горутина завершена по сигналу отмены контекста")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина в работе")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	time.Sleep(2 * time.Second) // Некоторая задержка
}

// Пример использования runtime.Goexit
func goexitExample() {
	go workerWithGoexit()
	time.Sleep(2 * time.Second) // Некоторая задержка
}

func workerWithGoexit() {
	defer fmt.Println("Горутина завершена по вызову goexit")
	for {
		fmt.Println("Горутина в работе")
		time.Sleep(time.Second)
		runtime.Goexit()
	}
}
