//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.

package main

import "fmt"

//функция объединяет значения , которые  близки между собой
func AddElement(set map[int][]float64, element float64) {
	//создание ключа
	key := (int(element) / 10) * 10
	//установление переменной карты
	value := set[key]
	//добавление элемента группе
	set[key] = append(value, element)

}
func main() {
	set := make(map[int][]float64)
	array := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _, element := range array {
		AddElement(set, element)
	}
	fmt.Println(set)
}
