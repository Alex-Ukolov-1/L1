package main

// Разработать программу, которая переворачивает слова в строке.
// пример

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow»

func main() {
	var arr1 string = "snow dog sun"
	arr := strings.Split(arr1, " ")
	arr1 = ""
	for i := len(arr); i > 0; i-- {
		arr1 += " " + (arr[i-1])
	}
	fmt.Println(arr1)
}

//как альтернатива
//  arr1 := "snow dog sun"
//	arr := strings.Split(arr1, " ")
//	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
//		arr[i], arr[j] = arr[j], arr[i]
//	}
//	fmt.Println(arr)
