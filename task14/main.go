package main

import "fmt"

/* Задача 14.
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

// Решение
/*
   Используя операторы switch case мы приводим к нужному нам типу.
   В функции whichType принимается аргумент типа interface.
   В блоке switch приводим к типу type. Если type int, выполнится case int итд.
   Приводить можно любой тип в том числе структуру.
   Если ни один case не выполнился, то сработает default
*/

func main() {
	var unknownType interface{}
	// unknownType = 1
	// unknownType = "1AC"
	// unknownType = true
	// unknownType = make(chan string)
	// unknownType = make(chan int)
	unknownType = make(chan bool)
	whichType(unknownType)

}

func whichType(unknownType interface{}) {
	switch v := unknownType.(type) {
	case int:
		fmt.Printf("Type int %d\n", v)
	case string:
		fmt.Printf("Type string %s\n", v)
	case bool:
		fmt.Printf("Type bool %v\n", v)
	case chan string:
		fmt.Printf("Chan string %v\n", v)
	case chan int:
		fmt.Printf("Chan int %v\n", v)
	case chan bool:
		fmt.Printf("Chan bool %v\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
