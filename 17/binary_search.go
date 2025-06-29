// binary_search.go
// ---------------------------
// Бинарный поиск в отсортированном слайсе.
// На вход: отсортированный слайс и искомое значение.
// На выход: индекс элемента или -1, если не найден.

package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9}
	res := binarySearch(arr, 3)
	if res == -1 {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Println("Элемент найден c индксом:", res)
	}
}

// binarySearch — реализует бинарный поиск в слайсе
func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1 // Задаём начальные границы поиска

	for left <= right { // Пока диапазон не сузился до невозможного
		mid := left + (right-left)/2 // Находим индекс середины диапазона

		if array[mid] == target { // Если нашли нужный элемент — возвращаем индекс
			return mid
		}
		if array[mid] > target { // Если середина больше искомого — ищем слева
			right = mid - 1
		} else { // Иначе — ищем справа
			left = mid + 1
		}
	}
	return -1 // Если не нашли — возвращаем -1
}
