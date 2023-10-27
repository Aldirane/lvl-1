package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	intArr := []int{5, 6, 7, 2, 1, 0}
	fmt.Printf("Before %v\n", intArr)
	fmt.Printf("After %v\n", quickSortStart(intArr))
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i

}
