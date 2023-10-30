package main

import (
	"fmt"
	"time"
)

// Задача 25.
// Реализовать собственную функцию sleep.

// Решение
/* 1) Принимаем на ввод время в секундах минутах и часах.
   2) Создаем переменную структуры типа time.Duration.
   3) Используем метод time.After в функции mySleep,
      который пишет в канал типа time.Time по прошествию определенного периода времени.
*/
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
