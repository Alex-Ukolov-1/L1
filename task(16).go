package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//первый способ сортировки
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	// выбираем точку
	pivotIndex := rand.Int() % len(a)
	// перемещаемся на право
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	// Элементы меньше, чем ось слева
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// Помещаем после последнего меньшего элемента
	a[left], a[right] = a[right], a[left]
	// спуск
	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}

//второй способ сортировки пузырьком
func sorrtt(mags []int) {
	var temp int
	for i := 0; i < len(mags); i++ {
		for j := 0; j < len(mags); j++ {
			if mags[i] < mags[j] {
				temp = mags[i]
				mags[i] = mags[j]
				mags[j] = temp
			}
		}
	}
	fmt.Println(mags)
}

func main() {
	//первый способ встроенный
	arr := []int{200, 100, 5, -300, -100}
	sort.Ints(arr)
	fmt.Println(arr)
	//второй способ
	mags := []int{200, 100, 5, -300, -100}
	sorrtt(mags)
	//третий способ
	cups := []int{200, 100, 5, -300, -100}
	v := quicksort(cups)
	fmt.Println(v)
}
