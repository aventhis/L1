// Реализация конвейера чисел с двумя этапами:
// генерация и обработка (*2) через каналы и горутины.
package main

import (
	"fmt"
)

// Функция-генератор: отправляет числа из массива в канал out и закрывает его по завершении.
func generator(arr []int, out chan<- int) {
	for _, v := range arr {
		out <- v
	}
	close(out)
}

// Функция-обработчик: читает числа из in, умножает на 2 и отправляет в out. Закрывает out по завершении.
func processor(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func main() {
	// Входные данные
	arr := []int{1, 2, 3, 4, 5}

	// Создаём два канала для конвейера
	ch := make(chan int)
	ch1 := make(chan int)

	// Запускаем этап генерации чисел
	go generator(arr, ch)

	// Запускаем этап обработки чисел
	go processor(ch, ch1)

	// Главная горутина читает результаты из второго канала и печатает их
	for v := range ch1 {
		fmt.Println(v)
	}
}
