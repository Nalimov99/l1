package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	fmt.Println("abcd: ", isUniqueString("abcd"))
	fmt.Println("abCdefAaf: ", isUniqueString("abCdefAaf"))
	fmt.Println("aabcd: ", isUniqueString("aabcd"))
	fmt.Println("Aa: ", isUniqueString("Aa"))
}

func isUniqueString(str string) (isUnique bool) {
	hash := make(map[rune]struct{})
	runes := []rune(strings.ToLower(str))

	for _, r := range runes {
		if _, ok := hash[r]; ok {
			return
		}

		hash[r] = struct{}{}
	}

	return true
}
