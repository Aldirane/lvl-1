package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	strSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	strSet := make(map[string]bool)
	for _, val := range strSlice {
		strSet[val] = true
	}
	fmt.Printf("%+v\n", strSet)
}
