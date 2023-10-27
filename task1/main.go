package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

func main() {
	a := Action{Human: Human{Name: "Bob", Age: 20}}
	a.walk()
}

type Human struct {
	Name string
	Age  int8
}
type Action struct {
	Human
}

func (a *Action) walk() {
	fmt.Printf("%s is walking\n", a.Name)
}
