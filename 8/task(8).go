package main

//Дана переменная int64. Разработать программу которая устанавливает n-й бит в 1 или 0

import (
	"fmt"
	"strconv"
	"strings"
)

//функция преобразования числа в биты , поиска  и замены в бит
func IntToBytes(num int64, position int, target int) (a string) {
	if target != 0 && target != 1 {
		return
	} else {
		v := strconv.FormatInt(num, 2)
		arr := strings.Split(v, "")
		for i := 0; i < len(v); i++ {
			if i == position {
				arr[i] = strconv.Itoa(target)
			}
		}
		bb := strings.Join(arr, "")
		v = bb
		return v
	}
}

func main() {
	//первая переменная - число которое надо в биты, вторая - номер по порядку битов
	//третия - на что меняем
	v := IntToBytes(20, 3, 1)
	//20->10100
	//20->10110
	fmt.Println(v)
}
