package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/* Задача 19.
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

/* Решение
   1) Мы сканируем с стандартного ввода строку.
   2) Добавляем строку в массив байтов.
   3) В цикле при неравенстве utf8.RuneCount(byteStr) >= 1, что означает Количество рун в срезе байтов
      обходим, декодируем первую руну из среза байтов, получаем руну и размер занимаемых ею байтов,
	  далее уменьшаем срез байтов от количества байтов последней полученной руны.
   4) Добавляем руну в срез рун.
   5) Обходим в цикле по индексам срез рун, переводим очередную руну в строку и добавляем в срез строк
   6) Получаем целую строку путем объединения среза строк.
*/

func main() {
	var str string
	var strArr []string
	fmt.Scan(&str)
	var byteStr = []byte(str)
	var strRune = []rune{}
	for utf8.RuneCount(byteStr) >= 1 {
		r, size := utf8.DecodeRune(byteStr)
		byteStr = byteStr[size:]
		strRune = append(strRune, r)
	}

	for i := len(strRune) - 1; i >= 0; i-- {
		strArr = append(strArr, string(strRune[i]))
	}
	res := strings.Join(strArr, "")
	fmt.Printf("%s\n", res)
}
