package main

//Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором.

import (
	"fmt"
	"math"
)

//структура point
type Point struct {
	x float64
	y float64
}

//функция установки значений для структуры point
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

//функция нахождения расстояния между двумя точками
func (atr *Point) Distance(point1, point2 *Point) float64 {
	a := point2.x - point1.x
	b := point2.y - point1.y
	return math.Sqrt(a*a + b*b)
}

func main() {
	A := NewPoint(1.0, 1.0)
	B := NewPoint(3.0, 3.0)
	fmt.Println(A.Distance(A, B))
}
