package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	flag.Parse()

	workersQty, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic("number of workers should be int")
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	seedChannel := make(chan int)

	for i := 0; i < workersQty; i++ {
		go runWorker(seedChannel, i)
	}

	runSeed(seedChannel, shutdown)
}

func runSeed(c chan int, shutdown chan os.Signal) {
	r := rand.New(rand.NewSource(99))
	for {
		select {
		case <-shutdown:
			close(c)
			return
		default:
			c <- r.Intn(11)
		}
	}
}

func runWorker(c chan int, id int) {
	for number := range c {
		fmt.Printf("Worker[%d]: %d\n", id, number)
	}
}
