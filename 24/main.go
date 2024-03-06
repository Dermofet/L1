package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p Point) DistanceTo(other Point) float64 {
	deltaX := other.x - p.x
	deltaY := other.y - p.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
