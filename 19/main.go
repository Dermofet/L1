package main

import (
	"fmt"
)

// reverseString переворачивает строку, учитывая символы Unicode
func reverseString(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем местами символы в срезе рун
	}
	return string(runes) // Преобразуем срез рун обратно в строку
}

func main() {
	input := "Пример строки с символами Unicode: 😀🚀💻"

	reversed := reverseString(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Println("Перевернутая строка:", reversed)
}
