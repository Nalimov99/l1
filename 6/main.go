package main

import (
	"context"
	"fmt"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	// каналы
	c1 := make(chan int)
	q := make(chan struct{})
	go func() {
		for {
			select {
			case n := <-c1:
				fmt.Println(n)
			case <-q:
				fmt.Println("Селект с другим каналом")
			}
		}
	}()
	loopAndWrite(c1)
	q <- struct{}{}

	c2 := make(chan int)
	go func() {
		for n := range c2 {
			fmt.Println(n)
		}
		fmt.Println("Закрытие канала")
	}()
	loopAndWrite(c2)
	close(c2)

	// Контекст
	ctx, cancel := context.WithCancel(context.Background())
	c3 := make(chan int)
	go func() {
		for {
			select {
			case n := <-c3:
				fmt.Println(n)
			case <-ctx.Done():
				fmt.Println("Контекст")
				return
			}
		}
	}()
	loopAndWrite(c3)
	cancel()

	fmt.Scanln()
}

func loopAndWrite(c chan int) {
	data := []int{1, 2, 3, 4, 5}
	for _, item := range data {
		c <- item
	}
}
