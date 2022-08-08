package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	input := []int{2, 4, 6, 8, 10}

	go channelWorker(input)
	go wgWorker(input)

	fmt.Scanln()
}

// Реализация с помощью каналов
func channelWorker(input []int) {
	c := make(chan int)
	go channelReader(c)
	for _, number := range input {
		go func(n int) {
			c <- n
		}(number)
	}
}

func channelReader(c chan int) {
	for number := range c {
		fmt.Println("Channel: ", number*number)
	}
}

// Реализация с помощью wg
func wgWorker(input []int) {
	var wg sync.WaitGroup
	wg.Add(len(input))

	for _, number := range input {
		go wgReader(number, &wg)
	}
	wg.Wait()
}

func wgReader(number int, wg *sync.WaitGroup) {
	fmt.Println("WG: ", number*number)
	wg.Done()
}
