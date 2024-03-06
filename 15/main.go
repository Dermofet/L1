/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

createHugeString - функция создания строки заданного размера
при присваивании justString среза строки v[:100] (т.е. justString будет ссылаться на копию базового слайса).
После выхода из функции someFunc() созданная строка никуда не денется из памяти, потому что
глобальная переменная justString ссылается на часть этой строки, что может привести к утечке памяти.
Решение - создание копии строки.
*/

package main

import (
	"strings"
)

// Функция для создания строки заданного размера
func createHugeString(size int) string {
	return strings.Repeat("Hi!", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
