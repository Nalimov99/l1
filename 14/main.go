package main

import (
	"fmt"
	"reflect"
)

/*Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	printType(1)
	printType("asd")
	printType(true)
	printType(make(chan int))
}

func printType(value interface{}) {
	fmt.Println(reflect.TypeOf(value))
}
