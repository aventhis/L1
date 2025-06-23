// Работа нескольких воркеров
// Реализовать постоянную запись данных в канал (в главной горутине).

// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

// Программа должна принимать параметром количество воркеров и
//  при старте создавать указанное число горутин-воркеров.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numWorkers := 3 // Количество воркеров
	var wg sync.WaitGroup 
	jobs := make(chan int) // Канал для задач

	// Запускаем воркеров
	for i:=1; i <= numWorkers; i++ {
		go Work(i, jobs, &wg)
	}

	// Добавляем задачи
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}

func Work(id int, jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker id %d recieved job %d\n", id, job)
	}
}