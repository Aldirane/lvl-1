package main

import (
	"fmt"
	"sync"
)

// Задание 7.
// Реализовать конкурентную запись данных в map.

// Решение
/*
	Так как чтобы записать значение в map мы обращаемся по ключу,
	то мы не можем конкурентно без блокировок избежать одновременную перезапись значения
	при работе нескольких горутин.
	Для этого мы должны использовать sync.Mutex и блокировать операции над map.

	- В цикле мы запускаем 100 горутин.

	1) В первой функции будет паника при попытке внести изменения в concurrentMap.

	2) Во второй функции доступ к словарю/мапе происходит с использованием блокировок,
	таким образом каждая горутина поочередно выполнится как только разблокируется доступ к словарю.

	!!! Если расскоментировать 40 строку где происходит обновление concurrentMap - будет паника.
*/

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
