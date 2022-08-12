package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func main() {
	slice := []int{10, 9, 8, 7, 7, 1, 5, 4, 0, 11, 44, 13}
	// Можно так:
	// sort.Ints(slice)

	qsort(slice)
	fmt.Println(slice)
}

func qsort[T numbers](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}

	left, right := 0, len(slice)-1
	pivot := len(slice) / 2
	slice[pivot], slice[right] = slice[right], slice[pivot]

	for i := range slice {
		if slice[i] < slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]

	qsort(slice[:left])
	qsort(slice[left+1:])

	return slice
}
