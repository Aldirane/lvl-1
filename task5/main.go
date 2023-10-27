package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func main() {
	var N int
	fmt.Scan(&N)
	t := time.Duration(N * int(time.Second))
	inOutCh := make(chan interface{})
	go func() {
		for i := 0; ; i++ {
			inOutCh <- i
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
