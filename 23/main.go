package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("До: ", slice)
	slice = removeIndex(slice, 4)
	fmt.Println("После: ", slice)
}

func removeIndex[T any](s []T, index int) []T {
	if index < 0 || index > len(s)-1 {
		return s
	}

	return append(s[:index], s[index+1:]...)
}
