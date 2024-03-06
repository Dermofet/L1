package main

import (
	"fmt"
	"sort"
)

// simpleIntersection - функция для поиска пересечения двух множеств без использования дополнительных структур данных
func simpleIntersection(set1, set2 []int) []int {
	result := []int{}
	for _, a := range set1 {
		for _, b := range set2 {
			if a == b {
				result = append(result, a) // Если элемент присутствует в обоих множествах, добавляем его в результат
				break
			}
		}
	}
	return result
}

// hashIntersection - функция для поиска пересечения двух множеств с использованием хэш-таблицы
func hashIntersection(set1, set2 []int) []int {
	setMap := make(map[int]bool) // Создаем хэш-таблицу для быстрого поиска элементов из первого множества
	result := []int{}

	// Заполняем хэш-таблицу элементами из первого множества
	for _, num := range set1 {
		setMap[num] = true
	}

	// Проверяем каждый элемент из второго множества на наличие в хэш-таблице
	for _, num := range set2 {
		if setMap[num] {
			result = append(result, num) // Если элемент присутствует в хэш-таблице, добавляем его в результат
		}
	}

	return result
}

// sortedIntersection - функция для поиска пересечения двух множеств с использованием отсортированных массивов
func sortedIntersection(set1, set2 []int) []int {
	sort.Ints(set1) // Сортируем оба множества
	sort.Ints(set2)
	result := []int{}
	i, j := 0, 0

	// Сравниваем элементы по очереди из обоих отсортированных массивов
	for i < len(set1) && j < len(set2) {
		if set1[i] < set2[j] {
			i++
		} else if set1[i] > set2[j] {
			j++
		} else {
			result = append(result, set1[i]) // Если элементы равны, добавляем его в результат
			i++
			j++
		}
	}
	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Поиск простого пересечения O(n^2)
	intersection := simpleIntersection(set1, set2)
	fmt.Println("Простое пересечение:", intersection)

	// Поиск пересечения с использованием хэш-таблицы O(n)
	intersection = hashIntersection(set1, set2)
	fmt.Println("Пересечение с использованием хэш-таблицы:", intersection)

	// Поиск отсортированного пересечения O(n*log(n))
	intersection = sortedIntersection(set1, set2)
	fmt.Println("Отсортированное пересечение:", intersection)
}
