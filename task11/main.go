package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Реализовать пересечение двух неупорядоченных множеств.

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
