package main

import (
	"fmt"
	"sync"
)

// Функция воркера для вычисления квадрата числа
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик группы синхронизации по завершении работы воркера
	for n := range jobs {
		fmt.Printf("Воркер %d вычисляет квадрат числа %d\n", id, n)
		results <- n * n // Отправляем результат вычисления квадрата числа в канал результатов
	}
}

// Функция для вычисления квадратов чисел с помощью воркеров
func calculateSquaresWithWorkers(massive []int, numWorkers int) {
	var wg sync.WaitGroup // Создаем группу синхронизации
	wg.Add(numWorkers)    // Увеличиваем счетчик группы синхронизации на количество воркеров

	// Создаем каналы для передачи задач и результатов
	jobs := make(chan int, len(massive))
	results := make(chan int, len(massive))

	// Запускаем воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results, &wg)
	}

	// Отправляем задачи в канал
	for _, n := range massive {
		jobs <- n
	}
	close(jobs) // Закрываем канал задач после отправки всех задач

	// Ждем завершения работы воркеров
	go func() {
		wg.Wait()      // Ждем завершения всех воркеров
		close(results) // Закрываем канал результатов после завершения работы всех воркеров
	}()

	// Получаем результаты и выводим их
	for res := range results {
		fmt.Println(res) // Выводим результаты вычислений квадратов чисел
	}
}

// Функция для вычисления квадратов чисел через sync.WaitGroup
func calculateSquaresWithWaitGroup(massive []int) {
	var wg sync.WaitGroup
	wg.Add(len(massive)) // Увеличиваем счетчик группы синхронизации на количество задач

	// Запускаем воркеров
	for _, n := range massive {
		go func(n int) {
			defer wg.Done()    // Уменьшаем счетчик группы синхронизации по завершении работы воркера
			fmt.Println(n * n) // Выводим результат вычисления квадрата числа
		}(n)
	}

	wg.Wait() // Ждем завершения выполнения всех задач
}

func main() {
	massive := []int{2, 4, 6, 8, 10} // Список чисел для вычисления квадратов
	numWorkers := 3                  // Количество воркеров для параллельного вычисления квадратов

	fmt.Println("Вычисление квадратов чисел через воркеров:")
	calculateSquaresWithWorkers(massive, numWorkers) // Вызов функции для вычисления квадратов с помощью воркеров

	fmt.Println("\nВычисление квадратов чисел через sync.WaitGroup:")
	calculateSquaresWithWaitGroup(massive) // Вызов функции для вычисления квадратов через sync.WaitGroup
}
