package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
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
