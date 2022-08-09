package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	// mainWg используется для последовательного выполнения разных реализаций
	var mainWg sync.WaitGroup

	input := []int{2, 4, 6, 8, 10}

	channel(input, &mainWg)
	mainWg.Wait()
	runWg(input, &mainWg)
	mainWg.Wait()
}

// Реализация с помощью каналов
func channel(input []int, mwg *sync.WaitGroup) {
	mwg.Add(1)

	c := make(chan int)
	go func() {
		for number := range c {
			fmt.Println("Channel: ", number*number)
		}
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
	wg.Add(len(input))

	for _, number := range input {
		go wgReader(number, &wg)
	}
	wg.Wait()
	mwg.Done()
}

func wgReader(number int, wg *sync.WaitGroup) {
	fmt.Println("WG: ", number*number)
	wg.Done()
}
