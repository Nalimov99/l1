package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("До: ", slice)
	slice = removeIndex(slice, 1)
	fmt.Println("После: ", slice)
}

func removeIndex[T any](s []T, index int) []T {
	if index < 0 || index > len(s)-1 {
		return s
	}

	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
