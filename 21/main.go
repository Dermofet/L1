package main

import (
	"fmt"
)

// LegacyPrinter - устаревший принтер
type LegacyPrinter struct{}

// Print - метод устаревшего принтера
func (p *LegacyPrinter) Print(text string) {
	fmt.Println("Устаревший принтер:", text)
}

// ModernPrinter - интерфейс для современного принтера
type ModernPrinter interface {
	PrintMessage(message string)
}

// Adapter - адаптер для преобразования устаревшего принтера в современный
type Adapter struct {
	LegacyPrinter *LegacyPrinter
}

// PrintMessage - метод адаптера для современного принтера
func (a *Adapter) PrintMessage(message string) {
	// Вызываем устаревший метод через адаптер
	a.LegacyPrinter.Print(message)
}

func main() {
	// Создаем устаревший принтер
	legacyPrinter := &LegacyPrinter{}

	// Создаем адаптер для устаревшего принтера
	adapter := &Adapter{
		LegacyPrinter: legacyPrinter,
	}

	// Используем адаптер как современный принтер
	adapter.PrintMessage("Привет, мир!")
}
