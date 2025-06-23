package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h *Human) ChangeName(name string) {
	h.Name = name
}

func (h *Human) String() string {
	return fmt.Sprintf("name - %s, age -  %d", h.Name, h.Age)
}

type Action struct {
	Human
}

func NewAction(name string, age int) *Action {
	return &Action {
		Human : Human {
			Name : name,
			Age: age,
		},
	}
}

func main() {
	a:= NewAction("Kate", 24)
	
	fmt.Println("before change: ", a)
	
	a.ChangeName("Irina")
	
	fmt.Println("after change: ", a)
}