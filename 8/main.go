package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	fmt.Println(setBit(100, 3, false))
}

func setBit(number int64, pos uint, shouldSetZero bool) int64 {
	if shouldSetZero {
		return number &^ (1 << pos)
	}

	return number | (1 << pos)
}
