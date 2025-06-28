package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// --- Способ 1: Остановка горутины по условию (флагу) ---
// Горутина выполняет работу в цикле и периодически проверяет флаг.
// Как только внешний код меняет флаг на false — цикл завершается и горутина выходит.
func exitByFlag() {
	flag := true

	go func() {
		for flag {
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond) // Имитация работы
		}
		// Горутина завершилась после изменения флага
		fmt.Println("Goroutine exited")
	}()

	time.Sleep(2 * time.Second) // Даём горутине поработать
	flag = false                // Сигнал к завершению (меняем флаг)
	time.Sleep(2 * time.Second) // Ждём, чтобы увидеть сообщение о завершении
}

// --- Способ 2: Остановка горутины через канал уведомления ---
// Горутина работает в цикле и с помощью select слушает канал stop.
// Как только канал закрывают (close(stop)), горутина завершает работу.
func exitByChannel() {
	stop := make(chan struct{}) // Канал-уведомление для остановки

	go func() {
		for {
			select {
			case <-stop:
				// Получен сигнал завершения через канал
				fmt.Println("Goroutine exited")
				return
			default:
				// Основная работа горутины
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second) // Даём горутине поработать
	close(stop)                 // Сигнал на завершение через закрытие канала
	time.Sleep(2 * time.Second) // Ждём, чтобы увидеть сообщение о завершении
}

// --- Способ 3: Остановка горутины через context.Context ---
// Горутина слушает канал завершения из контекста и корректно выходит по сигналу отмены.
func exitByContext() {
	// Создаём контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				// Контекст был отменён — завершаем горутину
				fmt.Println("Goroutine exited")
				return
			default:
				// Основная работа горутины
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second) // Даём горутине поработать
	cancel()                    // Отправляем сигнал отмены контекста (остановка)
	time.Sleep(2 * time.Second) // Ждём, чтобы увидеть сообщение о завершении
}

// --- Способ 4: Прекращение работы горутины через runtime.Goexit() ---
// Горутина завершает работу мгновенно через вызов Goexit.
// Все defer будут выполнены.
func exitByGoexit() {
	go func() {
		defer fmt.Println("Deferred: Goroutine is exiting (Goexit called)")
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("About to call runtime.Goexit()")
		runtime.Goexit() // Немедленно завершает горутину
		// Эта строка не будет выполнена
		fmt.Println("This will never be printed")
	}()
	time.Sleep(2 * time.Second) // Дадим горутине поработать и завершиться
}

func main() {
	exitByFlag()
	exitByChannel()
	exitByContext()
	exitByGoexit()
}
