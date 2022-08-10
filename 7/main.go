package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type MyMap struct {
	Map map[int]int
	mu  sync.Mutex
}

func (m *MyMap) Set(key, value int) {
	m.mu.Lock()
	m.Map[key] = value
	m.mu.Unlock()
}

func main() {
	r := rand.New(rand.NewSource(99))
	var wg sync.WaitGroup
	m := MyMap{
		Map: make(map[int]int),
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			key, val := r.Intn(11), r.Intn(100)
			m.Set(key, val)
			fmt.Printf("Field should be: {%d:%d}\n", key, val)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Map:\n", m.Map)
}
