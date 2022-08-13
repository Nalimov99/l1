package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x, y,
	}
}

func (p *Point) GetDistance() int {
	return int(math.Abs(float64(p.x - p.y)))
}

func main() {
	p := NewPoint(50, 100)
	fmt.Println(p.GetDistance())
}
