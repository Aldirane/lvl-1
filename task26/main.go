package main

import (
	"fmt"
	"strings"
)

/* Задача 26.
   Разработать программу, которая проверяет, что все символы в строке
   уникальные (true — если уникальные, false etc). Функция проверки должна быть
   регистронезависимой.
   Например:
   abcd — true
   abCdefAaf — false
   aabcd — false
*/

/* Решение
   1) создаем хеш-мапу unique.
   2) Считываем строку.
   3) Проходим по строке в цикле и добавляем очередной символ в хеш-мапу.
      Если там уже был ключ с таким значением,
	  тогда выходим из цикла и записываем в результат false. Иначе будет true.
*/

func main() {
	var unique = make(map[string]int)
	var str string
	fmt.Scan(&str)
	var res bool = true
	for _, val := range str {
		c := strings.ToLower(string(val))
		if _, ok := unique[c]; !ok {
			unique[c] = 1
		} else {
			res = false
			break
		}
	}
	fmt.Println(res)
}
