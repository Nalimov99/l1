package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})
	res := []string{}

	for _, str := range slice {
		m[str] = struct{}{}
	}

	for key := range m {
		res = append(res, key)
	}

	fmt.Println(res)
}
