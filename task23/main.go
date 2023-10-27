package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Удалить i-ый элемент из слайса

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
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func orderRemove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
