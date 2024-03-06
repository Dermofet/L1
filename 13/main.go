package main

import "fmt"

func main() {
	// Способ 1: Через замену
	a, b := 5, 10

	fmt.Println("Способ 1\nДо замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = b, a // Замена значений переменных

	fmt.Println("\nПосле замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Способ 2: Используя арифметические операции XOR
	a, b = 5, 10

	fmt.Println("\nСпособ 2\nДо замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a = a ^ b // Применяем XOR к a и b
	b = a ^ b
	a = a ^ b

	fmt.Println("\nПосле замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
