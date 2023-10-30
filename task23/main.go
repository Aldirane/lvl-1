package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Задача 23.
// Удалить i-ый элемент из слайса

/* Решение
   1) Мы считываем строку, и индекс.
   2) Строку разбиваем на слайс.
   3) Эффективное удаление - unOrderRemove:
      - Заменяем последнее значение, вместо удаляемого по индексу.
      - Образаем конец.
   4) Неэфективное удаление - orderRemove:
      - Возвращаем новый слайс объединяя первую часть до индекса, и после индекса.
*/

func main() {
	var someSlice []string
	var i int
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Scan(&i)
	str = strings.Trim(str, "\n")
	someSlice = strings.Split(str, " ")
	unOrderSlice := unOrderRemove(someSlice, i)
	fmt.Println(unOrderSlice)
	orderSlice := orderRemove(someSlice, i)
	fmt.Println(orderSlice)
}

// Порядок не важен
func unOrderRemove(s []string, i int) []string {
	if len(s) == 1 {
		return nil
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func orderRemove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
