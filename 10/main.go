package main

import (
	"fmt"
	"math"
	"sort"
)

// groupTemperatures - функция для группировки температур
func groupTemperatures(groups map[int][]float64, temperatures []float64, step float64) map[int][]float64 {
	for _, temp := range temperatures {
		group := int(math.Floor(temp / step))       // Вычисляем номер группы для текущей температуры
		groups[group] = append(groups[group], temp) // Добавляем температуру в соответствующую группу
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Исходные температуры
	step := 10.0                                                                 // Шаг для группировки температур

	groups := make(map[int][]float64)                      // Создаем карту для хранения групп температур
	groups = groupTemperatures(groups, temperatures, step) // Группируем температуры

	// Сортируем ключи (группы) для вывода в порядке возрастания
	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Выводим результаты
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key*int(step), groups[key])
	}
}
