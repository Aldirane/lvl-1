package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
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
