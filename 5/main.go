package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// getNSeconds - функция для получения времени от пользователя
func getNSeconds() int {
	fmt.Print("Введите время через которое воркеры должны завершить работу: ")

	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Printf("Не удалось считать время: %v\n", err)
	}

	return n
}

// work - функция для запуска работы воркеров
func work(t int) {
	msgChan := make(chan string) // Канал для передачи сообщений воркерам
	done := make(chan struct{})  // Канал для завершения работы воркеров
	var wg sync.WaitGroup        // Группа для ожидания завершения работы всех воркеров

	// Запуск воркера
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg := <-msgChan:
				fmt.Println(msg) // Выводим полученное сообщение
			case <-done:
				fmt.Println("Воркер завершил работу")
				return
			}
		}
	}()

	// Отправка сообщений воркеру и отсчет времени
	go func(t int) {
		timePoint := time.Now().Add(time.Duration(t) * time.Second) // Определение времени завершения работы воркеров
		for {
			select {
			case <-done:
				return
			default:
				remainingSeconds := time.Until(timePoint).Seconds() // Определение оставшегося времени
				if remainingSeconds <= 0 {
					close(done) // Закрываем канал завершения работы воркеров
					return
				}
				msgChan <- fmt.Sprintf("Секунд осталось %.0f", remainingSeconds) // Отправляем сообщение с оставшимся временем
				time.Sleep(1 * time.Second)                                      // Задержка на одну секунду
			}
		}
	}(t)

	// Обработка сигнала завершения программы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan // Ожидаем сигнал о завершении программы
		close(signalChan)
		close(done) // Закрываем канал завершения работы воркеров
	}()

	wg.Wait()      // Ожидаем завершения работы воркера
	close(msgChan) // Закрываем канал сообщений
}

func main() {
	nSeconds := getNSeconds() // Получаем время от пользователя

	work(nSeconds) // Запускаем работу воркеров
}
