package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	duration := time.Second * 2
	fmt.Println("after")
	fmt.Println("Sleep: ", duration)
	now := time.Now()
	fmt.Println("now: ", now)
	sleep(duration)
	fmt.Println("Wake up")
	fmt.Println("pass: ", time.Since(now))

	fmt.Println()

	fmt.Println("timer")
	fmt.Println("Sleep: ", duration)
	now = time.Now()
	fmt.Println("now: ", now)
	sleepTimer(duration)
	fmt.Println("Wake up")
	fmt.Println("pass: ", time.Since(now))

	fmt.Println()

	fmt.Println("ticker")
	fmt.Println("Sleep: ", duration)
	now = time.Now()
	fmt.Println("now: ", now)
	sleepTicker(duration)
	fmt.Println("Wake up")
	fmt.Println("pass: ", time.Since(now))
}

func sleep(d time.Duration) {
	<-time.After(d)
}

func sleepTimer(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
}

func sleepTicker(d time.Duration) {
	t := time.NewTicker(d)
	<-t.C
	t.Stop()
}
