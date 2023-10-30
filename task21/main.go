package main

import "fmt"

// Задача 21.
// Реализовать паттерн «адаптер» на любом примере

/* Решение
   Паттерн адаптер применяется тогда когда необходимо внедрить
   какую-либо структуру в другую структуру, но не напрямую, а используя
   другую структуру реализующую сигнатуру той структуры в которую надо внедрится.
   Таким образом у структуры Cat есть метод Purr, но нет метода Sound
   Мы создаем CatAdapter, внедряем в него структуру Cat,
   и создаем метод Sound в котором используем метод Purr структуры Cat.
   Таким образом мы адаптировали структуру Cat под сигнатуру структуры Animal.
*/

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
