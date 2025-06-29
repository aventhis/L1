// swap.go
// ---------------------------
// Обмен значениями двух переменных без использования третьей переменной.
// Способ 1: через сложение/вычитание.
// Способ 2: через XOR (битовую операцию).
package main

import "fmt"

func main() {
	// --- 1. Арифметический обмен ---
	a := 5
	b := 8

	fmt.Println("До обмена:a =", a, "b =", b)
	// Обмен значениями через сложение/вычитание
	a = a + b // a = 13 b = 8
	b = a - b // b = 5 a = 13
	a = a - b // b = 5 a = 8
	fmt.Println("После обмена:a =", a, "b =", b)
	fmt.Println()

	// --- 2. Обмен через XOR ---
	x := 12
	y := 42

	fmt.Println("До обмена (XOR): x =", x, "y =", y)
	// Обмен значениями через XOR
	x = x ^ y // x = 38, y = 42
	y = x ^ y // y = 12, x = 38
	x = x ^ y // x = 42, y = 12
	fmt.Println("После обмена (XOR): x =", x, "y =", y)
}
