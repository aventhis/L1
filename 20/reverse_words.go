// reverse_words.go
// ---------------------------------------------
// Разворот порядка слов в строке "на месте", без доп. срезов.
// Пример: "snow dog sun" → "sun dog snow".

package main

import "fmt"

func reverseWords(s string) string {
	// Преобразуем строку в руны (на случай Unicode)
	r := []rune(s)

	// 1. Развернуть всю строку целиком
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// 2. Разворачиваем каждое слово по отдельности
	start := 0 // начало слова
	for i := 0; i <= len(r); i++ {
		// i == len(r): обработать последнее слово (не заканчивается пробелом)
		// r[i] == ' ': обработать очередное слово
		if i == len(r) || r[i] == ' ' {
			// Развернуть слово r[start:i] на месте
			for l, r2 := start, i-1; l < r2; l, r2 = l+1, r2-1 {
				r[l], r[r2] = r[r2], r[l]
			}
			start = i + 1 // следующее слово
		}
	}

	return string(r)
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s)) // Ожидаем: "sun dog snow"
}
