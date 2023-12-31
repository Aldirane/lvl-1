package main

import (
	"fmt"
	"math"
)

/* Задача 24.
   Разработать программу нахождения расстояния между двумя точками, которые
   представлены в виде структуры Point с инкапсулированными параметрами x,y и
   конструктором.
*/

/* Решение
   1) Задаем структуру Point с двумя координатами x и y.
   2) Создаем новую структуру с помощью метода NewPoint.
   3) Находим расстояние функцией Distance,
   по теореме пифагора находим гипотенузу между двумя точками на плоскости.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func Distance(p1, p2 *Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	distance := Distance(p1, p2)
	fmt.Printf("Расстояние между точками: %f\n", distance)
}
