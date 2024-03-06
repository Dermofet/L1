package main

import "fmt"

func main() {
	// Последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустое множество
	set := make(map[string]bool)

	// Добавляем каждую строку в множество
	for _, str := range sequence {
		set[str] = true
	}

	// Выводим содержимое множества
	fmt.Println("Множество:", set)

	// Проверяем наличие элементов в множестве
	fmt.Println("Элемент 'cat' присутствует:", set["cat"])
	fmt.Println("Элемент 'bird' присутствует:", set["bird"])
}
