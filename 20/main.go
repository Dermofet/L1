package main

import (
	"fmt"
	"strings"
)

// reverse переворачивает слова в строке
func reverse(str string) string {
	// Создаем срез из строки
	words := strings.Split(str, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {

		// Меняем местами
		words[i], words[j] = words[j], words[i]

	}
	// Соединяем срез в строку
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverse(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутые слова: %s\n", reversed)
}
