package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте. Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

func main() {
	mainCh := make(chan interface{})
	signalChan := make(chan os.Signal, 1)
	var workers int
	wg := new(sync.WaitGroup)
	fmt.Scan(&workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			data, ok := <-mainCh
			if ok {
				fmt.Println(data)
			}
		}(wg)
	}
	scan := bufio.NewScanner(os.Stdin)
	go func() {
		for scan.Scan() {
			mainCh <- scan.Text()
		}
	}()
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
