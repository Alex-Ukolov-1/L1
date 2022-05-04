package main

import (
	"fmt"
)

type Count struct {
	count int
}

//функция счётчик для переменной структуры Count и канала
func count(atr *Count, channel chan int) {
	//счётчик
	atr.count++
	channel <- atr.count
}

func main() {
	channelA := make(chan int)
	//инициализация Count
	vv := Count{0}
	for i := 0; i < 100; i++ {
		//запуск горутины
		go count(&vv, channelA)
		<-channelA
	}
	//закрыть канал
	defer close(channelA)
	fmt.Println(vv)
}
