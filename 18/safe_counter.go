// safe_counter.go
// ---------------------------------------------
// Конкурентный счетчик: два способа (Mutex, Atomic)
// Запускает несколько горутин, каждая инкрементирует счетчик много раз.
// После завершения работы выводится итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ----------------- Вариант 1: sync.Mutex -----------------

type MutexCounter struct {
	counter int        // значение счетчика
	mu      sync.Mutex // мьютекс для синхронизации доступа
}

// Inc — инкрементирует счётчик безопасно для конкуренции
func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

// Value — потокобезопасно возвращает значение счётчика
func (c *MutexCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

// ----------------- Вариант 2: sync/atomic -----------------

type AtomicCounter struct {
	value int64
}

// Inc — потокобезопасный инкремент через atomic
func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

// Value — потокобезопасное получение значения через atomic
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("----- MutexCounter -----")

	c := MutexCounter{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				c.Inc() // каждая горутина инкрементирует 20 раз
			}
		}()
	}
	wg.Wait()
	fmt.Println("Итоговое значение (Mutex):", c.Value())
	fmt.Println()

	fmt.Println("----- AtomicCounter -----")
	wg = sync.WaitGroup{} // Обнуляем wg
	a := AtomicCounter{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				a.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Итоговое значение (Atomic):", a.Value())
}
