package main

import (
	"fmt"
)

// Human — структура, описывающая человека с именем и возрастом
type Human struct {
	Name string
	Age  int
}

// ChangeName — метод Human для изменения имени
func (h *Human) ChangeName(name string) {
	h.Name = name
}

// String — реализует интерфейс Stringer для красивого вывода Human
func (h *Human) String() string {
	return fmt.Sprintf("name - %s, age -  %d", h.Name, h.Age)
}

// Action — структура, "встраивающая" Human (композиция)
type Action struct {
	Human // Встраивание структуры Human
}

// NewAction — конструктор для создания Action c нужными полями Human
func NewAction(name string, age int) *Action {
	return &Action{
		Human: Human{
			Name: name,
			Age:  age,
		},
	}
}

func main() {
	// Создаем экземпляр Action с начальными значениями
	a := NewAction("Kate", 24)
	fmt.Println("before change: ", a)

	// Меняем имя через метод, унаследованный от Human
	a.ChangeName("Irina")
	fmt.Println("after change: ", a)
}
