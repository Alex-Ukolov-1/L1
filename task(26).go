package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные
// Функция проверки должна быть регистронезависимой.

func main() {
	s := "aBcD"
	//переводит всё в нижний регистр
	s = strings.ToLower(s)
	//создание массива и создание пробела между буквами
	ss := strings.Split(s, "")
	dd := ""
	//алгоритм для поиска уникальности символов
	for i := 0; i < len(ss); i++ {
		var b = 0
		for j := 0; j < len(ss); j++ {
			var aa string = ss[i]
			if aa == ss[j] {
				b++
			}
		}
		if b > 1 {
			b = 0
		} else {
			dd += ss[i]
		}
	}
	//fmt.Println(dd)
	if dd == s {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
