package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

func main() {
	var countUpTo, starter, goroutines int
	fmt.Println("Input: count up to, starter, goroutines")
	fmt.Scanf("%d %d %d", &countUpTo, &starter, &goroutines)
	countSome := Counter{
		wg:         new(sync.WaitGroup),
		mu:         new(sync.RWMutex),
		countUpTo:  countUpTo,
		starter:    starter,
		goroutines: goroutines,
	}
	res := countSome.incrementCount()
	fmt.Printf("Result = %d\n", res)
}

type Counter struct {
	wg         *sync.WaitGroup
	mu         *sync.RWMutex
	countUpTo  int
	starter    int
	goroutines int
}

func (cnt *Counter) incrementCount() int {
	for i := 0; i < cnt.goroutines; i++ {
		cnt.wg.Add(1)
		go func(cnt *Counter) {
			defer cnt.wg.Done()
			for {
				if cnt.starter < cnt.countUpTo {
					if cnt.mu.TryLock() {
						if cnt.starter < cnt.countUpTo {
							cnt.starter += 1

						}
					} else {
						continue
					}
					cnt.mu.Unlock()
				} else {
					break
				}

			}
		}(cnt)
	}
	cnt.wg.Wait()
	return cnt.starter
}
