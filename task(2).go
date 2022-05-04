package main

import (
	"fmt"
	"os"
	//"sync"
	
)

//массив
var box = [5]int{2, 4, 6, 8, 10}

//функция odd принимающая целочисленное значение и канал
func odd(arr int, channel chan int) {
	//передача результата в канал
	channel <- (arr * arr)
}

func main() {
	//создание канала
	channelA := make(chan int)
	//закрытие канала
	defer close(channelA)
	//цикл для запуска горутины и вывода значений из канала
	for i := 0; i < len(box); i++ {
		//Канал передается функции, выполняющейся в новой горутине
		//вызов функции в горутине
		go odd(box[i], channelA)
		//печать данных в os.Stdout
		fmt.Fprintln(os.Stdout, <-channelA)
	}
}

// метод с использованием пакета sync
//func main() {
//  инициализация массива
//	nums := []int{2, 4, 6, 8, 10}
// создание переменной типа sync.waitgroup
//	var wg sync.WaitGroup
//  добавляет длинну массива в счетчик WaitGroup
//	wg.Add(len(nums))
//  перебор массива
//	for _, num := range nums {
//      вызов анонимной функции горутиной
//		go func(num int) {
// печать результата
//			fmt.Println(num * num)
// уменьшает значение счётчика на еденицу
//			wg.Done()
// обьявляем об использовании переменной num для параметра num
//		}(num)
//	}
//  ждёт, пока счетчик WaitGroup не достигнет нуля.
//	wg.Wait()
//}

//2 метод создание массива
//var box = [5]int{2, 4, 6, 8, 10}
//
//функция для передачи в канал значений
//func odd(channel chan int) {
//	for i := 0; i < len(box); i++ {
//вывод в канал
//		channel <- (box[i] * box[i])
//	}
//}
//
//func main() {
//создание канала
//	channelA := make(chan int)
//вызов функции в горутине
//	go odd(channelA)
//закрытие канала
//	defer close(channelA)
//	for i := 0; i < len(box); i++ {
//		fmt.Fprintln(os.Stdout, <-channelA)
//	}
//}
