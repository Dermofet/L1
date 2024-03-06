package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Решение через Atomic
func sequenceByAtomic(nums []int64) int64 {
	var (
		wg     sync.WaitGroup // WaitGroup для ожидания завершения всех горутин
		result int64          // Переменная для хранения результата
	)

	// Обходим все числа в массиве
	for _, v := range nums {
		wg.Add(1) // Увеличиваем счетчик WaitGroup для каждой горутины
		go func(num int64) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы горутины
			// Атомик нам гарантирует потокобезопасное выполнение операции увеличения значения переменной
			atomic.AddInt64(&result, num*num)
		}(v) // Запускаем горутину с текущим числом из массива
	}

	wg.Wait() // Ждем завершения всех горутин

	return result // Возвращаем результат
}

// Решение через Mutex
func sequenceByMutex(nums []int64) int64 {
	var (
		wg     sync.WaitGroup // WaitGroup для ожидания завершения всех горутин
		mu     sync.Mutex     // Мьютекс для синхронизации доступа к переменной result
		result int64          // Переменная для хранения результата
	)

	// Обходим все числа в массиве
	for _, v := range nums {
		wg.Add(1) // Увеличиваем счетчик WaitGroup для каждой горутины
		go func(num int64) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы горутины
			mu.Lock()       // Блокируем мьютекс для синхронизации доступа к переменной result
			defer mu.Unlock() // Удаляем блокировку мьютекса после завершения работы горутины

			result += num * num // Выполняем операцию увеличения переменной result на квадрат текущего числа
		}(v) // Запускаем горутину с текущим числом из массива
	}

	wg.Wait() // Ждем завершения всех горутин

	return result // Возвращаем результат
}

func main() {
	nums := []int64{2, 4, 6, 8, 10} // Массив чисел

	// Вызываем функции для вычисления суммы квадратов чисел с использованием разных механизмов синхронизации
	fmt.Println(sequenceByAtomic(nums)) // Выводим результат работы функции sequenceByAtomic
	fmt.Println(sequenceByMutex(nums))  // Выводим результат работы функции sequenceByMutex
}
