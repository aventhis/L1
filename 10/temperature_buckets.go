// temperature_grouping.go
// ---------------------------
// Группировка температур по диапазонам
package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10

	groups := make(map[int][]float64)

	// Группируем значения по диапазонам
	for _, t := range arr {
		bucket := int(t) / step * step
		if t < 0 && t != float64(int(t)) && int(t)%step == 0 {
			bucket += step
		}
		groups[bucket] = append(groups[bucket], t)
	}

	// Печатаем результат
	for v, k := range groups {
		fmt.Printf("%d: %v ", v, k)
	}

}
