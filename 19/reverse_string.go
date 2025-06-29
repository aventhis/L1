// reverse_string.go
// ---------------------------------------------
// Программа для разворота строки (Unicode-безопасно).
// Корректно работает с кириллицей, emoji и другими символами.
// Использует []rune для правильной работы с Unicode.

package main

import "fmt"

// Функция reverseString разворачивает строку, корректно обрабатывая Unicode-символы (руны).
func reverseString(s string) {
	runes := []rune(s) // преобразуем строку в слайс рун (Unicode-safe)

	// вариант 1
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[j], runes[i] = runes[i], runes[j]
	}
	fmt.Println(string(runes))

	// еще возможные 2 варианты
	//for i := len(runes) - 1; i >= 0; i-- {
	//	//fmt.Printf("%c", runes[i]) // вариант 2
	//	fmt.Printf(string(runes[i])) // варинат 3
	//}
}

func main() {
	s := "Hello World!"
	s1 := "Привет"

	reverseString(s)
	reverseString(s1)
}
