// Конкурентное возведение в квадрат
// Написать программу, которая конкурентно рассчитает 
// значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

// Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.

package main

import (
	"fmt"
	"sync"
)

//решение с использованием WaitGroup
func squareWaitGroup(arr []int) {
	var wg sync.WaitGroup 

	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n*n)
		}(num)
	}

	wg.Wait()
}

//решение с использованием каналов
func squareChannel(arr []int) {
	ch := make(chan int)

	for _, num := range arr {
		go func(n int) {
			ch <- num * num
		}(num)
	}
	
	for i := len(arr)-1; i >= 0; i-- {
		fmt.Println(<- ch)
	}

	close(ch)
}

//решение с использованием каналов + WaitGroup
func squareChannelWG(arr []int) {
	ch := make(chan int, len(arr))
	var wg sync.WaitGroup 

	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- num * num
		}(num)
	}
	
	wg.Wait()
	close(ch)

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

