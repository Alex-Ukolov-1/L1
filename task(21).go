package main

import (
	"fmt"
)

//Реализовать паттерн «адаптер» на любом примере.

type lamp interface {
	fire()
}

type lampoff struct {
	action string
	*energy
}

type lampon struct {
	action string
	*energy
}

type energy struct {
	status string
}

func (a *lampoff) fire() {
	fmt.Println("off " + a.action)
}

func (a *lampon) fire() {
	fmt.Println("on " + a.action)
}

func main() {

	first := &lampoff{"case", &energy{"working"}}
	second := &lampon{"case", &energy{"working"}}
	first.fire()
	second.fire()
}
