package main

import (
	"fmt"
	"sort"
)

// binarySearch выполняет бинарный поиск элемента в отсортированном массиве arr.
// Возвращает индекс найденного элемента или -1, если элемент не найден.
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid // Элемент найден, возвращаем его индекс
		} else if arr[mid] < target {
			left = mid + 1 // Искать в правой половине массива
		} else {
			right = mid - 1 // Искать в левой половине массива
		}
	}
	return -1 // Элемент не найден
}

func main() {
	// Пример массива
	arr := []int{2, 5, 7, 9, 13, 18, 22, 27, 31, 36}

	// Сортировка массива (бинарный поиск работает только с отсортированными данными)
	sort.Ints(arr)

	for _, v := range []int{6, 15, 27} {
		// Выполнение бинарного поиска
		index := binarySearch(arr, v)

		// Вывод результата
		if index != -1 {
			fmt.Printf("Элемент %d найден в позиции %d\n", v, index)
		} else {
			fmt.Printf("Элемент %d не найден\n", v)
		}
	}
}
