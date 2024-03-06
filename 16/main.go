package main

import (
	"fmt"
)

// quickSort - функция для быстрой сортировки целочисленного массива
func quickSort(arr []int) {
	if len(arr) < 2 {
		return // Если длина массива меньше двух, он уже отсортирован
	}

	pivotIndex := len(arr) / 2 // Индекс опорного элемента
	pivot := arr[pivotIndex]   // Опорный элемент

	left := 0                // Левая граница
	right := len(arr) - 1    // Правая граница

	for left <= right {
		for arr[left] < pivot { // Находим элемент слева, который больше или равен опорному
			left++
		}
		for arr[right] > pivot { // Находим элемент справа, который меньше или равен опорному
			right--
		}
		if left <= right {
			// Обмен значениями
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// Рекурсивно сортируем подмассивы
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func main() {
	arr := []int{9, 3, 7, 5, 6, 4, 8, 2, 1}
	fmt.Println("Before sorting:", arr)
	quickSort(arr)
	fmt.Println("After sorting:", arr)
}
