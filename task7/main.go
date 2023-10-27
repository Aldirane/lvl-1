package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {
	//var dataNotMutex string = "notMutexData"
	var dataMutex string = "MutexData"
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	concurrentMap := make(map[string]int)
	mutexMap := make(map[string]int)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		// Goroutine without mutex will throw concurrent map writes error due to race condition
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			//concurrentMap[dataNotMutex] += i
		}(i, wg)
		wg.Add(1)
		// Goroutine with mutex no errors
		go func(i int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			mutexMap[dataMutex] += i
			mu.Unlock()
		}(i, wg, mu)
	}
	wg.Wait()
	fmt.Printf("%+v\n", concurrentMap)
	fmt.Printf("%+v\n", mutexMap)
}
