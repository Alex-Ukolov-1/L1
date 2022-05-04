package main

import (
	"fmt"
)
//структура хьюман
type Human struct {
	age     int
	name    string
	sername string
	height  int
	weight  int
}

//структура актион
type Action struct {
	Human
	grow  int
	sleep int
	eat   int
}

//функция репорт (с доступом к Human),которая выводит данные 
func (a *Human) report() {
	fmt.Println("Age", a.age, "Name", a.name, "Sername", a.sername, "Height", a.height, "Weight", a.weight)
}

func main() {
	//инициализация переменной структуры Action , которая в композиции с структурой human 
	conv := Action{Human{20, "ALEX", "UKOLOV", 185, 90}, 0, 0, 0}
	//вызов из переменной структуры action{human} метода report()
	conv.report()
}
