package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

/* Задание 9.
   Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
   массива, во второй — результат операции x*2, после чего данные из второго
   канала должны выводиться в stdout.
*/

/* Решение
   1) Мы сканируем строку обрабатываем и выдаем слайс чисел функцией scanIntSclice
   2) Создаем WaitGroup и запускаем две горутины.
   - Первая горутина пишет в канал chan1 значения чисел из слайса intSlice
   - Вторая горутина читает из chan1, пишет в chan2 полученные значения,
   - потом делает вывод значений из chan2 в стандартный выход.
*/

func main() {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	intSlice := scanIntSlice()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup, intSlice []int) {
		defer wg.Done()
		defer close(chan1)
		for _, val := range intSlice {
			chan1 <- val

		}
	}(wg, intSlice)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(chan2)
		for val := range chan1 {
			chan2 <- val * 2
			fmt.Print(<-chan2, " ")
		}
	}(wg)
	wg.Wait()
}

func scanIntSlice() []int {
	var intSlice []int
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	strSlice := strings.Split(strings.Trim(strings.Trim(str, "\n"), " "), " ")
	for _, val := range strSlice {
		i, err := strconv.Atoi(string(val))
		if err != nil {
			log.Fatalf("must be integer - %s", err)
		} else {
			intSlice = append(intSlice, i)
		}
	}
	return intSlice
}
