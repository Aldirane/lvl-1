package main

import "fmt"

// Задача 13.
// Поменять местами два числа без создания временной переменной.

/* Решение
   1) замена переменных путем присваивания друг друга в одной строке без временной переменной.
   2) Побитовое исключение xor - значения одной переменной из суммы двух переменных.
   3) Создание указателей и получение значений из указателей с помощью *.
*/

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
