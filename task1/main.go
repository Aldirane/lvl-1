package main

import "fmt"

/* Задание 1
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

/* Решение
Встраиваем структуру Human в структуру Action и вызываем метод walk у Human
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

func (h *Human) walk() {
	fmt.Printf("%s is walking\n", h.Name)
}
