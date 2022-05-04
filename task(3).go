package main

import (
	"fmt"
	//"sync"
)

//массив
var box = [5]int{2, 4, 6, 8, 10}

//функция принимающая канал
func odd(channel chan int) {
	for i := 0; i < len(box); i++ {
		//выводит результат в канал
		channel <- (box[i] * box[i])
	}
}

func main() {
	var s int
	//создание канала
	channelA := make(chan int)
	//вызов функции в горутине
	go odd(channelA)
	//закрытие канала
	defer close(channelA)
	//цикл для сложения
	for i := 0; i < len(box); i++ {
		//сложение суммы из канала
		s += (<-channelA)
	}
	//печать
	fmt.Println(s)
}

//func main() {
//
//  создание переменной типа sync.waitgroup
//	var wg sync.WaitGroup
//  создание переменной типа sync.Mutex
//	var mutex sync.Mutex
//
//	var sum int
//  массив
//	nums := []int{2, 4, 6, 8, 10}
//
//  добавляет длинну массива в счетчик WaitGroup
//	wg.Add(len(nums))
//  перебор массива
//	for _, num := range nums {
//  вызов анонимной функции горутиной
//		go func(num int) {
//  вызывающая goroutine блокируется до тех пор, пока мьютекс не будет доступен.
//			mutex.Lock()
//  разблокировка mutex
//			defer mutex.Unlock()
//
//			sum += num * num
//  уменьшает значение счётчика на еденицу
//			wg.Done()
//  обьявляем об использовании переменной num для параметра num
//		}(num)
//	}
//  ждёт, пока счетчик WaitGroup не достигнет нуля.
//	wg.Wait()
//
//	fmt.Println(sum)
//}
