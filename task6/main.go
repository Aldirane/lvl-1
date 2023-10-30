package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Задание 6.
// Реализовать все возможные способы остановки выполнения горутины.

/* Решение
	1) Первый способ - использовать команду Ctrl+C для завершения процессов,
       с помощью os.Signal через канал signalChan и завершить главный поток функции main по получению сигнала

	2) Второй способ - ожидать получение из канала stopChan
	   и пока не было принятно ничего из этого канала, горутина будет делать какую либо работу.

	3) Третий способ - принимать из канала какие то данные.
	   Если данные валидны то продолжать, если нет то завершить работу.

	4) Четвертый способ - завершить работу горутины по таймеру. - time.After(Duration).
*/

func main() {
	var N int = 2
	signalChan := make(chan os.Signal, 1)
	stopChan := make(chan bool)
	go secondGo(stopChan)
	time.Sleep(2 * time.Second)
	stopChan <- true
	someChan := make(chan interface{})
	go thirdGo(someChan)
	someChan <- "data"
	someChan <- "wrong data"
	go fourthGo(N)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Main goroutine out")
}

func secondGo(activeChan chan bool) {
	for {
		select {
		case <-activeChan:
			fmt.Println("Second goroutine finished working")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Second goroutine do some work")
		}
	}
}

func thirdGo(someChan chan interface{}) {
	for data := range someChan {
		if data == "data" {
			fmt.Println("Third goroutine do some work")
		} else {
			close(someChan)
		}
	}
	fmt.Println("Third goroutine out")
}

func fourthGo(N int) {
	dur := time.Duration(N * int(time.Second))
	<-time.After(dur)
	fmt.Println("Fourth gorutine out")
}
