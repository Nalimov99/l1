package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	slice := []int{1, 3, 2, 4, 5}

	fmt.Println(binarySearch(slice, 5))
}

func binarySearch(input []int, search int) (index int, ok bool) {
	slice := make([]int, len(input))
	copy(slice, input)
	sort.Ints(slice)

	start := 0
	end := len(slice) - 1
	for start <= end {
		mid := (start + end) / 2
		switch {
		case slice[mid] == search:
			index = mid
			return index, true
		case slice[mid] < search:
			start = mid + 1
		case slice[mid] > search:
			end = mid - 1
		}
	}

	return
}
