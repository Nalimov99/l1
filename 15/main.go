package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

/*
Возможные негативные последствия:
1) Строки в голанг представляют из себя иммутабельный слайс байтов. При ре-слайсе слайса копия изначального массива не создается
После инициализации строки состоящий из 1024 символов, береться 100 элементов,
а остальные байты будут продолжать висеть в памяти, без возможности удаление (т.к. у нас глобальная переменная)
через GC.
2) Символ представляет из себя набор нескольких байт, поэтому при использовании слайса на строку, может обрезаться часть
символьного представляние.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)[:100]
	m := make([]rune, 100)
	copy(m, runes)
	justString = string(m)
}

func createHugeString(len int) string {
	var sb strings.Builder
	for i := 0; i < len; i++ {
		sb.WriteString("€")
	}

	return sb.String()
}

func main() {
	someFunc()
	fmt.Println("\"" + justString + "\"")
}
