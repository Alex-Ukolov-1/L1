//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

//функция преобразует данные из массива в карту подмножеств
func Set(array *[]string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, value := range *array {
		if _, ok := result[value]; !ok {
			result[value] = struct{}{}
		}
	}
	return result
}

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(array)
	fmt.Println(Set(&array))
}