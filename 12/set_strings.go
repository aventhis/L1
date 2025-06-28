// set_of_strings.go
// ---------------------------
// Собственное множество строк.
// Дан слайс строк, требуется получить множество уникальных слов (set).
// Используем map[string]struct{} для хранения уникальных значений.

package main

import "fmt"

func main() {
	// Исходные данные
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём пустое множество (set) на map[string]struct{}
	set := make(map[string]struct{})

	// Добавляем все слова из слайса в set
	for _, v := range data {
		set[v] = struct{}{} // struct{}{} — минимальный пустой тип, не занимает памяти
	}

	// Выводим уникальные слова
	for val, _ := range set {
		fmt.Printf("%s ", val)
	}
}
