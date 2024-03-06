package main

import (
	"fmt"
)

func generateNumbers(numbers []int, out chan<- int) {
	defer close(out) // Закрываем канал после отправки всех чисел
	for _, num := range numbers {
		out <- num // Отправляем число в канал
	}
}

func doubleNumbers(in <-chan int, out chan<- int) {
	defer close(out) // Закрываем второй канал после обработки всех чисел
	for num := range in {
		out <- num * 2 // Удваиваем число и отправляем во второй канал
	}
}

func printResults(in <-chan int) {
	for num := range in {
		fmt.Println(num) // Выводим результаты в stdout
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем каналы
	numChannel := make(chan int)
	resultChannel := make(chan int)

	// Запускаем горутины для работы с каналами
	go generateNumbers(numbers, numChannel)
	go doubleNumbers(numChannel, resultChannel)
	printResults(resultChannel)
}
