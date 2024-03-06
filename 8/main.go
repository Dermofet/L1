package main

import "fmt"

// Устанавливает i-й бит в 1 или 0
func setBit(n int64, i int, bitValue int) int64 {
	// Создаем маску, установленную в 1 в позиции i
	mask := int64(1) << (i - 1)

	// Если bitValue равно 1, устанавливаем i-й бит в 1, иначе в 0
	if bitValue == 1 {
		return n | mask // Устанавливаем бит в 1
	} else {
		return n &^ mask // Устанавливаем бит в 0
	}
}

func main() {
	// 1010
	var num int64 = 10

	// Устанавливаем 3-й бит в 1
	// 1010 -> 1110
	num = setBit(num, 3, 1)
	fmt.Printf("Установлен 3-й бит: %b\n", num)

	// Устанавливаем 2-й бит в 0
	// 1110 -> 1100
	num = setBit(num, 2, 0)
	fmt.Printf("Установлен 5-й бит: %b\n", num)
}
