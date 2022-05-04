package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


//функция , которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
func gouru1(wg *sync.WaitGroup, n int) {
	c := make(chan int)
	go func() {
		for true {
			s := rand.Intn(1000)
			c <- s
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}
	}()
	go func() {
		for true {
			fmt.Println(<-c)
		}
	}()
	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	gouru1(&wg, 10)
}