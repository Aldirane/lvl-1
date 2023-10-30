package main

import "fmt"

// Задача 17.
// Реализовать бинарный поиск встроенными методами языка.

/* Решение
   Бинарный поиск позволяет за O log(N) найти искомое значение.
   Однако он работает только с отсортированным срезом.
   Ключевой принцип это поиск искомого значения начиная с середины.
   Если искомое значение меньше середины, то нужно искать в левой части от середины,
   и наоборот если больше то в правой части среза. Для этого мы постоянно в цикле
   вычисляем середины, и двигаем нижнюю или верхнюю границу
   в зависимости от неравенства искомого значения и середины.
   Если они равны, значит мы нашли значение по индексу равному текущей "середине".
   Или если нижняя граница стала больше верхней, значит значение не найдено.
*/

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
