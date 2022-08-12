package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	counter int64
}

func (c *Counter) GetCounter() int64 {
	// В данном случае можно было вернуть без атомика
	return atomic.LoadInt64(&c.counter)
}

func (c *Counter) IncCounter() {
	atomic.AddInt64(&c.counter, 1)
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	goroutines := 100

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			counter.IncCounter()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.GetCounter())
}
