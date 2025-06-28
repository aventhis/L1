// set_bit.go
// Программа для установки или сброса i-го бита в числе типа int64.
// Для установки используем битовые операции | и &^.
package main

import "fmt"

// setBitToOne устанавливает i-й бит в 1
func setBitToOne(i uint, n int64) int64 {
	mask := int64(1) << i
	return n | mask
}

// setBitToZero сбрасывает i-й бит в 0
func setBitToZero(i uint, n int64) int64 {
	mask := int64(1) << i
	return n &^ mask
}

func main() {
	fmt.Println("----- установка 1-го бита в 1 -----")

	var n int64 = 25
	fmt.Printf("До изменения: %d %b\n", n, n) // 25 (11001₂)
	new_n := setBitToOne(1, n)
	fmt.Printf("После изменения 1го бита на 1: %d %b\n", new_n, new_n) // 27 (11011₂)

	fmt.Println()

	fmt.Println("----- установка 0-го бита в 0 -----")

	n = 5
	fmt.Printf("До изменения: %d %b\n", n, n) // 5 (101₂)
	new_n = setBitToZero(0, n)
	fmt.Printf("После изменения 0го бита на 0: %d %b\n", new_n, new_n) // 4 (100₂)
}

// setBitToZero сбрасывает i-й бит в 0
