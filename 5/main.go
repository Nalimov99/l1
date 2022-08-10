package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(99))
	duration := time.Duration(5 * time.Second)
	c := make(chan int)

	go func() {
		for {
			c <- r.Intn(11)
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	<-time.After(duration)
}
