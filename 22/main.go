package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	a := big.NewInt(1 << 22)
	b := big.NewInt(1 << 21)
	res := big.NewInt(0)

	fmt.Printf("a: %d, b: %d\n", a, b)

	res.Mul(a, b)
	fmt.Printf("Умножение: %d\n", res)

	res.Div(a, b)
	fmt.Printf("Деление: %d\n", res)

	res.Add(a, b)
	fmt.Printf("Сложение: %d\n", res)

	res.Sub(a, b)
	fmt.Printf("Вычитание: %d\n", res)
}
