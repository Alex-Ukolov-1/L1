package main

import (
	"fmt"
	//"sync"
)

var box = [5]int{2, 4, 6, 8, 10}
var keys = map[int]int{}

//функция записи данных в канал
func odd(channel chan int) {
	for i := 0; i < len(box); i++ {
		channel <- (box[i])
	}
}

//реализация записи данных в map
func main() {
	channelA := make(chan int)
	go odd(channelA)
	for i := 0; i < len(box); i++ {
		keys[i] = (<-channelA)
	}
	defer close(channelA)
	fmt.Println(keys)

}

//package main
// Принимает map m, данные i, WaitGroup wg, Mutex mu.
//func write(m map[int]int, i int, wg *sync.WaitGroup, mu *sync.Mutex) {
//	// Блокирование мютекса перед записью и уменьшаем счетчик wg перед завершением.
//	mu.Lock()
//	defer wg.Done()
//	m[i] = i
//	//Разблокировие мютекса.
//	mu.Unlock()
//}

//func main() {
// Создаем map m, waitgroup wg и mutex.
//	m := make(map[int]int)
//	wg := sync.WaitGroup{}
//	mu := sync.Mutex{}
//
//	for i := 0; i < 5; i++ { //Запускаем горутины.
//		wg.Add(1)
//		go write(m, i, &wg, &mu)
//	}
//
//	// Проверка завершения всех горутин.
//	wg.Wait()
//	fmt.Println(m)
//}
