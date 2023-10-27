package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере

type Cat struct{}

func (c Cat) Purr() {
	fmt.Println("Муррр")
}

type Animal interface {
	Sound()
}

type CatAdapter struct {
	cat Cat
}

func (ca CatAdapter) Sound() {
	ca.cat.Purr()
}

func MakeSound(a Animal) {
	a.Sound()
}

func main() {
	cat := Cat{}
	catAdapter := CatAdapter{cat}
	MakeSound(catAdapter) // Вывод: Муррр
}
