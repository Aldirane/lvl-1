package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Задание 11.
// Реализовать пересечение двух неупорядоченных множеств.

/* Решение
   1) Считываем две строки.
   2) Переводим их в срез строк.
   3) Добавляем строки как ключи в хеш таблицу.
   4) Так делаем для обеих хеш мап.
   5) Вызываем функцию intersection,
    - в которой в цикле for range обходит первую хеш таблицу по ключам
    - и если ключ есть во второй таблице добавляет его в срез строк.
	- Это и будет пересечением множеств.
*/

func main() {
	set1 := make(map[string]bool)
	set2 := make(map[string]bool)
	reader1 := bufio.NewReader(os.Stdin)
	str1, _ := reader1.ReadString('\n')
	str1 = strings.Trim(str1, "\n")
	str1Slice := strings.Split(str1, " ")
	for _, val := range str1Slice {
		set1[string(val)] = true
	}
	reader2 := bufio.NewReader(os.Stdin)
	str2, _ := reader2.ReadString('\n')
	str2 = strings.Trim(str2, "\n")
	str2Slice := strings.Split(str2, " ")
	for _, val := range str2Slice {
		set2[string(val)] = true
	}
	intersection := intersection(set1, set2)
	fmt.Println(intersection)
}

func intersection(set1, set2 map[string]bool) []string {
	intersection := new([]string)
	for key := range set1 {
		if _, ok := set2[key]; ok {
			*intersection = append(*intersection, key)
		}
	}
	return *intersection
}
