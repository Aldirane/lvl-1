package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов с использованием конкурентных вычислений.
*/

func main() {
	var sum int
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	s := []int{2, 4, 6, 8, 10}
	for _, val := range s {
		wg.Add(1)
		go func(mu *sync.Mutex, wg *sync.WaitGroup, val int) {
			defer wg.Done()
			mu.Lock()
			sum += val * val
			mu.Unlock()
		}(&mu, &wg, val)
	}
	wg.Wait()
	fmt.Println(sum)
}
