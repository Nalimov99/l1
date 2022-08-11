package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	firstSlice := []int{1, 2, 3, 4, 5}
	secondSlice := []int{8, 10, 2, 0, 1, 3, 3}
	res := make(map[int]int)

	mapSet(res, firstSlice)
	mapSet(res, secondSlice)

	fmt.Println("Пересечения:")
	for key, num := range res {
		if num > 1 {
			fmt.Println(key)
		}
	}
}

func mapSet(m map[int]int, slice []int) {
	for _, n := range slice {
		m[n] += 1
	}
}
