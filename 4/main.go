package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// getNumberOfWorkers - функция для получения числа воркеров от пользователя
func getNumberOfWorkers() int {
	fmt.Print("Введите число воркеров: ")

	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Printf("Не удалось считать число воркеров: %v\n", err)
	}

	return n
}

// work - функция для запуска работы воркеров
func work(nWorkers int) {
	msgChan := make(chan string) // Канал для передачи сообщений воркерам
	done := make(chan struct{})  // Канал для завершения работы воркеров
	var wg sync.WaitGroup        // Группа для ожидания завершения работы всех воркеров

	// Запуск воркеров
	for workerID := 0; workerID < nWorkers; workerID++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case msg := <-msgChan:
					fmt.Println(msg) // Выводим полученное сообщение
				case <-done:
					fmt.Printf("Воркер %d завершил работу\n", workerID)
					return
				}
			}
		}(workerID)
	}

	// Отправка сообщений воркерам
	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			default:
				msgChan <- fmt.Sprintf("Сообщение #%d", i) // Генерируем и отправляем сообщение
			}
		}
	}()

	// Обработка сигнала завершения программы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan // Ожидаем сигнал о завершении программы
		close(signalChan)
		close(done) // Закрываем канал завершения работы воркеров
	}()

	wg.Wait()      // Ожидаем завершения работы всех воркеров
	close(msgChan) // Закрываем канал сообщений
}

func main() {
	nWorkers := getNumberOfWorkers() // Получаем число воркеров от пользователя

	work(nWorkers) // Запускаем работу воркеров
}
