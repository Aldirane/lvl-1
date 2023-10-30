package main

import "fmt"

// Задание 16.
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

/* Решение.
   Точка входа quickSortStart - функция для запуска сортировки.
   Ключевая функция в быстрой сортировки это функция partition и переменная pivot,
   относительно которой мы сортируем влево и вправо значения по индексам.
   Если текущее значение по индексу меньше pivot мы делаем своп по начальному и текущему индексам.
   После свопа перемещаем на 1 начальный индекс. Это необходимо для того чтобы разместить все
   значения меньше pivot'а по левую сторону.
   Далее делаем замену наибольшего значения, который меньше pivot'a по их индексам
   и возвращаем получившийся срез и индекс i в котором лежит новый pivot.

   Логика самих функций quickSort - рекурсивно вызывать саму себя,
   передавая в аргументы получившийся срез и индекс нового pivot'a.
   Далее вызывать себя обратно с индексами от минимального до pivot'a.
   И максимального от pivot'a. В результате возвращая срез, если переданные как аргументы
   индексы равны. Это значит что срез отсортирован. Таким образом мы отсортируем весь срез.
*/

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
