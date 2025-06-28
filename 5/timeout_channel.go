// Таймаут на канал
// Программа отправляет значения в канал и читает их из другой горутины.
// Через N секунд (таймаут) программа корректно завершается.
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// Получаем таймаут из флага командной строки, по умолчанию 3 секунды
	var seconds int
	flag.IntVar(&seconds, "seconds", 1, "timeout in seconds")
	flag.Parse()

	ch := make(chan int)                                        // Канал для передачи данных
	timeout := time.After(time.Duration(seconds) * time.Second) // Таймер-таймаут

	// Горутина-отправитель
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}()

	// Основная горутина: читает из канала, либо выходит по таймауту
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				// Канал закрыт, все данные получены — завершаем работу
				return
			}
			fmt.Println("Received value from channel:", val)
		case <-timeout:
			// Таймаут: завершаем программу
			fmt.Println("Timeout! Exiting program.")
			return
		}
	}

}
