package main

import (
	"fmt"
	"time"
)

// sleep - функция для приостановки выполнения программы на заданное количество секунд
func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало программы")

	// Вызываем нашу функцию sleep на 3 секунды
	sleep(3)

	fmt.Println("Прошло 3 секунды")

	fmt.Println("Конец программы")
}
