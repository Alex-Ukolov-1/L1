package main

import (
	"fmt"
)

//Удалить i-ый элемент из слайса.

//функция удаления i-элемента из слайса
func first(arr []string, i int) {
	deleted := append(arr[:i], arr[i+1:]...)
	fmt.Println(deleted)
}

func main() {
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	first(arr, 2)
}
