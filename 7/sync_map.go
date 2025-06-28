// concurrent_map.go
// Реализация конкурентно-безопасной записи в map двумя способами:
// 1. Обычный map + sync.Mutex
// 2. sync.Map (встроенная concurrent-map в Go)
// чтобы проверить go run -race sync_map.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// ---------- Вариант 1: Обычный map с sync.Mutex ----------

// SafeMap — обёртка над map с защитой через мьютекс.
type SafeMap struct {
	mu sync.Mutex // Мьютекс для синхронизации доступа
	m  map[int]int
}

// Set — потокобезопасная запись в map.
func (s *SafeMap) Set(key, value int) {
	s.mu.Lock()      // Захватываем мьютекс перед записью
	s.m[key] = value // Записываем значение
	s.mu.Unlock()    // Освобождаем мьютекс
}

// Get — потокобезопасное чтение из map.
func (s *SafeMap) Get(key int) (int, bool) {
	s.mu.Lock() // Захватываем мьютекс перед чтением
	val, ok := s.m[key]
	s.mu.Unlock() // Освобождаем мьютекс
	return val, ok
}

// MutexMap — демонстрация конкурентной записи в map с защитой через мьютекс.
func MutexMap() {
	fmt.Println("----- SafeMap with Mutex -----")
	data := SafeMap{
		m: make(map[int]int), // Инициализация map
	}

	var wg sync.WaitGroup // WaitGroup для ожидания завершения всех горутин

	// Запускаем 3 горутины, каждая пишет значения в map
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				data.Set(j, j*10) // Пишем значение по ключу
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	wg.Wait()

	// Выводим содержимое map (ключи от 0 до 4)
	for i := 0; i < 5; i++ {
		val, _ := data.Get(i)
		fmt.Printf("Key %d, Value %d\n", i, val)
	}
}

// ---------- Вариант 2: sync.Map ----------

// SyncMap — демонстрация конкурентной записи с sync.Map.
func SyncMap() {
	fmt.Println("----- sync.Map -----")
	var m sync.Map
	var wg sync.WaitGroup

	// Запускаем 3 горутины, каждая пишет значения в map
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				m.Store(j, j*10) // Пишем значение по ключу
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Printf("Key %v, Value %v\n", key, value)
		return true
	})
}

func main() {
	MutexMap()
	SyncMap()
}
