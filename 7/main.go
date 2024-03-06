package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мапу для хранения данных
	data := make(map[string]int)

	// Создаем мьютекс для безопасного доступа к мапе
	var mutex sync.Mutex

	// Количество горутин для записи данных
	numWorkers := 5

	// Канал для сигнала завершения работы горутин
	done := make(chan struct{})

	// Запускаем горутины для записи данных
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer func() {
				done <- struct{}{} // Отправляем сигнал о завершении работы горутины
			}()

			// Захватываем мьютекс перед записью в мапу
			mutex.Lock()
			// Записываем данные в мапу
			data[fmt.Sprintf("key%d", workerID)] = workerID
			// Освобождаем мьютекс после записи
			mutex.Unlock()
		}(i)
	}

	// Ждем завершения работы всех горутин
	for i := 0; i < numWorkers; i++ {
		<-done
	}

	// Выводим данные из мапы
	fmt.Println("Мапа после записи данных:")
	fmt.Println(data)
}
