package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

type Namer interface {
	Name() string
}

type Human struct {
	Firstname string
	Lastname  string
}

func (h *Human) GetFullName() string {
	return h.Firstname + " " + h.Lastname
}

type HumanAdapter struct {
	Person Human
}

func (h *HumanAdapter) Name() string {
	return h.Person.GetFullName()
}

type Animal struct {
	Alias string
}

func (a *Animal) Name() string {
	return a.Alias
}

type Caller struct {
}

func (c *Caller) CallByName(n Namer) {
	fmt.Println(n.Name())
}

func main() {
	caller := Caller{}
	animal := Animal{
		Alias: "Пёсель",
	}
	human := Human{
		Firstname: "Илья",
		Lastname:  "Налимов",
	}
	adapter := HumanAdapter{
		Person: human,
	}

	caller.CallByName(&animal)
	caller.CallByName(&adapter)
}
