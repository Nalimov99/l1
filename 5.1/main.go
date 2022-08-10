package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	r := rand.New(rand.NewSource(99))
	duration := time.Duration(3 * time.Second)
	c := make(chan int)

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			c <- r.Intn(11)
		}
	}
}
