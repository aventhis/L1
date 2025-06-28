// set_intersection.go
// ---------------------------
// Пересечение двух неупорядоченных множеств (слайсов)
// Находим элементы, присутствующие и в первом, и во втором слайсе.
// Используем map для эффективного поиска.

package main

import "fmt"

func main() {
	// Данные
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	// Создаём map для быстрого поиска элементов из первого множества
	setA := make(map[int]struct{})
	for _, val := range a {
		setA[val] = struct{}{}
	}

	// Ищем пересечение
	var intersection []int
	for _, val := range b {
		if _, ok := setA[val]; ok {
			intersection = append(intersection, val)
		}
	}

	fmt.Println(intersection)
}
