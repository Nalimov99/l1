package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	str := "snow dog sun"
	resultSlice := strings.Fields(str)

	for i, j := 0, len(resultSlice)-1; i < j; i, j = i+1, j-1 {
		resultSlice[i], resultSlice[j] = resultSlice[j], resultSlice[i]
	}

	fmt.Println(strings.Join(resultSlice, " "))
}
