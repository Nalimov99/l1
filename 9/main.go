package main

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала:
в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	writeChannel := make(chan int)
	doubleChannel := make(chan int)

	go func() {
		for n := range writeChannel {
			doubleChannel <- n * n
		}
		close(doubleChannel)
	}()

	go func() {
		for _, n := range input {
			writeChannel <- n
		}
		close(writeChannel)
	}()

	for n := range doubleChannel {
		fmt.Println(n)
	}
}
