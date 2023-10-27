package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	signalChan := make(chan os.Signal, 1)
	stopChan := make(chan bool)
	go secondGo(stopChan)
	time.Sleep(2 * time.Second)
	stopChan <- true
	someChan := make(chan interface{})
	go thirdGo(someChan)
	someChan <- "data"
	someChan <- "wrong data"
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
