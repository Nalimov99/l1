package main

import (
	"context"
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

	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer stop()

	seedChannel := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < workersQty; i++ {
		go runWorker(ctx, seedChannel, i)
	}

	r := rand.New(rand.NewSource(99))
	for {
		select {
		case <-ctx.Done():
			cancel()
		default:
			seedChannel <- r.Intn(11)
		}
	}
}

func runWorker(ctx context.Context, c chan int, id int) {
	for {
		select {
		case number := <-c:
			fmt.Printf("Worker[%d]: %d\n", id, number)
		case <-ctx.Done():
			return
		}
	}
}
