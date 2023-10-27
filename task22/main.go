package main

import (
	"fmt"
	"math"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20

func main() {
	var a, b float64 = math.Pow(2, 25), math.Pow(2, 21)
	fmt.Printf("%.2f\n", a*b)
	fmt.Printf("%.2f\n", a/b)
	fmt.Printf("%.2f\n", a+b)
	fmt.Printf("%.2f\n", a-b)
}
