package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
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
