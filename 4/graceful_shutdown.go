// Завершение по Ctrl+C
// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Решение: используем context.Context для оповещения всех воркеров о необходимости завершиться.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worker — функция-воркер, читает задачи из канала jobs и выводит их на экран.
// После завершения работы сообщает о завершении через WaitGroup.
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Если получен сигнал завершения через контекст
			fmt.Printf("worker %d done\n", id)
			return
		case job, ok := <-jobs:
			if !ok { // Если канал закрыт, выходим из воркера
				return
			}
			time.Sleep(500 * time.Millisecond) // Имитация работы
			fmt.Printf("worker %d received job %d\n", id, job)
		}
	}

}

func main() {
	var n int
	// Принимаем количество воркеров через флаг -n
	flag.IntVar(&n, "n", 3, "number of goroutines")
	flag.Parse()

	var wg sync.WaitGroup  // WaitGroup для ожидания завершения всех воркеров
	jobs := make(chan int) // Канал для передачи задач воркерам

	ctx, cancel := context.WithCancel(context.Background()) // Контекст для оповещения о завершении

	//Перехват SIGINT
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Горутина, ожидающая сигнал завершения и отменяющая контекст
	go func() {
		<-sigs
		fmt.Println("\nReceived shutdown signal, shutting down gracefully")
		cancel()
	}()

	// Запускаем n воркеров, каждому выдаём уникальный id
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// Генерируем задачи и отправляем их в канал jobs
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done(): // Если получен сигнал завершения — выходим из генерации задач
			break
		default:
			jobs <- i
		}
	}

	close(jobs)
	wg.Wait() // Ожидаем завершения всех воркеров
	fmt.Println("All goroutines finished")
}

//Я выбрала context.Context для завершения воркеров, потому что это стандартный и наиболее безопасный способ передачи
//сигнала об остановке во всех горутинах. Контекст проще расширять (например, для таймаутов),
//его удобно передавать в функции, и он исключает ошибки, связанные с закрытием каналов.
//Такой подход считается лучшей практикой для graceful shutdown в Go.
