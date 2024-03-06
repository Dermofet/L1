package main

import "fmt"

// Определение структуры Human (Человек)
type Human struct {
	Name string // Имя
	Age  int    // Возраст
}

// SetName - метод для изменения значения поля Name (Изменить имя)
func (h *Human) SetName(name string) {
	h.Name = name
}

// SetAge - метод для изменения значения поля Age (Изменить возраст)
func (h *Human) SetAge(age int) {
	h.Age = age
}

// ShowInfo - метод для вывода информации (Показать информацию)
func (h *Human) ShowInfo() {
	fmt.Printf("Name: %s, Age: %d\n", h.Name, h.Age)
}

// Определение структуры Action (Действие) с встраиванием структуры Human
type Action struct {
	Human // встраивание структуры Human в структуру Action
}

// NewAction - конструктор структуры Action (Создание нового действия)
func NewAction(name string, age int) *Action {
	return &Action{Human{Name: name, Age: age}}
}

func main() {
	// Создание экземпляра структуры Action (Создание нового действия)
	action := NewAction("John", 30)

	// Вызов методов экземпляра (Вызов методов экземпляра действия)
	action.ShowInfo()

	// Изменение полей экземпляра (Изменение полей экземпляра)
	action.SetName("Jane") // Изменить имя на "Jane"
	action.SetAge(25)      // Изменить возраст на 25

	// Вызов методов экземпляра (Вызов методов экземпляра)
	action.ShowInfo() // Показать информацию о действии

	// Вызов методов класса Human (Вызов методов класса Human)
	action.Human.ShowInfo() // Показать информацию о человеке
}
