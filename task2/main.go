package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	s := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	for idx, val := range s {
		wg.Add(1)
		go func(val, idx int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val, idx, wg)
	}
	wg.Wait()
}
