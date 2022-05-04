package main

import (
	"fmt"
	"os"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
//массива, во второй — результат операции x*2, после чего данные из второго
//канала должны выводиться в stdout.

var box = [5]int{2, 4, 6, 8, 10}

//пишутся значения из массива
func tochannel(channel chan int) {
	for i := 0; i < len(box); i++ {
		channel <- box[i]
	}
}

//пишется результат операции x*2
func fromchannel(arr []int, channel2 chan int) {
	for _, bb := range arr {
		channel2 <- bb * 2
	}
}

func main() {
	channelA := make(chan int)
	channelB := make(chan int)

	ve := []int{}
	go tochannel(channelA)
	for i := 0; i < len(box); i++ {
		//добавление результата в массив
		ve = append(ve, <-channelA)
	}

	go fromchannel(ve, channelB)
	defer close(channelA)
	defer close(channelB)
	for _, bb := range box {
		//вывод результата
		fmt.Fprintln(os.Stdout, <-channelB, "=2*", bb)
	}

}
