package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func main() {
	// mwg используется для последовательного выполнениния разных реализаций
	var mwg sync.WaitGroup
	input := []int{2, 4, 6, 8, 10}

	channel(input, &mwg)
	mwg.Wait()
	runWg(input, &mwg)
	mwg.Wait()
}

// Реализация с помощью канала
func channel(input []int, mwg *sync.WaitGroup) {
	mwg.Add(1)

	c := make(chan int)
	var sum int64

	go func() {
		for number := range c {
			atomic.AddInt64(&sum, int64(number*number))
		}
		fmt.Println("Channel: ", sum)
		mwg.Done()
	}()

	var counter int64
	for _, number := range input {
		go func(n int) {
			atomic.AddInt64(&counter, 1)
			c <- n
			if atomic.LoadInt64(&counter) == int64(len(input)) {
				close(c)
			}
		}(number)
	}
}

// Реализация с помощью wg
func runWg(input []int, mwg *sync.WaitGroup) {
	mwg.Add(1)

	var wg sync.WaitGroup
	var sum int64

	wg.Add(len(input))
	for _, number := range input {
		go func(n int64) {
			atomic.AddInt64(&sum, int64(n*n))
			wg.Done()
		}(int64(number))
	}
	wg.Wait()
	fmt.Println("WG: ", sum)

	mwg.Done()
}
