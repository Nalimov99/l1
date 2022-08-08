package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)
*/

type Human struct {
	Firstname string
	Lastname  string
}

type Action struct {
	Human
}

func (h *Human) GetFullName() string {
	return h.Firstname + " " + h.Lastname
}

func main() {
	action := Action{
		Human{
			Firstname: "Ilia",
			Lastname:  "Nalimov",
		},
	}

	fmt.Println(action.GetFullName())
	fmt.Println(action.Human.GetFullName())
}
