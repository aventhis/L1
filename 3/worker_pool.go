// Работа нескольких воркеров
// Реализовать постоянную запись данных в канал (в главной горутине).

// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

// Программа должна принимать параметром количество воркеров и
//
//	при старте создавать указанное число горутин-воркеров.
package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// worker — функция-воркер, читает задачи из канала jobs и выводит их на экран.
// После завершения работы сообщает о завершении через WaitGroup.
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(500 * time.Millisecond) // Имитация работы
		fmt.Printf("Worker id %d recieved job %d\n", id, job)
	}
}

func main() {
	var n int
	// Принимаем количество воркеров через флаг -n
	flag.IntVar(&n, "n", 3, "number of goroutines")
	flag.Parse()

	var wg sync.WaitGroup  // WaitGroup для ожидания завершения всех воркеров
	jobs := make(chan int) // Канал для передачи задач воркерам

	// Запускаем n воркеров, каждому выдаём уникальный id
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Генерируем задачи и отправляем их в канал jobs
	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs) // Закрываем канал после отправки всех задач
	wg.Wait()   // Ожидаем завершения всех воркеров
}
