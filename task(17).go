package main

import (
	"fmt"
	"sort"
)

//бинарный поиск встроенными методами языка
func main() {
	mas := []int{999, 342, 234, 2, 0, 23, 4363, 235}
	sort.Ints(mas)
	a := sort.SearchInts(mas, 23)
	fmt.Println(mas[a])
}
