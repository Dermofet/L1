package main

import "fmt"

func removeElement(slice []int, index int) []int {
	// Проверяем, что индекс находится в пределах слайса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Создаем новый слайс для хранения элементов перед удаляемым элементом
	before := make([]int, index)
	copy(before, slice[:index])

	// Создаем новый слайс для хранения элементов после удаляемого элемента
	after := make([]int, len(slice)-index-1)
	copy(after, slice[index+1:])

	// Объединяем два слайса в один
	result := append(before, after...)

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2 // Индекс элемента, который нужно удалить

	fmt.Println(slice) // Выведет: [1 2 3 4 5]

	// Удаляем элемент по указанному индексу
	slice = removeElement(slice, index)

	// Выводим результат
	fmt.Println(slice) // Выведет: [1 2 4 5]
}
