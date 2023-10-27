package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(needle int, haystack []int) (int, bool) {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return 0, false
	}
	return low, true
}

func main() {
	needle := 45
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	idx, found := binarySearch(needle, items)
	if found {
		fmt.Printf("Found %d at index %d\n", needle, idx)
	} else {
		fmt.Printf("Needle %d not found.\n", needle)
	}
}
