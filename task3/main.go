package main

import (
	"fmt"
	"sync"
)

/* Задание 3
   Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
   квадратов с использованием конкурентных вычислений.
*/

/* Решение
   1) Создаем две переменные:
      - mu - sync.Mutex для синхронизации действий горутин с какой либо переменной (значением)
      - wg - sync.WaitGroup - для ожидания завершения горутин wg.Add(1)

   2) Таким образом в горутинах мы блокируем взаимодействие с переменной sum mu.Lock,
	  тем самым не давая другим горутинам одновременно перезаписать
      одну и ту же переменную.
	  Как действие закончено мы разблокируем mu.Unlock и убавляем при выходе
      счетчик активных горутин wg.Done
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
