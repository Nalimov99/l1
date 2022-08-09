package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
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

	var wg sync.WaitGroup
	wg.Add(workersQty)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	seedChannel := make(chan int)

	for i := 0; i < workersQty; i++ {
		go runWorker(seedChannel, i, &wg)
	}

	runSeed(seedChannel, shutdown, &wg)
}

func runSeed(c chan int, shutdown chan os.Signal, wg *sync.WaitGroup) {
	r := rand.New(rand.NewSource(99))
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			c <- r.Intn(11)
		case <-shutdown:
			close(c)
			wg.Wait()
			fmt.Println("Gracefull shutdown")
			return
		}
	}
}

func runWorker(c chan int, id int, wg *sync.WaitGroup) {
	for number := range c {
		fmt.Printf("Worker[%d]: %d\n", id, number)
	}
	wg.Done()
}
