package main

import (
	"fmt"
	"strings"
)

// isUnique проверяет, что все символы в строке уникальны (регистронезависимо)
func isUnique(str string) bool {
	// Создаем карту для отслеживания встреченных символов
	charMap := make(map[rune]bool)

	// Проходим по каждому символу в строке
	for _, char := range strings.ToLower(str) {
		// Если символ уже встречался в строке, возвращаем false
		if charMap[char] {
			return false
		}
		// Иначе отмечаем символ как встреченный
		charMap[char] = true
	}

	// Если мы досмотрели до конца строки и не встретили повторяющихся символов, возвращаем true
	return true
}

func main() {
	// Тестовые строки
	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}

	// Проверяем каждую строку
	for _, str := range testStrings {
		fmt.Printf("Строка '%s': %v\n", str, isUnique(str))
	}
}
