// Конкурентное возведение в квадрат
// Написать программу, которая конкурентно рассчитает
// значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

// Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.

package main

import (
	"fmt"
	"sync"
)

// Вариант 1: Использование WaitGroup для ожидания завершения всех горутин
func squareWaitGroup(arr []int) {
	var wg sync.WaitGroup

	// Для каждого числа запускаем отдельную горутину
	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n) // Выводим квадрат числа
		}(num)
	}

	wg.Wait() // Ожидаем завершения всех горутин
}

// Вариант 2: Использование канала для получения результатов от горутин
func squareChannel(arr []int) {
	ch := make(chan int)

	for _, num := range arr {
		go func(n int) {
			ch <- num * num // Отправляем результат в канал
		}(num)
	}

	// Получаем результаты из канала
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(<-ch)
	}

	close(ch)
}

// Вариант 3: Канал + WaitGroup для синхронизации и передачи результатов
func squareChannelWG(arr []int) {
	ch := make(chan int, len(arr)) // Буферизированный канал
	var wg sync.WaitGroup

	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- num * num // Отправляем результат в канал
		}(num)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	close(ch) // После завершения всех горутин закрываем канал

	// Читаем результаты из канала
	for res := range ch {
		fmt.Println(res)
	}
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	fmt.Println("**Вариант 1: WaitGroup**")
	squareWaitGroup(arr)

	fmt.Println("**Вариант 2: Channel**")
	squareChannel(arr)

	fmt.Println("**Вариант 3: Channel и WaitGroup**")
	squareChannelWG(arr)
}
