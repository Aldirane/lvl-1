package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* Задача 20.
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

/* Решение
   1) Читаем из стандартного ввода os.Stdin
   2) Разбиваем строку на подстроки и получаем срез строк strSlice
   3) Индексируемся с конца по срезу строк strSlice, и добавляем в новый срез строк invertedStrArr
   4) Объединяем invertedStrArr в одну строку и получаем искомый результат.
*/

func main() {
	var invertedStrArr []string
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	str = strings.Trim(str, "\n")
	strSlice := strings.Split(str, " ")
	for i := len(strSlice) - 1; i >= 0; i-- {
		invertedStrArr = append(invertedStrArr, strSlice[i])
	}
	result := strings.Join(invertedStrArr, " ")
	fmt.Println(result)
}
