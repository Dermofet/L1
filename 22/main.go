package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Используем большие числа для переменных a и b
	a, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
	b, _ := new(big.Int).SetString("987654321098765432109876543210", 10)

	fmt.Println("Переменная a: ", a)
	fmt.Println("Переменная b: ", b)

	// Умножение a на b
	mul := new(big.Int).Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление a на b
	div := new(big.Int).Div(a, b)
	fmt.Println("Деление:", div)

	// Сложение a и b
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сложение:", sum)

	// Вычитание b из a
	sub := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание:", sub)
}
