package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	var s, m, h int
	fmt.Println("Введите секунды s минуты m часы h")
	fmt.Scan(&s, &m, &h)
	fmt.Println("Start")
	dur := time.Duration(s*int(time.Second) + m*int(time.Minute) + h*int(time.Hour))
	mySleep(dur)
	fmt.Println("End")
}
