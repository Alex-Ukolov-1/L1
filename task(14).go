package main

import (
	"fmt"
)

//функция определяющая тип
func knowtype(parametr interface{}) {
	switch parametr.(type) {
	case int:
		fmt.Println("type of int", parametr)
	case string:
		fmt.Println("type of string", parametr)
	case bool:
		fmt.Println("type of bool", parametr)
	case chan int:
		fmt.Println("type of chan int", parametr)
	case chan string:
		fmt.Println("type of chan string", parametr)
	case chan bool:
		fmt.Println("type of chan bool", parametr)
	default:
		fmt.Println("type of undefined", parametr)
	}
}

func main() {
	knowtype(1)
	knowtype("2")
	knowtype(true)
	knowtype(make(chan int))
	knowtype(make(chan string))
	knowtype(make(chan bool))
}
