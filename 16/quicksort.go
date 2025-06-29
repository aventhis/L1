// quicksort.go
// ---------------------------
// Быстрая сортировка массива целых чисел с использованием рекурсии.
// Для выбора опорного элемента берём первый элемент.
// Разделяем массив на две части: меньше и больше (или равные) pivot.
package main

import "fmt"

func main() {
	arr := []int{24, 12, 4, 5, 10, -24}
	res := quickSort(arr)
	fmt.Println(res)
}

// quickSort — быстрая сортировка (QuickSort) слайса целых чисел.
// arr — слайс чисел, который нужно отсортировать.
// Возвращает новый отсортированный слайс.
func quickSort(arr []int) []int {
	var res []int

	// Если длина массива меньше 2 — он уже отсортирован
	if len(arr) < 2 {
		return arr
	}
	// Выбираем опорный элемент (pivot) — первый элемент массива
	pivot := arr[0]
	var left, right []int

	// Разделяем элементы: меньше pivot — налево, остальные — направо
	for _, num := range arr[1:] { // пропускаем pivot!
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	// Рекурсивно сортируем левую и правую части, собираем результат
	res = append(res, quickSort(left)...)
	res = append(res, pivot)
	res = append(res, quickSort(right)...)
	return res
}
