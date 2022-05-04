package main

import (
	"fmt"
)

func main() {
	//Поменять местами два числа без создания временной переменной.
	a, b := 10, 5
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
