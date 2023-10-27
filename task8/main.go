package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Дана переменная int64. Разработать программу,
// которая устанавливает i-й бит в 1 или 0.

func main() {
	var num int64
	fmt.Println("Введите число")
	fmt.Scan(&num)
	binNum := fmt.Sprintf("%b", num)
	fmt.Printf("Binary num %s\n", binNum)
	l := len(binNum)
	var bit int
	var val string
	fmt.Printf("Всего битов в числе %d = %d\n", num, l)
	fmt.Printf("Введите какой бит по счету поменять до %d бита\n", l)
	fmt.Println("Введите на какое значение поменять: 0 или 1")
	fmt.Scan(&bit, &val)
	var bSlice []string
	for i := l; i > 0; i-- {
		if i == bit {
			bSlice = append(bSlice, val)
		} else {
			bSlice = append(bSlice, string(binNum[l-i]))
		}
	}
	binNum = strings.Join(bSlice, "")
	fmt.Println(binNum)
	res := BinaryStringToInt(binNum)
	fmt.Printf("Получилось число %d\n", res)
}

func BinaryStringToInt(binStr string) int {
	var sum int
	l := len(binStr)
	for i := 0; i < l; i++ {
		x, _ := strconv.Atoi(string(binStr[i]))
		s := func(x, i int) int {
			if x == 0 {
				return 0
			} else if i == 1 {
				return 1
			}
			var res int = 2
			for j := 2; j < i; j++ {
				res *= 2
			}
			return res
		}(x, l-i)
		sum += s
	}
	return sum
}
