package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	var a, b = 1, 2
	fmt.Printf("Before a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After a = %d, b = %d\n", a, b)
	a, b = a^(a+b), b^(a+b)
	fmt.Printf("After xor a = %d, b = %d\n", a, b)
	pointA := &a
	pointB := &b
	a, b = *pointB, *pointA
	fmt.Printf("After point a = %d, b = %d\n", a, b)
}
