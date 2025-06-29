// type_switch.go
// -------------------------------------
// Определение типа переменной в runtime (через type switch и type assertion).
// Программа принимает переменную interface{} и выводит её тип.
// Обрабатываем: int, string, bool, chan, а остальные — "Unknown type".
package main

import "fmt"

func main() {
	// Примеры для всех запрошенных типов
	specifyType(32)
	specifyType("text")
	specifyType(true)
	intCh := make(chan int)
	specifyType(intCh)
	stringCh := make(chan string)
	specifyType(stringCh)
}

func specifyType(n interface{}) {
	// Используем type switch для int, string, bool, chan
	switch v := n.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("unknown type")
	}
}
