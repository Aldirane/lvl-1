package main

import (
	"fmt"
	"time"
)

/* Задача 5
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

/* Решение
Мы принимаем int N и задаем продолжительность в секундах N * time.Second
Затем запускаем две горутины, 1я отправляет значения в бесконечном цикле в канал inOutCh
2я принимает значения из канал. Канал небуферизуется, то есть то что отправляется сразу можно принимать.
*/

func main() {
	var N int
	fmt.Scan(&N)
	t := time.Duration(N * int(time.Second))
	inOutCh := make(chan interface{})
	go func() {
		for i := 0; ; i++ {
			inOutCh <- i
			time.Sleep(1000 * time.Microsecond)
		}
	}()
	go func() {
		for val := range inOutCh {
			fmt.Print(val, " ")
		}
	}()
	time.Sleep(t)
	fmt.Println("\nTime Out")
}
