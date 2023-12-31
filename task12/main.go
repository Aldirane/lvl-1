package main

import "fmt"

/* Задача 12.
   Имеется последовательность строк - (cat, cat, dog, cat, tree)
   создать для нее собственное множество.
*/

/* Решение
   1) Создаем срез заданных строк.
   2) Создаем пустую хеш таблицу.
   3) И добавляем строки как ключи в хеш таблицу
   - Значения ключей уникальны и будут множеством.
*/

func main() {
	strSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	strMap := make(map[string]bool)
	for _, val := range strSlice {
		strMap[val] = true
	}
	setSlice := []string{}
	for key := range strMap {
		setSlice = append(setSlice, key)
	}
	fmt.Printf("%v\n", setSlice)
}
