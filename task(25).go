package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep

//реализация функции sleep
func sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Print("a")

	}

}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Print("b")

	}

}

func main() {
	//функция a с горутиной
	go a()
	sleep(5)
	//функция b с горутиной
	go b()
	sleep(5)
	fmt.Println("end main()")
}
